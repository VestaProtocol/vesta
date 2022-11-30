package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/VestaProtocol/vesta/testutil/keeper"
	"github.com/VestaProtocol/vesta/testutil/nullify"
	"github.com/VestaProtocol/vesta/x/vm/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestCronjobsQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.VmKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNCronjobs(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetCronjobsRequest
		response *types.QueryGetCronjobsResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetCronjobsRequest{
				Contract: msgs[0].Contract,
			},
			response: &types.QueryGetCronjobsResponse{Cronjobs: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetCronjobsRequest{
				Contract: msgs[1].Contract,
			},
			response: &types.QueryGetCronjobsResponse{Cronjobs: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetCronjobsRequest{
				Contract: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Cronjobs(wctx, tc.request)
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

func TestCronjobsQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.VmKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNCronjobs(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllCronjobsRequest {
		return &types.QueryAllCronjobsRequest{
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
			resp, err := keeper.CronjobsAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Cronjobs), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Cronjobs),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.CronjobsAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Cronjobs), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Cronjobs),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.CronjobsAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.Cronjobs),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.CronjobsAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
