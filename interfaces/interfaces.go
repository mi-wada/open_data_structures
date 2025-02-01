package interfaces

// Queue interface represents FIFO queue.
type Queue[T any] interface {
	Add(x T)
	Remove() T
}

// Stack interface represents LIFO queue.
type Stack[T any] interface {
	Push(x T)
	Pop() T
}

// Deque interface represents Two-way queue.
type Deque[T any] interface {
	AddFirst(x T)
	RemoveFirst() T
	AddLast(x T)
	RemoveLast() T
}

// List interface represents list.
type List[T any] interface {
	// Size returns size of list.
	Size() int
	// Get returns the value at i.
	// If i is out of range index, panic.
	Get(i int) T
	// Set replace the value x at i, and returns replaced value.
	// If i is out of range index, panic.
	Set(i int, x T) T
	// Add adds the value x at i. And it move following values to back.
	// 0 <= i <= [List.Size]. If i is out of range index, panic.
	Add(i int, x T)
	// Remove removes the value at i. Ad it move following values to front.
	// If i is out of range index, panic.
	Remove(i int) T
}

// USet interface represents unordered set.
type USet[T any] interface {
	// Size returns size of set.
	Size() int
	// Add adds the value x to set.
	// If added returns true, else return false.
	Add(x T) bool
	// Remove removes the value x from set.
	// If removed returns that value, else returns nil.
	Remove(x T) *T
	// Find finds the value x.
	// If x is in set returns that value, else returns nil.
	Find(x T) *T
}

// SSet interface represents sorted set.
type SSet[T comparable] interface {
	// Size returns size of set.
	Size() int
	// Add adds the value x to set.
	// If added returns true, else return false.
	Add(x T) bool
	// Remove removes the value x from set.
	// If removed returns that value, else returns nil.
	Remove(x T) *T
	// Find finds the value >= x.
	// If found return it, else returns nil.
	Find(x T) *T
}
