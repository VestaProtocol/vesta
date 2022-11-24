package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// ProgramKeyPrefix is the prefix to retrieve all Program
	ProgramKeyPrefix = "Program/value/"
)

// ProgramKey returns the store key to retrieve a Program from the index fields
func ProgramKey(
	name string,
) []byte {
	var key []byte

	nameBytes := []byte(name)
	key = append(key, nameBytes...)
	key = append(key, []byte("/")...)

	return key
}
