package keeper

import (
	"github.com/VestaProtocol/vesta/x/vm/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetCronjobs set a specific cronjobs in the store from its index
func (k Keeper) SetCronjobs(ctx sdk.Context, cronjobs types.Cronjobs) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CronjobsKeyPrefix))
	b := k.cdc.MustMarshal(&cronjobs)
	store.Set(types.CronjobsKey(
		cronjobs.Contract,
	), b)
}

// GetCronjobs returns a cronjobs from its index
func (k Keeper) GetCronjobs(
	ctx sdk.Context,
	contract string,
) (val types.Cronjobs, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CronjobsKeyPrefix))

	b := store.Get(types.CronjobsKey(
		contract,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveCronjobs removes a cronjobs from the store
func (k Keeper) RemoveCronjobs(
	ctx sdk.Context,
	contract string,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CronjobsKeyPrefix))
	store.Delete(types.CronjobsKey(
		contract,
	))
}

// GetAllCronjobs returns all cronjobs
func (k Keeper) GetAllCronjobs(ctx sdk.Context) (list []types.Cronjobs) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CronjobsKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Cronjobs
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
