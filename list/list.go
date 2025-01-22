package list

import "slices"

type List[T any] []T

func (l List[T]) Len() int {
	return len(l)
}

func (l List[T]) Add(i int, v T) {
	l = slices.Insert(l, i, v) //nolint
}

func (l List[T]) Remove(i int) {
	l = slices.Delete(l, i, i) //nolint
}
