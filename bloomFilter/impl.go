package bloom

import (
	"hash"

	"github.com/twmb/murmur3"
)

type MurMur3Hasher struct {
}

func NewMurMur3Hasher() *MurMur3Hasher {
	return &MurMur3Hasher{}
}

func (h *MurMur3Hasher) GetHashes(n uint64) []hash.Hash64 {
	hashers := make([]hash.Hash64, n)
	for i := 0; uint64(i) < n; i++ {
		hashers[i] = murmur3.SeedNew64(uint64(i))
	}
	return hashers
}
