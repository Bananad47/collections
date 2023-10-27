package linked_list

import "github.com/collections/iter"

type LinkedList[T any] interface {
	PushBack(val ...T)
	Insert(val T, pos int)
	Delete(pos int) bool
	Size() int
	IsEmpty() bool
	Iterator() iter.Iterator[T]
}

type linkedList[T any] struct {
	nd *node[T]
}

type node[T any] struct {
	val  T
	next *node[T]
}

func (l *linkedList[T]) Size() int {
	nd := l.nd
	res := 0
	for nd != nil {
		nd = nd.next
		res++
	}
	return res
}

func (l *linkedList[T]) IsEmpty() bool {
	return l.nd == nil
}

func (l *linkedList[T]) PushBack(vals ...T) {
	if l.IsEmpty() {
		l.nd = &node[T]{
			val: vals[0],
		}
		vals = vals[1:]
	}
	nd := l.nd
	for nd.next != nil {
		nd = nd.next
	}
	for _, el := range vals {
		nd.next = &node[T]{
			val: el,
		}
		nd = nd.next
	}
}

func (l *linkedList[T]) _pushFront(val T) {
	t := *l.nd
	nd := &node[T]{
		val:  val,
		next: &t,
	}
	l.nd = nd
}

//Insert вставляет значение после позиции, если пос == 0, то вставляет значение вперед, если пос > размера, то в конец
func (l *linkedList[T]) Insert(val T, pos int) {
	if l.IsEmpty() {
		l.nd = &node[T]{
			val: val,
		}
		return
	}
	if pos == -1 {
		l._pushFront(val)
		return
	}
	nd := l.nd
	for i := 0; i != pos && nd.next != nil; i++ {
		nd = nd.next
	}
	nd.next = &node[T]{
		val:  val,
		next: nd.next,
	}
}

func (l *linkedList[T]) Delete(pos int) bool {
	if l.IsEmpty() {
		return false
	}
	if pos == 0 {
		l.nd = l.nd.next
		return true
	}
	nd := l.nd
	for i := 1; i != pos && nd.next != nil; i++ {
		nd = nd.next
	}
	if nd.next == nil {
		return false
	}
	nd.next = nd.next.next
	return true
}

func (l *linkedList[T]) Iterator() iter.Iterator {
	return &linkedListIterator[T]{
		nd: l.nd,
	}
}

func NewLinkedList[T any](constructors ...func(list LinkedList[T])) LinkedList[T] {
	ll := &linkedList[T]{}
	for _, el := range constructors {
		el(ll)
	}
	return ll
}

func AddDefaultValueConstructor[T any](vals ...T) func(LinkedList[T]) {
	return func(ll LinkedList[T]) {
		ll.PushBack(vals...)
	}
}
