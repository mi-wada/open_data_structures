package sset

type SSet[T comparable] map[T]struct{}

func (ss SSet[T]) Add(v T) {}

func (ss SSet[T]) Len() int { return len(ss) }

func (ss SSet[T]) Find(v T) (T, bool) {
	// return most neaest value
	panic("TODO: implement later")
}

func (ss SSet[T]) Remove(v T) {
	delete(ss, v)
}
