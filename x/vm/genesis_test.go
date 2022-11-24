package vm_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "vesta/testutil/keeper"
	"vesta/testutil/nullify"
	"vesta/x/vm"
	"vesta/x/vm/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		ContractsList: []types.Contracts{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		ContractsCount: 2,
		ProgramList: []types.Program{
			{
				Name: "0",
			},
			{
				Name: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.VmKeeper(t)
	vm.InitGenesis(ctx, *k, genesisState)
	got := vm.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.ContractsList, got.ContractsList)
	require.Equal(t, genesisState.ContractsCount, got.ContractsCount)
	require.ElementsMatch(t, genesisState.ProgramList, got.ProgramList)
	// this line is used by starport scaffolding # genesis/test/assert
}
