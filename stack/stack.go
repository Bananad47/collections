package stack

type Stack[T any] interface {
	Top() T
	Pop()
	Push(T)
	IsEmpty() bool
	Size() int
}

type stack[T any] struct {
	arr []T
}

func (s *stack[T]) Size() int {
	return len(s.arr)
}

func (s *stack[T]) Top() T {
	var val T
	if s.IsEmpty() {
		return val
	}
	return s.arr[len(s.arr)-1]
}

func (s *stack[T]) Pop() {
	if !s.IsEmpty() {
		s.arr = s.arr[:len(s.arr)-1]
	}
}

func (s *stack[T]) Push(el T) {
	s.arr = append(s.arr, el)
}

func (s *stack[T]) IsEmpty() bool {
	return len(s.arr) == 0
}

func NewStack[T any]() Stack[T] {
	return &stack[T]{
		arr: []T{},
	}
}
