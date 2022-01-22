package registry

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/sonr-io/sonr/testutil/sample"
	registrysimulation "github.com/sonr-io/sonr/x/registry/simulation"
	"github.com/sonr-io/sonr/x/registry/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = registrysimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgRegisterService = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRegisterService int = 100

	opWeightMsgRegisterName = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRegisterName int = 100

	opWeightMsgAccessName = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgAccessName int = 100

	opWeightMsgUpdateName = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateName int = 100

	opWeightMsgAccessService = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgAccessService int = 100

	opWeightMsgUpdateService = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateService int = 100

	opWeightMsgCreateWhois = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateWhois int = 100

	opWeightMsgUpdateWhois = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateWhois int = 100

	opWeightMsgDeleteWhois = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteWhois int = 100

	opWeightMsgCreateWhatis = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateWhatis int = 100

	opWeightMsgUpdateWhatis = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateWhatis int = 100

	opWeightMsgDeleteWhatis = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteWhatis int = 100

	opWeightMsgCreateThereis = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateThereis int = 100

	opWeightMsgUpdateThereis = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateThereis int = 100

	opWeightMsgDeleteThereis = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteThereis int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	registryGenesis := types.GenesisState{
		WhoisList: []types.Whois{
			{
				Creator: sample.AccAddress(),
				Index:   "0",
			},
			{
				Creator: sample.AccAddress(),
				Index:   "1",
			},
		},
		WhatisList: []types.Whatis{
			{
				Creator: sample.AccAddress(),
				Index:   "0",
			},
			{
				Creator: sample.AccAddress(),
				Index:   "1",
			},
		},
		ThereisList: []types.Thereis{
			{
				Creator: sample.AccAddress(),
				Index:   "0",
			},
			{
				Creator: sample.AccAddress(),
				Index:   "1",
			},
		},
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&registryGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgRegisterService int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgRegisterService, &weightMsgRegisterService, nil,
		func(_ *rand.Rand) {
			weightMsgRegisterService = defaultWeightMsgRegisterService
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRegisterService,
		registrysimulation.SimulateMsgRegisterService(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgRegisterName int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgRegisterName, &weightMsgRegisterName, nil,
		func(_ *rand.Rand) {
			weightMsgRegisterName = defaultWeightMsgRegisterName
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRegisterName,
		registrysimulation.SimulateMsgRegisterName(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgAccessName int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgAccessName, &weightMsgAccessName, nil,
		func(_ *rand.Rand) {
			weightMsgAccessName = defaultWeightMsgAccessName
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgAccessName,
		registrysimulation.SimulateMsgAccessName(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateName int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateName, &weightMsgUpdateName, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateName = defaultWeightMsgUpdateName
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateName,
		registrysimulation.SimulateMsgUpdateName(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgAccessService int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgAccessService, &weightMsgAccessService, nil,
		func(_ *rand.Rand) {
			weightMsgAccessService = defaultWeightMsgAccessService
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgAccessService,
		registrysimulation.SimulateMsgAccessService(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateService int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateService, &weightMsgUpdateService, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateService = defaultWeightMsgUpdateService
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateService,
		registrysimulation.SimulateMsgUpdateService(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateWhois int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateWhois, &weightMsgCreateWhois, nil,
		func(_ *rand.Rand) {
			weightMsgCreateWhois = defaultWeightMsgCreateWhois
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateWhois,
		registrysimulation.SimulateMsgCreateWhois(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateWhois int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateWhois, &weightMsgUpdateWhois, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateWhois = defaultWeightMsgUpdateWhois
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateWhois,
		registrysimulation.SimulateMsgUpdateWhois(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteWhois int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteWhois, &weightMsgDeleteWhois, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteWhois = defaultWeightMsgDeleteWhois
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteWhois,
		registrysimulation.SimulateMsgDeleteWhois(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateWhatis int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateWhatis, &weightMsgCreateWhatis, nil,
		func(_ *rand.Rand) {
			weightMsgCreateWhatis = defaultWeightMsgCreateWhatis
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateWhatis,
		registrysimulation.SimulateMsgCreateWhatis(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateWhatis int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateWhatis, &weightMsgUpdateWhatis, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateWhatis = defaultWeightMsgUpdateWhatis
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateWhatis,
		registrysimulation.SimulateMsgUpdateWhatis(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteWhatis int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteWhatis, &weightMsgDeleteWhatis, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteWhatis = defaultWeightMsgDeleteWhatis
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteWhatis,
		registrysimulation.SimulateMsgDeleteWhatis(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateThereis int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateThereis, &weightMsgCreateThereis, nil,
		func(_ *rand.Rand) {
			weightMsgCreateThereis = defaultWeightMsgCreateThereis
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateThereis,
		registrysimulation.SimulateMsgCreateThereis(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateThereis int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateThereis, &weightMsgUpdateThereis, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateThereis = defaultWeightMsgUpdateThereis
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateThereis,
		registrysimulation.SimulateMsgUpdateThereis(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteThereis int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteThereis, &weightMsgDeleteThereis, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteThereis = defaultWeightMsgDeleteThereis
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteThereis,
		registrysimulation.SimulateMsgDeleteThereis(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
