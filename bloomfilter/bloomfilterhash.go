package bloomfilter

import (
	"encoding"
	"hash/fnv"
)

type BloomFilterHash[T encoding.BinaryMarshaler] struct {
	bitset  []bool
	numHash int
}

// NewBloomFilter creates a new BloomFilter with the specified size and number of hash functions
func NewBloomFilter[T encoding.BinaryMarshaler](size int, numHash int) *BloomFilterHash[T] {
	return &BloomFilterHash[T]{
		bitset:  make([]bool, size),
		numHash: numHash,
	}
}

// Add inserts a new element into the Bloom filter
func (bf *BloomFilterHash[T]) Add(element T) {
	for i := 0; i < bf.numHash; i++ {
		index := bf.hash(element, i) % len(bf.bitset)
		bf.bitset[index] = true
	}
}

// Contains checks if the Bloom filter possibly contains the given element
func (bf *BloomFilterHash[T]) Contains(element T) bool {
	for i := 0; i < bf.numHash; i++ {
		index := bf.hash(element, i) % len(bf.bitset)
		if !bf.bitset[index] {
			return false
		}
	}
	return true
}

// hash generates a hash value for the element using the specified hash function index
func (bf *BloomFilterHash[T]) hash(element T, index int) int {
	hash := fnv.New32a()
	data, err := element.MarshalBinary()
	if err != nil {
		return 0
	}
	_, err = hash.Write(data)
	if err != nil {
		return 0
	}
	hashValue := hash.Sum32()
	return int(hashValue) + index
}
