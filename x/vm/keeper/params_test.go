package keeper_test

import (
	"testing"

	testkeeper "github.com/TheMarstonConnell/vesta/testutil/keeper"
	"github.com/TheMarstonConnell/vesta/x/vm/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.VmKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
