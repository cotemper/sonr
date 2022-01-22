package registry

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sonr-io/sonr/x/registry/keeper"
	"github.com/sonr-io/sonr/x/registry/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the whois
	for _, elem := range genState.WhoisList {
		k.SetWhois(ctx, elem)
	}
	// Set all the whatis
	for _, elem := range genState.WhatisList {
		k.SetWhatis(ctx, elem)
	}
	// Set all the thereis
	for _, elem := range genState.ThereisList {
		k.SetThereis(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.WhoisList = k.GetAllWhois(ctx)
	genesis.WhatisList = k.GetAllWhatis(ctx)
	genesis.ThereisList = k.GetAllThereis(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
