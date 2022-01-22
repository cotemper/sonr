package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/sonr-io/sonr/testutil/keeper"
	"github.com/sonr-io/sonr/testutil/nullify"
	"github.com/sonr-io/sonr/x/registry/keeper"
	"github.com/sonr-io/sonr/x/registry/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNWhatis(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Whatis {
	items := make([]types.Whatis, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetWhatis(ctx, items[i])
	}
	return items
}

func TestWhatisGet(t *testing.T) {
	keeper, ctx := keepertest.RegistryKeeper(t)
	items := createNWhatis(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetWhatis(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestWhatisRemove(t *testing.T) {
	keeper, ctx := keepertest.RegistryKeeper(t)
	items := createNWhatis(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveWhatis(ctx,
			item.Index,
		)
		_, found := keeper.GetWhatis(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestWhatisGetAll(t *testing.T) {
	keeper, ctx := keepertest.RegistryKeeper(t)
	items := createNWhatis(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllWhatis(ctx)),
	)
}
