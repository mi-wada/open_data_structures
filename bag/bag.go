package bag

import "slices"

type Bag[T comparable] map[T]int

func NewBag[T comparable]() Bag[T] {
	return make(Bag[T])
}

func (b Bag[T]) Add(v T) {
	b[v]++
}

func (b Bag[T]) Len() int {
	res := 0
	for _, v := range b {
		res += v
	}
	return res
}

func (b Bag[T]) Find(v T) (T, bool) {
	if _, ok := b[v]; ok {
		return v, true
	} else {
		return *new(T), false
	}
}

func (b Bag[T]) FindAll(v T) ([]T, bool) {
	if cnt, ok := b[v]; ok {
		return slices.Repeat([]T{v}, cnt), true
	} else {
		return nil, false
	}
}

func (b Bag[T]) Remove(v T) (T, bool) {
	var deletedV T
	var deleted bool
	if _, ok := b.Find(v); ok {
		deletedV = v
		deleted = true
	}

	b[v]--
	if b[v] <= 0 {
		delete(b, v)
	}

	return deletedV, deleted
}
