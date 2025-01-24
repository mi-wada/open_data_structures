package dualarraydeque

import "github.com/mi-wada/open_data_structures/arraystack"

type DualArrayDeque[T any] struct {
	front arraystack.ArrayStack[T]
	back  arraystack.ArrayStack[T]
}

func New[T any]() DualArrayDeque[T] {
	return DualArrayDeque[T]{
		front: arraystack.New[T](),
		back:  arraystack.New[T](),
	}
}

func (d DualArrayDeque[T]) Len() int {
	return d.front.Len() + d.back.Len()
}

func (d DualArrayDeque[T]) Get(i int) T {
	if i < d.front.Len() {
		return d.front.Get(d.front.Len() - i - 1)
	} else {
		return d.back.Get(i - d.front.Len())
	}
}

func (d *DualArrayDeque[T]) Set(i int, v T) T {
	if i < d.front.Len() {
		return d.front.Set(d.front.Len()-i-1, v)
	} else {
		return d.back.Set(i-d.front.Len(), v)
	}
}

func (d *DualArrayDeque[T]) Add(i int, v T) {
	if i < d.front.Len() {
		d.front.Add(d.front.Len()-i, v)
	} else {
		d.back.Add(i-d.front.Len(), v)
	}

	d.balance()
}

func (d *DualArrayDeque[T]) Remove(i int) T {
	var removed T
	if i < d.front.Len() {
		removed = d.front.Remove(d.front.Len() - i - 1)
	} else {
		removed = d.back.Remove(i - d.front.Len())
	}

	d.balance()
	return removed
}

func (d *DualArrayDeque[T]) balance() {
	if d.back.Len() > 3*d.front.Len() ||
		d.front.Len() > 3*d.back.Len() {
		n := d.Len()

		nf := n / 2
		sf := make([]T, nf)
		for i := 0; i < nf; i++ {
			sf[nf-i-1] = d.Get(i)
		}

		nb := n - nf
		sb := make([]T, nb)
		for i := 0; i < nb; i++ {
			sb[i] = d.Get(i + nf)
		}
		d.front = arraystack.ArrayStack[T](sf)
		d.back = arraystack.ArrayStack[T](sb)
	}
}
