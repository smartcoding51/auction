package keeper

import (
	"auction/x/auction/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetAuction set a specific auction in the store from its index
func (k Keeper) SetAuction(ctx sdk.Context, auction types.Auction) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AuctionKeyPrefix))
	b := k.cdc.MustMarshal(&auction)
	store.Set(types.AuctionKey(
		auction.Index,
	), b)
}

// GetAuction returns a auction from its index
func (k Keeper) GetAuction(
	ctx sdk.Context,
	index string,

) (val types.Auction, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AuctionKeyPrefix))

	b := store.Get(types.AuctionKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveAuction removes a auction from the store
func (k Keeper) RemoveAuction(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AuctionKeyPrefix))
	store.Delete(types.AuctionKey(
		index,
	))
}

// GetAllAuction returns all auction
func (k Keeper) GetAllAuction(ctx sdk.Context) (list []types.Auction) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AuctionKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Auction
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
