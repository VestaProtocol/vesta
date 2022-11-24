package keeper

import (
	"fmt"

	"vesta/x/vm/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dop251/goja"
)

func (k Keeper) applyStandardLib(ctx sdk.Context, creator sdk.AccAddress, contractAddress sdk.AccAddress, vm *goja.Runtime, readonly bool) {
	err := vm.Set("SENDER", creator.String())
	if err != nil {
		ctx.Logger().Error(err.Error())
		return
	}
	err = vm.Set("CONTRACT", contractAddress.String())
	if err != nil {
		ctx.Logger().Error(err.Error())
		return
	}

	if !readonly {
		err = vm.Set("sendTokens", func(call goja.FunctionCall) goja.Value {
			reciever, err := sdk.AccAddressFromBech32(call.Argument(0).String())
			if err != nil {
				return goja.ValueFalse()
			}

			coin, err := sdk.ParseCoinNormalized(call.Argument(1).String())
			if err != nil {
				return goja.ValueFalse()
			}

			err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, creator, types.ModuleName, sdk.NewCoins(coin))
			if err != nil {
				return goja.ValueFalse()
			}
			err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, reciever, sdk.NewCoins(coin))
			if err != nil {
				return goja.ValueFalse()
			}
			return goja.ValueTrue()
		})
		if err != nil {
			ctx.Logger().Error(err.Error())
			return
		}

		err = vm.Set("withdrawTokens", func(call goja.FunctionCall) goja.Value {
			reciever, err := sdk.AccAddressFromBech32(call.Argument(0).String())
			if err != nil {
				return goja.ValueFalse()
			}

			coin, err := sdk.ParseCoinNormalized(call.Argument(1).String())
			if err != nil {
				return goja.ValueFalse()
			}

			err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, contractAddress, types.ModuleName, sdk.NewCoins(coin))
			if err != nil {
				return goja.ValueFalse()
			}
			err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, reciever, sdk.NewCoins(coin))
			if err != nil {
				return goja.ValueFalse()
			}
			return goja.ValueTrue()
		})
		if err != nil {
			ctx.Logger().Error(err.Error())
			return
		}

		err = vm.Set("save", func(call goja.FunctionCall) goja.Value {
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

	err = vm.Set("read", func(call goja.FunctionCall) goja.Value {
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

	err = vm.Set("require", func(call goja.FunctionCall) goja.Value {
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
		k.applyStandardLib(ctx, creator, contractAddress, newvm, readonly)

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

	err = vm.Set("fetch", func(call goja.FunctionCall) goja.Value {
		moduleName := call.Argument(0).String()
		entryPoint := call.Argument(1).String()
		args := call.Arguments[2:]

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

		if !readonly {
			res, err = k.executeContract(ctx, program.Name, code, entryPoint, creator, args)
			if err != nil {
				return goja.Null()
			}
		} else {
			res, err = k.queryContract(ctx, program.Name, code, entryPoint, args)
			if err != nil {
				return goja.Null()
			}
		}

		ctx.GasMeter().ConsumeGas(types.DefaultGasValues().Import, "fetching contract")

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
}
