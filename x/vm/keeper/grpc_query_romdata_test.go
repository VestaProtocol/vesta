package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "vesta/testutil/keeper"
	"vesta/testutil/nullify"
	"vesta/x/vm/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestRomdataQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.VmKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNRomdata(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetRomdataRequest
		response *types.QueryGetRomdataResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetRomdataRequest{
				Index: msgs[0].Index,
			},
			response: &types.QueryGetRomdataResponse{Romdata: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetRomdataRequest{
				Index: msgs[1].Index,
			},
			response: &types.QueryGetRomdataResponse{Romdata: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetRomdataRequest{
				Index: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Romdata(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestRomdataQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.VmKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNRomdata(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllRomdataRequest {
		return &types.QueryAllRomdataRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.RomdataAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Romdata), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Romdata),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.RomdataAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Romdata), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Romdata),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.RomdataAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.Romdata),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.RomdataAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
