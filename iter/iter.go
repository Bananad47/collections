package iter

type Iterator[T any] interface {
	Next() bool
	Value() T
	ChangeValue(T)
	IsEnd() bool
}
