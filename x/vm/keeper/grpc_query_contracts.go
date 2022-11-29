package keeper

import (
	"context"

	"github.com/VestaProtocol/vesta/x/vm/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ContractsAll(c context.Context, req *types.QueryAllContractsRequest) (*types.QueryAllContractsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var contractss []types.Contracts
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	contractsStore := prefix.NewStore(store, types.KeyPrefix(types.ContractsKey))

	pageRes, err := query.Paginate(contractsStore, req.Pagination, func(key []byte, value []byte) error {
		var contracts types.Contracts
		if err := k.cdc.Unmarshal(value, &contracts); err != nil {
			return err
		}

		contractss = append(contractss, contracts)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllContractsResponse{Contracts: contractss, Pagination: pageRes}, nil
}

func (k Keeper) Contracts(c context.Context, req *types.QueryGetContractsRequest) (*types.QueryGetContractsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	contracts, found := k.GetContracts(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetContractsResponse{Contracts: contracts}, nil
}
