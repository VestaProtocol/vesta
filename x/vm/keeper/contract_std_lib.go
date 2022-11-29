package keeper

import (
	"fmt"

	"github.com/TheMarstonConnell/vesta/x/vm/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/dop251/goja"
)

func (k Keeper) applyStandardLib(ctx sdk.Context, creator sdk.AccAddress, contractName string, contractAddress sdk.AccAddress, vm *goja.Runtime, readonly bool) {
	std := vm.NewObject()
	context := vm.NewObject()

	contractExports := vm.NewObject()
	contractFunctions := vm.NewObject()
	contractQueries := vm.NewObject()

	err := context.Set("SENDER", creator.String())
	if err != nil {
		ctx.Logger().Error(err.Error())
		return
	}
	err = contractExports.Set("address", contractAddress.String())
	if err != nil {
		ctx.Logger().Error(err.Error())
		return
	}

	err = contractExports.Set("name", contractName)
	if err != nil {
		ctx.Logger().Error(err.Error())
		return
	}

	err = contractExports.Set("functions", contractFunctions)
	if err != nil {
		ctx.Logger().Error(err.Error())
		return
	}

	err = contractExports.Set("queries", contractQueries)
	if err != nil {
		ctx.Logger().Error(err.Error())
		return
	}

	err = vm.Set("CONTRACT", contractExports)
	if err != nil {
		ctx.Logger().Error(err.Error())
		return
	}

	for _, v := range k.injectors {
		module := vm.NewObject()
		err := std.Set(v.Name(), module)
		if err != nil {
			ctx.Logger().Error(err.Error())
			return
		}

		err = v.Inject(ctx, vm, creator, contractName, contractAddress, module, readonly)
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

	err = std.Set("require", func(call goja.FunctionCall) goja.Value {
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

		ctc := vm.Get("CONTRACT").ToObject(vm)
		if ctc == nil {
			e := sdkerrors.Wrapf(sdkerrors.ErrKeyNotFound, "cannot find contract")
			ctx.Logger().Error(e.Error())
			return goja.Null()
		}

		ctx.GasMeter().ConsumeGas(types.DefaultGasValues().Import, "importing library")

		return ctc
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

		return vm.ToValue(res)
	})
	if err != nil {
		ctx.Logger().Error(err.Error())
		return
	}

	err = vm.Set("STD", std)
	if err != nil {
		ctx.Logger().Error(err.Error())
		return
	}

	err = vm.Set("CTX", context)
	if err != nil {
		ctx.Logger().Error(err.Error())
		return
	}
}
