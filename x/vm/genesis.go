package vm

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"vesta/x/vm/keeper"
	"vesta/x/vm/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the contracts
	for _, elem := range genState.ContractsList {
		k.SetContracts(ctx, elem)
	}

	// Set contracts count
	k.SetContractsCount(ctx, genState.ContractsCount)
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.ContractsList = k.GetAllContracts(ctx)
	genesis.ContractsCount = k.GetContractsCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
