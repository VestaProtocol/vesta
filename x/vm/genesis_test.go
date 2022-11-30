package vm_test

import (
	"testing"

	keepertest "github.com/VestaProtocol/vesta/testutil/keeper"
	"github.com/VestaProtocol/vesta/testutil/nullify"
	"github.com/VestaProtocol/vesta/x/vm"
	"github.com/VestaProtocol/vesta/x/vm/types"
	"github.com/stretchr/testify/require"
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
		RomdataList: []types.Romdata{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		CronjobsList: []types.Cronjobs{
			{
				Contract: "0",
			},
			{
				Contract: "1",
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
	require.ElementsMatch(t, genesisState.RomdataList, got.RomdataList)
	require.ElementsMatch(t, genesisState.CronjobsList, got.CronjobsList)
	require.ElementsMatch(t, genesisState.CronjobsList, got.CronjobsList)
	// this line is used by starport scaffolding # genesis/test/assert
}
