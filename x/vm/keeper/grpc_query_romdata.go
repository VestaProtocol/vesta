package keeper

import (
	"context"

	"github.com/TheMarstonConnell/vesta/x/vm/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) RomdataAll(c context.Context, req *types.QueryAllRomdataRequest) (*types.QueryAllRomdataResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var romdatas []types.Romdata
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	romdataStore := prefix.NewStore(store, types.KeyPrefix(types.RomdataKeyPrefix))

	pageRes, err := query.Paginate(romdataStore, req.Pagination, func(key []byte, value []byte) error {
		var romdata types.Romdata
		if err := k.cdc.Unmarshal(value, &romdata); err != nil {
			return err
		}

		romdatas = append(romdatas, romdata)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllRomdataResponse{Romdata: romdatas, Pagination: pageRes}, nil
}

func (k Keeper) Romdata(c context.Context, req *types.QueryGetRomdataRequest) (*types.QueryGetRomdataResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetRomdata(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetRomdataResponse{Romdata: val}, nil
}
