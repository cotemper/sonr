package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// ThereisKeyPrefix is the prefix to retrieve all Thereis
	ThereisKeyPrefix = "Thereis/value/"
)

// ThereisKey returns the store key to retrieve a Thereis from the index fields
func ThereisKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
