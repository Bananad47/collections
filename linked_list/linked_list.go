package linked_list

type LinkedList[T any] interface {
	PushBack(val T)
	Insert(val T, pos int)
	Print()
	Delete(val T, pos int)
	Size() int
}
