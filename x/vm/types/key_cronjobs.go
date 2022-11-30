package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// CronjobsKeyPrefix is the prefix to retrieve all Cronjobs
	CronjobsKeyPrefix = "Cronjobs/value/"
)

// CronjobsKey returns the store key to retrieve a Cronjobs from the index fields
func CronjobsKey(
	contract string,
) []byte {
	var key []byte

	contractBytes := []byte(contract)
	key = append(key, contractBytes...)
	key = append(key, []byte("/")...)

	return key
}
