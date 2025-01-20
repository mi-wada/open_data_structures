package chapter1

type RingBuffer struct {
	s        []string
	firstCur int
}

func NewRingBuffer(size int) *RingBuffer {
	return &RingBuffer{s: make([]string, size), firstCur: 0}
}

func (rb *RingBuffer) Push(v string) {
	rb.s[rb.firstCur%rb.Len()] = v
	rb.firstCur++
}

func (rb *RingBuffer) First() string {
	return rb.s[rb.firstIdx()]
}

func (rb *RingBuffer) Last() string {
	return rb.s[rb.lastIdx()]
}

func (rb *RingBuffer) Full() bool {
	return rb.firstCur >= len(rb.s)-1
}

func (rb *RingBuffer) firstIdx() int {
	return rb.firstCur % rb.Len()
}

func (rb *RingBuffer) lastIdx() int {
	return (rb.firstCur + rb.Len() - 1) % rb.Len()
}

func (rb *RingBuffer) Len() int {
	return len(rb.s)
}
