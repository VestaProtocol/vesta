package keeper

import (
	"context"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/dop251/goja"

	"github.com/VestaProtocol/vesta/x/vm/types"
)

func (k msgServer) Instantiate(goCtx context.Context, msg *types.MsgInstantiate) (*types.MsgInstantiateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, found := k.GetProgram(ctx, msg.Name)
	if found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrConflict, "a program with that name already exists")
	}

	cd, ok := sdk.NewIntFromString(msg.Code)
	if !ok {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "cannot parse code as int")
	}

	code, found := k.GetContracts(ctx, cd.Uint64())
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrConflict, "no contract with that code exists")
	}

	contractAddress, err := k.getContractAddress(ctx, msg.Name)
	if err != nil {
		return nil, err
	}

	p := types.Program{
		Name:    msg.Name,
		Creator: msg.Creator,
		Code:    []uint64{cd.Uint64()},
		Address: contractAddress.String(),
	}

	k.SetProgram(ctx, p)

	address, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	argsString := strings.Split(msg.Args, ",")
	vals := make([]goja.Value, len(argsString))
	for i, s := range argsString {
		vals[i] = goja.ValueString(s)
	}

	_, err = k.initContract(ctx, msg.Name, code.Source, address, vals)
	if err != nil {
		ctx.Logger().Error(err.Error())
		return nil, err
	}

	return &types.MsgInstantiateResponse{}, nil
}
