package keeper

import (
	"context"

	"auction/x/auction/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) AuctionAll(goCtx context.Context, req *types.QueryAllAuctionRequest) (*types.QueryAllAuctionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var auctions []types.Auction
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	auctionStore := prefix.NewStore(store, types.KeyPrefix(types.AuctionKeyPrefix))

	pageRes, err := query.Paginate(auctionStore, req.Pagination, func(key []byte, value []byte) error {
		var auction types.Auction
		if err := k.cdc.Unmarshal(value, &auction); err != nil {
			return err
		}

		auctions = append(auctions, auction)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllAuctionResponse{Auction: auctions, Pagination: pageRes}, nil
}

func (k Keeper) Auction(goCtx context.Context, req *types.QueryGetAuctionRequest) (*types.QueryGetAuctionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetAuction(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetAuctionResponse{Auction: val}, nil
}
