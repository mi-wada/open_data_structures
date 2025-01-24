package arraydeque

type ArrayDeque[T any] struct {
	s    []T
	head int
	len  int
}

func New[T any]() ArrayDeque[T] {
	return ArrayDeque[T]{s: make([]T, 1), head: 0, len: 0}
}

func (d ArrayDeque[T]) Len() int {
	return d.len
}

func (d ArrayDeque[T]) tail() int {
	return (d.head + d.len) % len(d.s)
}

func (d ArrayDeque[T]) Get(i int) T {
	return d.s[(d.head+i)%len(d.s)]
}

func (d *ArrayDeque[T]) Set(i int, v T) T {
	removed := d.s[(d.head+i)%len(d.s)]
	d.s[(d.head+i)%len(d.s)] = v
	return removed
}

func (d *ArrayDeque[T]) Add(i int, v T) {
	d.Resize()
	if i < d.len/2 {
		if d.head == 0 {
			d.head = len(d.s) - 1
		} else {
			d.head--
		}
		for j := 0; j < i; j++ {
			d.s[(d.head+i+j-1+len(d.s))%len(d.s)] = d.s[(d.head+i+j)%len(d.s)]
		}
	} else {
		for j := d.len - 1; j >= i; j-- {
			d.s[(d.head+j+1)%len(d.s)] = d.s[(d.head+j)%len(d.s)]
		}
	}
	d.s[(d.head+i)%len(d.s)] = v
	d.len++
}

func (d *ArrayDeque[T]) Remove(i int) T {
	removed := d.s[(d.head+i)%len(d.s)]
	if i < d.len/2 {
		if d.head == len(d.s)-1 {
			d.head = 0
		} else {
			d.head++
		}
		for j := 0; j < i; j++ {
			d.s[(d.head+i+j-1+len(d.s))%len(d.s)] = d.s[(d.head+i+j)%len(d.s)]
		}
	} else {
		for j := d.len - 1; j >= i; j-- {
			d.s[(d.head+j+1)%len(d.s)] = d.s[(d.head+j)%len(d.s)]
		}
	}
	d.len--
	return removed
}

func (d *ArrayDeque[T]) Resize() {
	if d.Len() < len(d.s) {
		return
	}
	d.s = append(d.s[d.head:], d.s[:d.tail()]...)
	d.s = append(d.s, make([]T, len(d.s))...)

	d.head = 0
}
