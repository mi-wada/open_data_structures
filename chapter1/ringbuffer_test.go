package chapter1_test

import (
	"testing"

	"github.com/mi-wada/open_data_structures/chapter1"
)

func TestRingBuffer_Push(t *testing.T) {
	rb := chapter1.NewRingBuffer(2)

	if got, want := rb.First(), ""; got != want {
		t.Errorf("got %s, expected %s", got, want)
	}
	if got, want := rb.Last(), ""; got != want {
		t.Errorf("got %s, expected %s", got, want)
	}

	rb.Push("0")
	if got, want := rb.First(), ""; got != want {
		t.Errorf("got %s, expected %s", got, want)
	}
	if got, want := rb.Last(), "0"; got != want {
		t.Errorf("got %s, expected %s", got, want)
	}

	rb.Push("1")
	if got, want := rb.First(), "0"; got != want {
		t.Errorf("got %s, expected %s", got, want)
	}
	if got, want := rb.Last(), "1"; got != want {
		t.Errorf("got %s, expected %s", got, want)
	}

	rb.Push("2")
	if got, want := rb.First(), "1"; got != want {
		t.Errorf("got %s, expected %s", got, want)
	}
	if got, want := rb.Last(), "2"; got != want {
		t.Errorf("got %s, expected %s", got, want)
	}
}
