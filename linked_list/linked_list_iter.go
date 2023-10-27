package linked_list

type linkedListIterator[T any] struct {
	nd *node[T]
}

func (i *linkedListIterator[T]) IsEnd() bool {
	if i.nd == nil {
		return true
	}
	return false
}

func (i *linkedListIterator[T]) Next() bool {
	if i.IsEnd() {
		return false
	}
	i.nd = i.nd.next
	return true
}

func (i *linkedListIterator[T]) Value() T {
	var res T
	if !i.IsEnd() {
		res = i.nd.val
	}
	return res
}

func (i *linkedListIterator[T]) ChangeValue(val T) {
	if i.IsEnd() {
		return
	}
	i.nd.val = val
}
