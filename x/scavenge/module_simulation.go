package scavenge

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/mehdisbys/scavenge/testutil/sample"
	scavengesimulation "github.com/mehdisbys/scavenge/x/scavenge/simulation"
	"github.com/mehdisbys/scavenge/x/scavenge/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = scavengesimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgSubmitScavenge = "op_weight_msg_submit_scavenge"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSubmitScavenge int = 100

	opWeightMsgCommitSolution = "op_weight_msg_commit_solution"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCommitSolution int = 100

	opWeightMsgRevealSolution = "op_weight_msg_reveal_solution"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRevealSolution int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	scavengeGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&scavengeGenesis)
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

	var weightMsgSubmitScavenge int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgSubmitScavenge, &weightMsgSubmitScavenge, nil,
		func(_ *rand.Rand) {
			weightMsgSubmitScavenge = defaultWeightMsgSubmitScavenge
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSubmitScavenge,
		scavengesimulation.SimulateMsgSubmitScavenge(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCommitSolution int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCommitSolution, &weightMsgCommitSolution, nil,
		func(_ *rand.Rand) {
			weightMsgCommitSolution = defaultWeightMsgCommitSolution
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCommitSolution,
		scavengesimulation.SimulateMsgCommitSolution(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgRevealSolution int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgRevealSolution, &weightMsgRevealSolution, nil,
		func(_ *rand.Rand) {
			weightMsgRevealSolution = defaultWeightMsgRevealSolution
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRevealSolution,
		scavengesimulation.SimulateMsgRevealSolution(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
