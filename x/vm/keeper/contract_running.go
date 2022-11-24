package keeper

import (
	"vesta/x/vm/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/robertkrimen/otto"
)

func (k msgServer) applyStandardLib(ctx sdk.Context, creator sdk.AccAddress, vm *otto.Otto) {
	vm.Set("sendTokens", func(call otto.FunctionCall) otto.Value {
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
}

func (k msgServer) buildContract(ctx sdk.Context, sourceCode string, entry string, creator sdk.AccAddress) (string, error) {
	k.Logger(ctx).Info(sourceCode)

	vm := otto.New()

	vm.Set("sender", creator.String())

	k.applyStandardLib(ctx, creator, vm)

	_, err := vm.Run(sourceCode)
	if err != nil {
		return "", err
	}

	r, err := vm.Call(entry, nil)
	if err != nil {
		return "", err
	}

	return r.String(), nil
}
