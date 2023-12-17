package bloomfiter

type BloomFilter[T any] interface {
	add(value T)
	exists(value T)
}
