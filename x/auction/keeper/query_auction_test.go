package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "auction/testutil/keeper"
	"auction/testutil/nullify"
	"auction/x/auction/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestAuctionQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.AuctionKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNAuction(keeper, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetAuctionRequest
		response *types.QueryGetAuctionResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetAuctionRequest{
				Index: msgs[0].Index,
			},
			response: &types.QueryGetAuctionResponse{Auction: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetAuctionRequest{
				Index: msgs[1].Index,
			},
			response: &types.QueryGetAuctionResponse{Auction: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetAuctionRequest{
				Index: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Auction(wctx, tc.request)
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

func TestAuctionQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.AuctionKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNAuction(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllAuctionRequest {
		return &types.QueryAllAuctionRequest{
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
			resp, err := keeper.AuctionAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Auction), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Auction),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.AuctionAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Auction), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Auction),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.AuctionAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.Auction),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.AuctionAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
