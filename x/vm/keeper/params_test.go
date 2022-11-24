package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "vesta/testutil/keeper"
	"vesta/x/vm/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.VmKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
