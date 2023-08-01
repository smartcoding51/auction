package keeper_test

import (
	"strconv"
	"testing"

	keepertest "auction/testutil/keeper"
	"auction/testutil/nullify"
	"auction/x/auction/keeper"
	"auction/x/auction/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNAuction(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Auction {
	items := make([]types.Auction, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetAuction(ctx, items[i])
	}
	return items
}

func TestAuctionGet(t *testing.T) {
	keeper, ctx := keepertest.AuctionKeeper(t)
	items := createNAuction(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetAuction(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestAuctionRemove(t *testing.T) {
	keeper, ctx := keepertest.AuctionKeeper(t)
	items := createNAuction(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveAuction(ctx,
			item.Index,
		)
		_, found := keeper.GetAuction(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestAuctionGetAll(t *testing.T) {
	keeper, ctx := keepertest.AuctionKeeper(t)
	items := createNAuction(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllAuction(ctx)),
	)
}
