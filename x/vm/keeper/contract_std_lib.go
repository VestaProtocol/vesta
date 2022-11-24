package keeper

import (
	"fmt"

	"vesta/x/vm/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/robertkrimen/otto"
)

func (k Keeper) applyStandardLib(ctx sdk.Context, creator sdk.AccAddress, contractAddress sdk.AccAddress, vm *otto.Otto, readonly bool) {
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
		err = vm.Set("sendTokens", func(call otto.FunctionCall) otto.Value {
			reciever, err := sdk.AccAddressFromBech32(call.Argument(0).String())
			if err != nil {
				return otto.FalseValue()
			}

			coin, err := sdk.ParseCoinNormalized(call.Argument(1).String())
			if err != nil {
				return otto.FalseValue()
			}

			err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, creator, types.ModuleName, sdk.NewCoins(coin))
			if err != nil {
				return otto.FalseValue()
			}
			err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, reciever, sdk.NewCoins(coin))
			if err != nil {
				return otto.FalseValue()
			}
			return otto.TrueValue()
		})
		if err != nil {
			ctx.Logger().Error(err.Error())
			return
		}

		err = vm.Set("withdrawTokens", func(call otto.FunctionCall) otto.Value {
			reciever, err := sdk.AccAddressFromBech32(call.Argument(0).String())
			if err != nil {
				return otto.FalseValue()
			}

			coin, err := sdk.ParseCoinNormalized(call.Argument(1).String())
			if err != nil {
				return otto.FalseValue()
			}

			err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, contractAddress, types.ModuleName, sdk.NewCoins(coin))
			if err != nil {
				return otto.FalseValue()
			}
			err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, reciever, sdk.NewCoins(coin))
			if err != nil {
				return otto.FalseValue()
			}
			return otto.TrueValue()
		})
		if err != nil {
			ctx.Logger().Error(err.Error())
			return
		}

		err = vm.Set("save", func(call otto.FunctionCall) otto.Value {
			key := call.Argument(0).String()
			data := call.Argument(1).String()

			size, ok := sdk.NewIntFromString(fmt.Sprintf("%d", len(data)))
			if !ok {
				return otto.FalseValue()
			}

			toSave := types.Romdata{
				Index: fmt.Sprintf("%s%s", contractAddress.String(), key),
				Data:  data,
			}

			k.SetRomdata(ctx, toSave)

			ctx.GasMeter().ConsumeGas(size.Uint64()*types.DefaultGasValues().Write, "saving data")

			return otto.TrueValue()
		})
		if err != nil {
			ctx.Logger().Error(err.Error())
			return
		}
	}

	err = vm.Set("read", func(call otto.FunctionCall) otto.Value {
		key := call.Argument(0).String()

		data, found := k.GetRomdata(ctx, fmt.Sprintf("%s%s", contractAddress.String(), key))
		if !found {
			return otto.NullValue()
		}

		size, ok := sdk.NewIntFromString(fmt.Sprintf("%d", len(data.Data)))
		if !ok {
			return otto.NullValue()
		}

		ctx.GasMeter().ConsumeGas(size.Uint64()*types.DefaultGasValues().Read, "reading data")

		val, err := otto.New().Eval(data.Data)
		if err != nil {
			return otto.NullValue()
		}

		return val
	})
	if err != nil {
		ctx.Logger().Error(err.Error())
		return
	}

	err = vm.Set("require", func(call otto.FunctionCall) otto.Value {
		moduleName := call.Argument(0).String()

		program, found := k.GetProgram(ctx, moduleName)
		if !found {
			return otto.NullValue()
		}

		source, ok := k.GetContracts(ctx, program.Code)
		if !ok {
			return otto.NullValue()
		}

		code := source.Source

		newvm := otto.New()
		k.applyStandardLib(ctx, creator, contractAddress, newvm, readonly)

		_, err := newvm.Run(code)
		if err != nil {
			return otto.NullValue()
		}

		r, err := newvm.Get("contractFunctions")
		if err != nil {
			return otto.NullValue()
		}

		ctx.GasMeter().ConsumeGas(types.DefaultGasValues().Import, "importing library")

		return r
	})
	if err != nil {
		ctx.Logger().Error(err.Error())
		return
	}

	err = vm.Set("fetch", func(call otto.FunctionCall) otto.Value {
		moduleName := call.Argument(0).String()
		entryPoint := call.Argument(1).String()
		argList := call.ArgumentList[2:]

		var args []string = make([]string, len(argList))
		for i, v := range argList {
			args[i] = v.String()
		}

		program, found := k.GetProgram(ctx, moduleName)
		if !found {
			return otto.NullValue()
		}

		source, ok := k.GetContracts(ctx, program.Code)
		if !ok {
			return otto.NullValue()
		}

		code := source.Source

		var res string
		var err error

		if !readonly {
			res, err = k.executeContract(ctx, program.Name, code, entryPoint, creator, args)
			if err != nil {
				return otto.NullValue()
			}
		} else {
			res, err = k.queryContract(ctx, program.Name, code, entryPoint, args)
			if err != nil {
				return otto.NullValue()
			}
		}

		ctx.GasMeter().ConsumeGas(types.DefaultGasValues().Import, "fetching contract")

		val, err := otto.New().Eval(res)
		if err != nil {
			return otto.NullValue()
		}

		return val
	})
	if err != nil {
		ctx.Logger().Error(err.Error())
		return
	}
}
