package arraystack

type ArrayStack[T any] []T

func New[T any]() ArrayStack[T] {
	return make(ArrayStack[T], 0)
}

func (as ArrayStack[T]) Len() int {
	return len(as)
}

func (as *ArrayStack[T]) Insert(i int, x T) {
	*as = append((*as)[:i], append([]T{x}, (*as)[i:]...)...)
}

func (as *ArrayStack[T]) Remove(i int) T {
	x := (*as)[i]
	*as = append((*as)[:i], (*as)[i+1:]...)
	return x
}

func (as *ArrayStack[T]) Push(x T) {
	as.Insert(as.Len(), x)
}

func (as *ArrayStack[T]) Pop() T {
	return as.Remove(as.Len() - 1)
}
