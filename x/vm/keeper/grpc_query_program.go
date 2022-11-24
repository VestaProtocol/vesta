package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"vesta/x/vm/types"
)

func (k Keeper) ProgramAll(c context.Context, req *types.QueryAllProgramRequest) (*types.QueryAllProgramResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var programs []types.Program
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	programStore := prefix.NewStore(store, types.KeyPrefix(types.ProgramKeyPrefix))

	pageRes, err := query.Paginate(programStore, req.Pagination, func(key []byte, value []byte) error {
		var program types.Program
		if err := k.cdc.Unmarshal(value, &program); err != nil {
			return err
		}

		programs = append(programs, program)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllProgramResponse{Program: programs, Pagination: pageRes}, nil
}

func (k Keeper) Program(c context.Context, req *types.QueryGetProgramRequest) (*types.QueryGetProgramResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetProgram(
		ctx,
		req.Name,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetProgramResponse{Program: val}, nil
}
