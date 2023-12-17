package bloomfiter

import (
	"encoding"
)

type BloomFilter[T encoding.BinaryMarshaler] interface {
	Add(value T)
	Contains(value T) bool
}
