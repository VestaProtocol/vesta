package keeper

import (
	"context"
	"strings"

	"vesta/x/vm/types"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

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

	val, err := k.queryContract(ctx, req.Name, code, req.Query, strings.Split(req.Args, ","))
	if err != nil {
		return nil, err
	}

	return &types.QueryDetailResponse{Response: val}, nil
}
