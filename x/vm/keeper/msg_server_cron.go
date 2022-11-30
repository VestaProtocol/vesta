package keeper

import (
	"context"

	"github.com/VestaProtocol/vesta/x/vm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) Cron(goCtx context.Context, msg *types.MsgCron) (*types.MsgCronResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	i, ok := sdk.NewIntFromString(msg.Interval)
	if !ok {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "cannot parse interval to int")
	}

	program, found := k.GetProgram(ctx, msg.Contract)
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrKeyNotFound, "cannot find contract")
	}

	if program.Creator != msg.Creator {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "not your contract")
	}

	crn := types.Cronjobs{
		Contract: msg.Contract,
		Function: msg.Function,
		Interval: i.Int64(),
	}

	k.SetCronjobs(ctx, crn)

	return &types.MsgCronResponse{}, nil
}
