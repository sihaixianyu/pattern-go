package iter

type Collection[T any] interface {
	NewIterator() Iterator[T]
}

type Iterator[T any] interface {
	HasNext() bool
	GetNext() *T
}
