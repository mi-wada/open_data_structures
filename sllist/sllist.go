// A package for Singly-linked list.
package sllist

type SLList[T any] struct {
	head *Node[T]
	tail *Node[T]
	len  int
}

type Node[T any] struct {
	x    T
	next *Node[T]
}

func newNode[T any](x T) Node[T] {
	return Node[T]{
		x:    x,
		next: nil,
	}
}

func New[T any]() SLList[T] {
	return SLList[T]{
		head: nil,
		tail: nil,
		len:  0,
	}
}

func (sl SLList[T]) Len() int {
	return sl.len
}

func (sl *SLList[T]) Push(x T) {
	n := newNode(x)
	if sl.head != nil {
		n.next = sl.head
	}
	if sl.tail == nil {
		sl.tail = &n
	}
	sl.head = &n
	sl.len++
}

func (sl *SLList[T]) Pop() T {
	removed := sl.head.x
	sl.head = sl.head.next
	return removed
}

func (sl *SLList[T]) Add(x T) {
	n := newNode(x)
	if sl.tail != nil {
		sl.tail.next = &n
	}
	if sl.head == nil {
		sl.head = &n
	}
	sl.tail = &n
	sl.len++
}

func (sl *SLList[T]) Remove() T {
	return sl.Pop()
}
