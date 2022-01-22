package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sonr-io/sonr/x/object/types"
)

// func (k Keeper) AppendPost() uint64 {
// 	 count := k.GetPostCount()
// 	 store.Set()
// 	 k.SetPostCount()
// 	 return count
// }

func (k Keeper) GetPostCount(ctx sdk.Context) uint64 {
	// Get the store using storeKey (which is "blog") and PostCountKey (which is "Post-count-")
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.MemStoreKey))
	// Convert the PostCountKey to bytes
	byteKey := []byte(types.MemStoreKey)
	// Get the value of the count
	bz := store.Get(byteKey)
	// Return zero if the count value is not found (for example, it's the first post)
	if bz == nil {
		return 0
	}
	// Convert the count into a uint64
	return binary.BigEndian.Uint64(bz)
}
