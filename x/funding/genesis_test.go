package funding_test

import (
	"testing"

	keepertest "github.com/VestaProtocol/vesta/testutil/keeper"
	"github.com/VestaProtocol/vesta/testutil/nullify"
	"github.com/VestaProtocol/vesta/x/funding"
	"github.com/VestaProtocol/vesta/x/funding/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.FundingKeeper(t)
	funding.InitGenesis(ctx, *k, genesisState)
	got := funding.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
