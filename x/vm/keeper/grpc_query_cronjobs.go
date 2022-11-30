package keeper

import (
	"context"

	"github.com/VestaProtocol/vesta/x/vm/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) CronjobsAll(c context.Context, req *types.QueryAllCronjobsRequest) (*types.QueryAllCronjobsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var cronjobss []types.Cronjobs
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	cronjobsStore := prefix.NewStore(store, types.KeyPrefix(types.CronjobsKeyPrefix))

	pageRes, err := query.Paginate(cronjobsStore, req.Pagination, func(key []byte, value []byte) error {
		var cronjobs types.Cronjobs
		if err := k.cdc.Unmarshal(value, &cronjobs); err != nil {
			return err
		}

		cronjobss = append(cronjobss, cronjobs)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllCronjobsResponse{Cronjobs: cronjobss, Pagination: pageRes}, nil
}

func (k Keeper) Cronjobs(c context.Context, req *types.QueryGetCronjobsRequest) (*types.QueryGetCronjobsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetCronjobs(
		ctx,
		req.Contract,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetCronjobsResponse{Cronjobs: val}, nil
}
