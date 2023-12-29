package bloomfilter

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

type CustomTypeMulti struct {
	Value      string
	floatValue float32
}

func (c CustomTypeMulti) MarshalBinary() ([]byte, error) {
	return json.Marshal(c)
}

func TestGivenBloomFilterWhenObjectAddedThenObjectIsReported(t *testing.T) {
	// Create a new Bloom filter with 10 bits and 3 hash functions
	bloomFilter := NewBloomFilter[CustomTypeMulti](10, 3)
	apple := CustomTypeMulti{
		Value:      "apple",
		floatValue: rand.Float32(),
	}
	banana := CustomTypeMulti{
		Value:      "banana",
		floatValue: rand.Float32(),
	}

	bloomFilter.Add(apple)
	assert.True(t, bloomFilter.Contains(apple), "Expected 'apple' to be in the Bloom filter")
	assert.False(t, bloomFilter.Contains(banana), "Expected 'orange' to not be in the Bloom filter")
}

func TestGivenBloomFilterWhen(t *testing.T) {
	// Create a new Bloom filter with 10 bits and 3 hash functions
	bloomFilter := NewBloomFilter[CustomTypeMulti](10, 3)
	apple := CustomTypeMulti{
		Value:      "apple",
		floatValue: rand.Float32(),
	}
	banana := CustomTypeMulti{
		Value:      "banana",
		floatValue: rand.Float32(),
	}

	bloomFilter.Add(apple)
	assert.True(t, bloomFilter.Contains(apple), "Expected 'apple' to be in the Bloom filter")
	assert.False(t, bloomFilter.Contains(banana), "Expected 'orange' to not be in the Bloom filter")
}

func TestBloomFilter_AddAndContains_GivenNonEmptyFilter_WhenAddingAndCheckingForContainment_ThenExpectCorrectResults(t *testing.T) {
	//given
	bloomFilter := NewBloomFilter[CustomTypeMulti](10, 3)
	//when
	apple := CustomTypeMulti{
		Value:      "apple",
		floatValue: rand.Float32(),
	}
	banana := CustomTypeMulti{
		Value:      "banana",
		floatValue: rand.Float32(),
	}

	bloomFilter.Add(apple)
	//then
	assert.True(t, bloomFilter.Contains(apple), "Expected 'apple' to be in the Bloom filter")
	assert.False(t, bloomFilter.Contains(banana), "Expected 'orange' to not be in the Bloom filter")
}

func TestBloomFilter_Empty_GivenEmptyFilter_WhenCheckingForContainment_ThenExpectNoMatches(t *testing.T) {
	// Given
	bloomFilter := NewBloomFilter[CustomTypeMulti](10, 3)
	apple := CustomTypeMulti{
		Value:      "apple",
		floatValue: rand.Float32(),
	}

	// When/Then
	assert.False(t, bloomFilter.Contains(apple), "Expected 'apple' to not be in an empty Bloom filter")
}
