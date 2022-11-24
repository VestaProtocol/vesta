package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"vesta/x/vm/types"
)

// SetProgram set a specific program in the store from its index
func (k Keeper) SetProgram(ctx sdk.Context, program types.Program) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProgramKeyPrefix))
	b := k.cdc.MustMarshal(&program)
	store.Set(types.ProgramKey(
		program.Name,
	), b)
}

// GetProgram returns a program from its index
func (k Keeper) GetProgram(
	ctx sdk.Context,
	name string,

) (val types.Program, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProgramKeyPrefix))

	b := store.Get(types.ProgramKey(
		name,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveProgram removes a program from the store
func (k Keeper) RemoveProgram(
	ctx sdk.Context,
	name string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProgramKeyPrefix))
	store.Delete(types.ProgramKey(
		name,
	))
}

// GetAllProgram returns all program
func (k Keeper) GetAllProgram(ctx sdk.Context) (list []types.Program) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProgramKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Program
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
