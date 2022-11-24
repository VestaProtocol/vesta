package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"vesta/x/vm/types"
)

// SetRomdata set a specific romdata in the store from its index
func (k Keeper) SetRomdata(ctx sdk.Context, romdata types.Romdata) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RomdataKeyPrefix))
	b := k.cdc.MustMarshal(&romdata)
	store.Set(types.RomdataKey(
		romdata.Index,
	), b)
}

// GetRomdata returns a romdata from its index
func (k Keeper) GetRomdata(
	ctx sdk.Context,
	index string,

) (val types.Romdata, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RomdataKeyPrefix))

	b := store.Get(types.RomdataKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveRomdata removes a romdata from the store
func (k Keeper) RemoveRomdata(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RomdataKeyPrefix))
	store.Delete(types.RomdataKey(
		index,
	))
}

// GetAllRomdata returns all romdata
func (k Keeper) GetAllRomdata(ctx sdk.Context) (list []types.Romdata) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RomdataKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Romdata
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
