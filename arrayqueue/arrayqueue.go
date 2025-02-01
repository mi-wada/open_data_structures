package arrayqueue

import "github.com/mi-wada/open_data_structures/interfaces"

type ArrayQueue[T any] struct {
	s    []T
	head int
	size int
}

func New[T any]() ArrayQueue[T] {
	return ArrayQueue[T]{s: make([]T, 0), head: 0, size: 0}
}

func (aq ArrayQueue[T]) Size() int {
	return aq.size
}

var _ interfaces.Queue[struct{}] = &ArrayQueue[struct{}]{}

func (aq *ArrayQueue[T]) Add(x T) {
	aq.resize()
	aq.s[(aq.head+aq.Size())%len(aq.s)] = x
	aq.size++
}

func (aq *ArrayQueue[T]) Remove() T {
	removed := aq.s[aq.head]
	aq.head = (aq.head + 1) % len(aq.s)
	return removed
}

func (aq *ArrayQueue[T]) resize() {
	if aq.Size() >= len(aq.s) || len(aq.s) > aq.Size()*3 {
		aq.s = append(aq.s[aq.head:], aq.s[0:aq.head]...)
		aq.s = append(aq.s, make([]T, aq.Size()+1)...)
		aq.head = 0
	}
}
