package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// AuctionKeyPrefix is the prefix to retrieve all Auction
	AuctionKeyPrefix = "Auction/value/"
)

// AuctionKey returns the store key to retrieve a Auction from the index fields
func AuctionKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
