package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "vesta/testutil/keeper"
	"vesta/testutil/nullify"
	"vesta/x/vm/keeper"
	"vesta/x/vm/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNRomdata(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Romdata {
	items := make([]types.Romdata, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetRomdata(ctx, items[i])
	}
	return items
}

func TestRomdataGet(t *testing.T) {
	keeper, ctx := keepertest.VmKeeper(t)
	items := createNRomdata(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetRomdata(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}

func TestRomdataRemove(t *testing.T) {
	keeper, ctx := keepertest.VmKeeper(t)
	items := createNRomdata(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveRomdata(ctx,
			item.Index,
		)
		_, found := keeper.GetRomdata(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestRomdataGetAll(t *testing.T) {
	keeper, ctx := keepertest.VmKeeper(t)
	items := createNRomdata(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllRomdata(ctx)),
	)
}
