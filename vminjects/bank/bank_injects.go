package bankinjects

import (
	"fmt"

	"github.com/VestaProtocol/roma/x/vm/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dop251/goja"
)

type Injector struct {
	keeper interface{}
}

func (i *Injector) Name() string {
	return "bank"
}

func NewInjector(keeper interface{}) *Injector {
	return &Injector{
		keeper: keeper,
	}
}

func (i *Injector) Inject(ctx sdk.Context, gj *goja.Runtime, creator sdk.AccAddress, contractName string, contractAddress sdk.AccAddress, vm *goja.Object, readonly bool) error {
	if !readonly {

		err := vm.Set("sendTokens", func(call goja.FunctionCall) goja.Value {
			reciever, err := sdk.AccAddressFromBech32(call.Argument(0).String())
			if err != nil {
				return goja.ValueFalse()
			}

			coin, err := sdk.ParseCoinNormalized(call.Argument(1).String())
			if err != nil {
				return goja.ValueFalse()
			}

			err = i.keeper.(BankKeeper).SendCoinsFromAccountToModule(ctx, creator, types.ModuleName, sdk.NewCoins(coin))
			if err != nil {
				return goja.ValueFalse()
			}
			err = i.keeper.(BankKeeper).SendCoinsFromModuleToAccount(ctx, types.ModuleName, reciever, sdk.NewCoins(coin))
			if err != nil {
				return goja.ValueFalse()
			}
			return goja.ValueTrue()
		})
		if err != nil {
			ctx.Logger().Error(err.Error())
			return err
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

			err = i.keeper.(BankKeeper).SendCoinsFromAccountToModule(ctx, contractAddress, types.ModuleName, sdk.NewCoins(coin))
			if err != nil {
				return goja.ValueFalse()
			}
			err = i.keeper.(BankKeeper).SendCoinsFromModuleToAccount(ctx, types.ModuleName, reciever, sdk.NewCoins(coin))
			if err != nil {
				return goja.ValueFalse()
			}
			return goja.ValueTrue()
		})
		if err != nil {
			ctx.Logger().Error(err.Error())
			return err
		}

		err = vm.Set("mint", func(call goja.FunctionCall) goja.Value {
			amtString := call.Argument(0).String()

			amt, ok := sdk.NewIntFromString(amtString)
			if !ok {
				return goja.ValueFalse()
			}

			coins := sdk.NewCoins(sdk.NewCoin(fmt.Sprintf("vesta-%s", contractName), amt))

			err = i.keeper.(BankKeeper).MintCoins(ctx, types.ModuleName, coins)
			if err != nil {
				return goja.ValueFalse()
			}

			err = i.keeper.(BankKeeper).SendCoinsFromModuleToAccount(ctx, types.ModuleName, contractAddress, coins)
			if err != nil {
				return goja.ValueFalse()
			}

			return goja.ValueTrue()
		})
		if err != nil {
			ctx.Logger().Error(err.Error())
			return err
		}

		err = vm.Set("burn", func(call goja.FunctionCall) goja.Value {
			amtString := call.Argument(0).String()

			amt, ok := sdk.NewIntFromString(amtString)
			if !ok {
				return goja.ValueFalse()
			}

			coins := sdk.NewCoins(sdk.NewCoin(fmt.Sprintf("vesta-%s", contractName), amt))

			err = i.keeper.(BankKeeper).SendCoinsFromAccountToModule(ctx, contractAddress, types.ModuleName, coins)
			if err != nil {
				return goja.ValueFalse()
			}

			err = i.keeper.(BankKeeper).BurnCoins(ctx, types.ModuleName, coins)
			if err != nil {
				return goja.ValueFalse()
			}

			return goja.ValueTrue()
		})
		if err != nil {
			ctx.Logger().Error(err.Error())
			return err
		}

	}

	err := vm.Set("token", gj.ToValue(fmt.Sprintf("vesta-%s", contractName)))
	if err != nil {
		ctx.Logger().Error(err.Error())
		return err
	}

	err = vm.Set("balance", func(call goja.FunctionCall) goja.Value {
		reciever, err := sdk.AccAddressFromBech32(call.Argument(0).String())
		if err != nil {
			return goja.Undefined()
		}

		denom := call.Argument(1).String()

		ctx.Logger().Error(denom)

		bal := i.keeper.(BankKeeper).GetBalance(ctx, reciever, denom)

		val := gj.ToValue(bal)

		return val
	})
	if err != nil {
		ctx.Logger().Error(err.Error())
		return err
	}

	return nil
}
