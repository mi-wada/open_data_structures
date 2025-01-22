package uset

type USet[T comparable] map[T]struct{}

func NewUSet[T comparable]() USet[T] {
	return make(USet[T])
}

func (us USet[T]) Add(v T) {
	us[v] = struct{}{}
}

func (us USet[T]) Len() int {
	return len(us)
}

func (us USet[T]) Remove(v T) (T, bool) {
	var deletedV T
	var deleted bool
	if _, ok := us.Find(v); ok {
		deletedV = v
		deleted = true
	}

	delete(us, v)
	return deletedV, deleted
}

func (us USet[T]) Find(v T) (T, bool) {
	if _, ok := us[v]; ok {
		return v, true
	} else {
		var zero T
		return zero, false
	}
}
