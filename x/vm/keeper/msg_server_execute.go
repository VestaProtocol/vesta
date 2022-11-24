package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"vesta/x/vm/types"
)

func (k msgServer) Execute(goCtx context.Context, msg *types.MsgExecute) (*types.MsgExecuteResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	program, found := k.GetProgram(ctx, msg.Contract)
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrKeyNotFound, "cannot find contract '%s'", msg.Contract)
	}

	source, ok := k.GetContracts(ctx, program.Code)
	if !ok {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrKeyNotFound, "cannot find contract source")
	}

	code := source.Source

	address, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	val, err := k.buildContract(ctx, code, msg.Function, address)
	if err != nil {
		return nil, err
	}

	return &types.MsgExecuteResponse{Console: val}, nil
}
