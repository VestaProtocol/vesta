package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

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

	_, found = k.GetContracts(ctx, cd.Uint64())
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrConflict, "no contract with that code exists")
	}

	address, err := k.getContractAddress(ctx, msg.Name)
	if err != nil {
		return nil, err
	}

	p := types.Program{
		Name:    msg.Name,
		Creator: msg.Creator,
		Code:    []uint64{cd.Uint64()},
		Address: address.String(),
	}

	k.SetProgram(ctx, p)

	return &types.MsgInstantiateResponse{}, nil
}
