package hash

import (
	"github.com/spaolacci/murmur3"
)

// MurmurHash3 computes a 64-bit MurmurHash3 hash for the given string and seed.
func MurmurHash3(str string, seed uint32) uint64 {
	hasher := murmur3.New128WithSeed(seed)
	hasher.Write([]byte(str))

	h1, _ := hasher.Sum128()
	return h1 // returns only the first 64 bits
}
