package keeper_test

import (
	"strconv"
	"testing"

	keepertest "github.com/VestaProtocol/vesta/testutil/keeper"
	"github.com/VestaProtocol/vesta/testutil/nullify"
	"github.com/VestaProtocol/vesta/x/vm/keeper"
	"github.com/VestaProtocol/vesta/x/vm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNCronjobs(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Cronjobs {
	items := make([]types.Cronjobs, n)
	for i := range items {
		items[i].Contract = strconv.Itoa(i)

		keeper.SetCronjobs(ctx, items[i])
	}
	return items
}

func TestCronjobsGet(t *testing.T) {
	keeper, ctx := keepertest.VmKeeper(t)
	items := createNCronjobs(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetCronjobs(ctx,
			item.Contract,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}

func TestCronjobsRemove(t *testing.T) {
	keeper, ctx := keepertest.VmKeeper(t)
	items := createNCronjobs(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveCronjobs(ctx,
			item.Contract,
		)
		_, found := keeper.GetCronjobs(ctx,
			item.Contract,
		)
		require.False(t, found)
	}
}

func TestCronjobsGetAll(t *testing.T) {
	keeper, ctx := keepertest.VmKeeper(t)
	items := createNCronjobs(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllCronjobs(ctx)),
	)
}
