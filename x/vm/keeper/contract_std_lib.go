package keeper

import (
	"fmt"

	"vesta/x/vm/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dop251/goja"
)

func (k Keeper) applyStandardLib(ctx sdk.Context, creator sdk.AccAddress, contractName string, contractAddress sdk.AccAddress, vm *goja.Runtime, readonly bool) {
	std := vm.NewObject()

	err := std.Set("SENDER", creator.String())
	if err != nil {
		ctx.Logger().Error(err.Error())
		return
	}
	err = std.Set("CONTRACT", contractAddress.String())
	if err != nil {
		ctx.Logger().Error(err.Error())
		return
	}

	err = std.Set("NAME", contractName)
	if err != nil {
		ctx.Logger().Error(err.Error())
		return
	}

	for _, v := range k.injectors {
		err := v.Inject(ctx, creator, contractName, contractAddress, std, readonly)
		if err != nil {
			ctx.Logger().Error(err.Error())
			return
		}
	}

	if !readonly {

		err = std.Set("write", func(call goja.FunctionCall) goja.Value {
			key := call.Argument(0).String()
			data := call.Argument(1).String()

			size, ok := sdk.NewIntFromString(fmt.Sprintf("%d", len(data)))
			if !ok {
				return goja.ValueFalse()
			}

			toSave := types.Romdata{
				Index: fmt.Sprintf("%s%s", contractAddress.String(), key),
				Data:  data,
			}

			k.SetRomdata(ctx, toSave)

			ctx.GasMeter().ConsumeGas(size.Uint64()*types.DefaultGasValues().Write, "saving data")

			return goja.ValueTrue()
		})
		if err != nil {
			ctx.Logger().Error(err.Error())
			return
		}
	}

	err = std.Set("read", func(call goja.FunctionCall) goja.Value {
		key := call.Argument(0).String()

		data, found := k.GetRomdata(ctx, fmt.Sprintf("%s%s", contractAddress.String(), key))
		if !found {
			return goja.Null()
		}

		size, ok := sdk.NewIntFromString(fmt.Sprintf("%d", len(data.Data)))
		if !ok {
			return goja.Null()
		}

		ctx.GasMeter().ConsumeGas(size.Uint64()*types.DefaultGasValues().Read, "reading data")

		val, err := goja.New().RunString(data.Data)
		if err != nil {
			return goja.Null()
		}

		return val
	})
	if err != nil {
		ctx.Logger().Error(err.Error())
		return
	}

	err = std.Set("import", func(call goja.FunctionCall) goja.Value {
		moduleName := call.Argument(0).String()

		program, found := k.GetProgram(ctx, moduleName)
		if !found {
			return goja.Null()
		}

		source, ok := k.GetContracts(ctx, program.Code)
		if !ok {
			return goja.Null()
		}

		code := source.Source

		newvm := goja.New()
		k.applyStandardLib(ctx, creator, contractName, contractAddress, newvm, readonly)

		_, err := newvm.RunString(code)
		if err != nil {
			return goja.Null()
		}

		r := newvm.Get("contractFunctions")
		if r == nil {
			return goja.Null()
		}

		ctx.GasMeter().ConsumeGas(types.DefaultGasValues().Import, "importing library")

		return r
	})
	if err != nil {
		ctx.Logger().Error(err.Error())
		return
	}

	err = std.Set("fetch", func(call goja.FunctionCall) goja.Value {
		moduleName := call.Argument(0).String()
		entryPoint := call.Argument(1).String()
		fetchType := call.Argument(2).String()
		args := call.Arguments[3:]

		ctx.GasMeter().ConsumeGas(types.DefaultGasValues().Import, "fetching contract")

		program, found := k.GetProgram(ctx, moduleName)
		if !found {
			return goja.Null()
		}

		source, ok := k.GetContracts(ctx, program.Code)
		if !ok {
			return goja.Null()
		}

		code := source.Source

		var res string
		var err error

		if fetchType == "POST" {
			res, err = k.executeContract(ctx, program.Name, code, entryPoint, creator, args)
			if err != nil {
				return goja.Null()
			}
		} else if fetchType == "GET" {
			res, err = k.queryContract(ctx, program.Name, code, entryPoint, args)
			if err != nil {
				return goja.Null()
			}
		} else {
			return goja.Null()
		}

		val, err := goja.New().RunString(res)
		if err != nil {
			return goja.Null()
		}

		return val
	})
	if err != nil {
		ctx.Logger().Error(err.Error())
		return
	}

	err = vm.Set("std", std)
	if err != nil {
		ctx.Logger().Error(err.Error())
		return
	}
}
