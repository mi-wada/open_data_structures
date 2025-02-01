package arraystack

import "github.com/mi-wada/open_data_structures/interfaces"

type ArrayStack[T any] struct {
	s []T
}

func New[T any]() ArrayStack[T] {
	return ArrayStack[T]{s: make([]T, 0)}
}

var _ interfaces.List[struct{}] = &ArrayStack[struct{}]{}

func (as ArrayStack[T]) Size() int { return len(as.s) }

func (as ArrayStack[T]) Get(i int) T {
	return as.s[i]
}

func (as *ArrayStack[T]) Set(i int, x T) T {
	y := as.s[i]
	as.s[i] = x
	return y
}

func (as *ArrayStack[T]) Add(i int, x T) {
	as.s = append(append(as.s[:i], x), as.s[i:]...)
}

func (as *ArrayStack[T]) Remove(i int) T {
	removed := as.s[i]
	as.s = append(as.s[:i], as.s[i+1:]...)
	return removed
}

var _ interfaces.Stack[int] = &ArrayStack[int]{}

func (as *ArrayStack[T]) Push(x T) {
	as.Add(as.Size(), x)
}

func (as *ArrayStack[T]) Pop() T {
	return as.Remove(as.Size() - 1)
}
