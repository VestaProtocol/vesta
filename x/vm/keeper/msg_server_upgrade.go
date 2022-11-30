package keeper

import (
	"context"

	"github.com/VestaProtocol/vesta/x/vm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) Upgrade(goCtx context.Context, msg *types.MsgUpgrade) (*types.MsgUpgradeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	program, found := k.GetProgram(ctx, msg.Contract)
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrKeyNotFound, "a program with that name cannot be found")
	}

	cd, ok := sdk.NewIntFromString(msg.Code)
	if !ok {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "cannot parse code as int")
	}

	_, found = k.GetContracts(ctx, cd.Uint64())
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrConflict, "no contract with that code exists")
	}

	if program.Creator != msg.Creator {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "you do not own this contract")
	}

	program.Code = append(program.Code, cd.Uint64())

	k.SetProgram(ctx, program)

	return &types.MsgUpgradeResponse{}, nil
}
