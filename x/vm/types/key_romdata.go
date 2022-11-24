package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// RomdataKeyPrefix is the prefix to retrieve all Romdata
	RomdataKeyPrefix = "Romdata/value/"
)

// RomdataKey returns the store key to retrieve a Romdata from the index fields
func RomdataKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
