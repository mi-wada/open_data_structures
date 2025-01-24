package arraystack

type ArrayStack[T any] []T

func New[T any]() ArrayStack[T] {
	return make(ArrayStack[T], 0)
}

func (as ArrayStack[T]) Len() int {
	return len(as)
}

func (as ArrayStack[T]) Get(i int) T {
	return as[i]
}

func (as *ArrayStack[T]) Set(i int, x T) T {
	removed := (*as)[i]
	(*as)[i] = x
	return removed
}

func (as *ArrayStack[T]) Add(i int, x T) {
	*as = append((*as)[:i], append([]T{x}, (*as)[i:]...)...)
}

func (as *ArrayStack[T]) Remove(i int) T {
	x := (*as)[i]
	*as = append((*as)[:i], (*as)[i+1:]...)
	return x
}

func (as *ArrayStack[T]) Push(x T) {
	as.Add(as.Len(), x)
}

func (as *ArrayStack[T]) Pop() T {
	return as.Remove(as.Len() - 1)
}
