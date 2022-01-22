package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sonr-io/sonr/x/registry/types"
)

// SetThereis set a specific thereis in the store from its index
func (k Keeper) SetThereis(ctx sdk.Context, thereis types.Thereis) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ThereisKeyPrefix))
	b := k.cdc.MustMarshal(&thereis)
	store.Set(types.ThereisKey(
		thereis.Index,
	), b)
}

// GetThereis returns a thereis from its index
func (k Keeper) GetThereis(
	ctx sdk.Context,
	index string,

) (val types.Thereis, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ThereisKeyPrefix))

	b := store.Get(types.ThereisKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveThereis removes a thereis from the store
func (k Keeper) RemoveThereis(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ThereisKeyPrefix))
	store.Delete(types.ThereisKey(
		index,
	))
}

// GetAllThereis returns all thereis
func (k Keeper) GetAllThereis(ctx sdk.Context) (list []types.Thereis) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ThereisKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Thereis
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
