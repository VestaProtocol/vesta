package keeper_test

import (
	"strconv"
	"testing"

	keepertest "github.com/TheMarstonConnell/vesta/testutil/keeper"
	"github.com/TheMarstonConnell/vesta/testutil/nullify"
	"github.com/TheMarstonConnell/vesta/x/vm/keeper"
	"github.com/TheMarstonConnell/vesta/x/vm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNProgram(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Program {
	items := make([]types.Program, n)
	for i := range items {
		items[i].Name = strconv.Itoa(i)

		keeper.SetProgram(ctx, items[i])
	}
	return items
}

func TestProgramGet(t *testing.T) {
	keeper, ctx := keepertest.VmKeeper(t)
	items := createNProgram(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetProgram(ctx,
			item.Name,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}

func TestProgramRemove(t *testing.T) {
	keeper, ctx := keepertest.VmKeeper(t)
	items := createNProgram(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveProgram(ctx,
			item.Name,
		)
		_, found := keeper.GetProgram(ctx,
			item.Name,
		)
		require.False(t, found)
	}
}

func TestProgramGetAll(t *testing.T) {
	keeper, ctx := keepertest.VmKeeper(t)
	items := createNProgram(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllProgram(ctx)),
	)
}
