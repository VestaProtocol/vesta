package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"vesta/x/vm/types"
)

func (k msgServer) Execute(goCtx context.Context, msg *types.MsgExecute) (*types.MsgExecuteResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	contractId, ok := sdk.NewIntFromString(msg.Contract)
	if !ok {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "cannot convert contract id to int.")
	}

	source, ok := k.GetContracts(ctx, contractId.Uint64())
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
