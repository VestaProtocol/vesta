package keeper

import (
	"encoding/binary"

	"github.com/TheMarstonConnell/vesta/x/vm/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetContractsCount get the total number of contracts
func (k Keeper) GetContractsCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.ContractsCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetContractsCount set the total number of contracts
func (k Keeper) SetContractsCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.ContractsCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendContracts appends a contracts in the store with a new id and update the count
func (k Keeper) AppendContracts(
	ctx sdk.Context,
	contracts types.Contracts,
) uint64 {
	// Create the contracts
	count := k.GetContractsCount(ctx)

	// Set the ID of the appended value
	contracts.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ContractsKey))
	appendedValue := k.cdc.MustMarshal(&contracts)
	store.Set(GetContractsIDBytes(contracts.Id), appendedValue)

	// Update contracts count
	k.SetContractsCount(ctx, count+1)

	return count
}

// SetContracts set a specific contracts in the store
func (k Keeper) SetContracts(ctx sdk.Context, contracts types.Contracts) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ContractsKey))
	b := k.cdc.MustMarshal(&contracts)
	store.Set(GetContractsIDBytes(contracts.Id), b)
}

// GetContracts returns a contracts from its id
func (k Keeper) GetContracts(ctx sdk.Context, id uint64) (val types.Contracts, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ContractsKey))
	b := store.Get(GetContractsIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveContracts removes a contracts from the store
func (k Keeper) RemoveContracts(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ContractsKey))
	store.Delete(GetContractsIDBytes(id))
}

// GetAllContracts returns all contracts
func (k Keeper) GetAllContracts(ctx sdk.Context) (list []types.Contracts) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ContractsKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Contracts
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetContractsIDBytes returns the byte representation of the ID
func GetContractsIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetContractsIDFromBytes returns ID in uint64 format from a byte array
func GetContractsIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
