package keeper

import (
	"context"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/dop251/goja"

	"github.com/TheMarstonConnell/vesta/x/vm/types"
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

	argsString := strings.Split(msg.Args, ",")
	vals := make([]goja.Value, len(argsString))
	for i, s := range argsString {
		vals[i] = goja.ValueString(s)
	}

	val, err := k.executeContract(ctx, msg.Contract, code, msg.Function, address, vals)
	if err != nil {
		ctx.Logger().Error(err.Error())
		return nil, err
	}

	return &types.MsgExecuteResponse{Console: val}, nil
}
