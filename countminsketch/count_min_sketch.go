package countminsketch

import (
	"errors"
	"hash/fnv"
)

type hashFunction func(data []byte) uint32
type CountMinSketch struct {
	width  int
	depth  int
	table  [][]int
	hashFn hashFunction
}

func NewCountMinSketch(width, depth int) (*CountMinSketch, error) {
	if width <= 0 {
		return nil, errors.New("name cannot be empty")
	}
	if depth <= 0 {
		return nil, errors.New("age must be between 0 and 120")
	}
	table := make([][]int, depth)
	for i := range table {
		table[i] = make([]int, width)
	}

	return &CountMinSketch{
		width:  width,
		depth:  depth,
		table:  table,
		hashFn: hashFnv,
	}, nil
}

func (cms *CountMinSketch) Increment(item []byte) {
	for i := 0; i < cms.depth; i++ {
		hashValue := cms.hashFn(append(item, []byte(string(rune(i)))...))
		index := hashValue % uint32(cms.width)
		cms.table[i][index]++
	}
}

func (cms *CountMinSketch) Estimate(item []byte) int {
	minCount := int(^uint(0) >> 1)
	for i := 0; i < cms.depth; i++ {
		hashValue := cms.hashFn(append(item, []byte(string(rune(i)))...))
		index := hashValue % uint32(cms.width)
		count := cms.table[i][index]
		if count < minCount {
			minCount = count
		}
	}
	return minCount
}
func (cms *CountMinSketch) clear() {
	cms.table = make([][]int, cms.depth)
	for i := range cms.table {
		cms.table[i] = make([]int, cms.width)
	}
}
func hashFnv(data []byte) uint32 {
	hash := fnv.New32a()
	hash.Write(data)
	return hash.Sum32()
}
