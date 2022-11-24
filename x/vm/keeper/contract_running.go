package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/robertkrimen/otto"
	"github.com/tendermint/tendermint/crypto"
)

func NewContractAddress(name string) sdk.AccAddress {
	return sdk.AccAddress(crypto.AddressHash([]byte(name)))
}

func (k msgServer) newContractAccount(ctx sdk.Context, address sdk.AccAddress) authtypes.AccountI {
	baseAcc := k.accountKeeper.NewAccountWithAddress(ctx, address)

	return baseAcc
}

func (k msgServer) getContractAddress(ctx sdk.Context, contractName string) (sdk.AccAddress, error) {
	address := NewContractAddress(contractName)

	acc := k.accountKeeper.GetAccount(ctx, address)
	if acc != nil {
		return acc.GetAddress(), nil
	}

	acc = k.newContractAccount(ctx, address)

	k.accountKeeper.SetAccount(ctx, acc)

	return acc.GetAddress(), nil
}

func (k msgServer) buildContract(ctx sdk.Context, name string, sourceCode string, entry string, creator sdk.AccAddress) (string, error) {
	k.Logger(ctx).Info(sourceCode)

	vm := otto.New()

	address, err := k.getContractAddress(ctx, name)
	if err != nil {
		return "", err
	}

	k.applyStandardLib(ctx, creator, address, vm)

	_, err = vm.Run(sourceCode)
	if err != nil {
		return "", err
	}

	r, err := vm.Get("contractFunctions")
	if err != nil {
		return "", err
	}

	function, err := r.Object().Get(entry)
	if err != nil {
		return "", err
	}

	res, err := function.Call(function, nil)
	if err != nil {
		return "", err
	}

	return res.String(), nil
}
