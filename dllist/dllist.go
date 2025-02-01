package dllist

type DLList[T any] struct {
	root *Node[T]
	len  int
}

func New[T any]() DLList[T] {
	root := newNode(*new(T))
	return DLList[T]{
		root: &root,
		len:  0,
	}
}

type Node[T any] struct {
	x    T
	prev *Node[T]
	next *Node[T]
}

func (n *Node[T]) getNext(i int) *Node[T] {
	if i == 0 {
		return n
	}
	return n.next.getNext(i - 1)
}

func (n *Node[T]) getBack(i int) *Node[T] {
	if i == 0 {
		return n
	}
	return n.prev.getBack(i - 1)
}

func newNode[T any](x T) Node[T] {
	return Node[T]{
		x:    x,
		prev: nil,
		next: nil,
	}
}

func (dl *DLList[T]) getNode(i int) *Node[T] {
	if i == 0 {
		return dl.root.next
	}

	if i < dl.len/2 {
		return dl.root.next.getNext(i - 1)
	} else {
		return dl.root.prev.getBack(i - 1)
	}
}

func (dl DLList[T]) get(i int) T {
	return dl.getNode(i).x
}

func (dl *DLList[T]) set(i int, x T) T {
	removed := dl.getNode(i).x
	dl.getNode(i).x = x
	return removed
}

func (dl *DLList[T]) AddBefore(n *Node[T], x T) {
	u := newNode(x)
	u.prev = n.prev
	u.prev.next = &u
	u.next = n
	n.prev = &u
}
