package arrayqueue

type ArrayQueue[T any] struct {
	s    []T
	head int
	len  int
}

func New[T any]() ArrayQueue[T] {
	return ArrayQueue[T]{s: make([]T, 1), head: 0, len: 0}
}

func (aq ArrayQueue[T]) Len() int {
	return aq.len
}

func (aq ArrayQueue[T]) tail() int {
	return (aq.head + aq.len) % len(aq.s)
}

func (aq *ArrayQueue[T]) Push(x T) {
	aq.Resize()

	aq.s[aq.tail()] = x
	aq.len++
}

func (aq *ArrayQueue[T]) Pop() T {
	x := aq.s[aq.head]
	aq.head = (aq.head + 1) % len(aq.s)
	aq.len--
	return x
}

func (aq *ArrayQueue[T]) Resize() {
	if aq.Len() < len(aq.s) {
		return
	}
	aq.s = append(aq.s[aq.head:], aq.s[:aq.tail()]...)
	aq.s = append(aq.s, make([]T, len(aq.s))...)

	aq.head = 0
}
