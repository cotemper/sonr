package channel_test

import (
	"testing"

	keepertest "github.com/sonr-io/sonr/testutil/keeper"
	"github.com/sonr-io/sonr/testutil/nullify"
	"github.com/sonr-io/sonr/x/channel"
	"github.com/sonr-io/sonr/x/channel/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		HowIsList: []types.HowIs{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ChannelKeeper(t)
	channel.InitGenesis(ctx, *k, genesisState)
	got := channel.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.HowIsList, got.HowIsList)
	// this line is used by starport scaffolding # genesis/test/assert
}
