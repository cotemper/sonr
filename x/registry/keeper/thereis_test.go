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

func createNThereis(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Thereis {
	items := make([]types.Thereis, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetThereis(ctx, items[i])
	}
	return items
}

func TestThereisGet(t *testing.T) {
	keeper, ctx := keepertest.RegistryKeeper(t)
	items := createNThereis(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetThereis(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestThereisRemove(t *testing.T) {
	keeper, ctx := keepertest.RegistryKeeper(t)
	items := createNThereis(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveThereis(ctx,
			item.Index,
		)
		_, found := keeper.GetThereis(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestThereisGetAll(t *testing.T) {
	keeper, ctx := keepertest.RegistryKeeper(t)
	items := createNThereis(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllThereis(ctx)),
	)
}
