package auction_test

import (
	"testing"

	keepertest "auction/testutil/keeper"
	"auction/testutil/nullify"
	"auction/x/auction"
	"auction/x/auction/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		SystemInfo: types.SystemInfo{
			AuctionId: 78,
		},
		AuctionList: []types.Auction{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.AuctionKeeper(t)
	auction.InitGenesis(ctx, *k, genesisState)
	got := auction.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.SystemInfo, got.SystemInfo)
	require.ElementsMatch(t, genesisState.AuctionList, got.AuctionList)
	// this line is used by starport scaffolding # genesis/test/assert
}
