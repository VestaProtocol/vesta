package keeper_test

import (
	"testing"

	keepertest "github.com/VestaProtocol/vesta/testutil/keeper"
	"github.com/VestaProtocol/vesta/testutil/nullify"
	"github.com/VestaProtocol/vesta/x/vm/keeper"
	"github.com/VestaProtocol/vesta/x/vm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func createNContracts(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Contracts {
	items := make([]types.Contracts, n)
	for i := range items {
		items[i].Id = keeper.AppendContracts(ctx, items[i])
	}
	return items
}

func TestContractsGet(t *testing.T) {
	keeper, ctx := keepertest.VmKeeper(t)
	items := createNContracts(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetContracts(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestContractsRemove(t *testing.T) {
	keeper, ctx := keepertest.VmKeeper(t)
	items := createNContracts(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveContracts(ctx, item.Id)
		_, found := keeper.GetContracts(ctx, item.Id)
		require.False(t, found)
	}
}

func TestContractsGetAll(t *testing.T) {
	keeper, ctx := keepertest.VmKeeper(t)
	items := createNContracts(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllContracts(ctx)),
	)
}

func TestContractsCount(t *testing.T) {
	keeper, ctx := keepertest.VmKeeper(t)
	items := createNContracts(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetContractsCount(ctx))
}
