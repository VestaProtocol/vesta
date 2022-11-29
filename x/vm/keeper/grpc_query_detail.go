package keeper

import (
	"context"
	"strings"

	"github.com/VestaProtocol/vesta/x/vm/types"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/dop251/goja"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Detail(goCtx context.Context, req *types.QueryDetailRequest) (*types.QueryDetailResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	program, found := k.GetProgram(ctx, req.Name)
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrKeyNotFound, "cannot find contract '%s'", req.Name)
	}

	source, ok := k.GetContracts(ctx, program.Code)
	if !ok {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrKeyNotFound, "cannot find contract source")
	}

	code := source.Source

	argsString := strings.Split(req.Args, ",")
	vals := make([]goja.Value, len(argsString))
	for i, s := range argsString {
		vals[i] = goja.ValueString(s)
	}

	val, err := k.queryContract(ctx, req.Name, code, req.Query, vals)
	if err != nil {
		return nil, err
	}

	return &types.QueryDetailResponse{Response: val}, nil
}
