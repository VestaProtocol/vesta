package keeper_test

import (
	"testing"

	testkeeper "github.com/VestaProtocol/vesta/testutil/keeper"
	"github.com/VestaProtocol/vesta/x/funding/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.FundingKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
	require.EqualValues(t, params.ValidTokens, k.ValidTokens(ctx))
}
