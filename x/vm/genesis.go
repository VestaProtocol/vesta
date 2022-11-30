package vm

import (
	"github.com/VestaProtocol/vesta/x/vm/keeper"
	"github.com/VestaProtocol/vesta/x/vm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the contracts
	for _, elem := range genState.ContractsList {
		k.SetContracts(ctx, elem)
	}

	// Set contracts count
	k.SetContractsCount(ctx, genState.ContractsCount)
	// Set all the program
	for _, elem := range genState.ProgramList {
		k.SetProgram(ctx, elem)
	}
	// Set all the romdata
	for _, elem := range genState.RomdataList {
		k.SetRomdata(ctx, elem)
	}
	// Set all the cronjobs
	for _, elem := range genState.CronjobsList {
		k.SetCronjobs(ctx, elem)
	}
	// Set all the cronjobs
	for _, elem := range genState.CronjobsList {
		k.SetCronjobs(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.ContractsList = k.GetAllContracts(ctx)
	genesis.ContractsCount = k.GetContractsCount(ctx)
	genesis.ProgramList = k.GetAllProgram(ctx)
	genesis.RomdataList = k.GetAllRomdata(ctx)
	genesis.CronjobsList = k.GetAllCronjobs(ctx)
	genesis.CronjobsList = k.GetAllCronjobs(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
