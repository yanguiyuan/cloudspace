package bloomfilter

import (
	"hash/fnv"
)

type BloomFilter struct {
	bitSet  []bool
	hashers []Hasher
}

type Hasher interface {
	hash(data []byte) uint32
}

type fnvHasher struct {
}

func (h *fnvHasher) hash(data []byte) uint32 {
	hash := fnv.New32()
	hash.Write(data)
	return hash.Sum32()
}

func New(size int, numHashers int) *BloomFilter {
	return &BloomFilter{
		bitSet:  make([]bool, size),
		hashers: generateHashers(numHashers),
	}
}

func (bf *BloomFilter) Add(data []byte) {
	for _, h := range bf.hashers {
		index := h.hash(data) % uint32(len(bf.bitSet))
		bf.bitSet[index] = true
	}
}

func (bf *BloomFilter) Contains(data []byte) bool {
	for _, h := range bf.hashers {
		index := h.hash(data) % uint32(len(bf.bitSet))
		if !bf.bitSet[index] {
			return false
		}
	}
	return true
}

func generateHashers(count int) []Hasher {
	hashers := make([]Hasher, count)
	for i := 0; i < count; i++ {
		hashers[i] = &fnvHasher{}
	}
	return hashers
}
