package arrayqueue_test

import (
	"testing"

	"github.com/mi-wada/open_data_structures/arrayqueue"
)

func TestArrayQueue(t *testing.T) {
	aq := arrayqueue.New[int]()
	if aq.Len() != 0 {
		t.Errorf("New() failed")
	}

	aq.Push(1)
	if aq.Len() != 1 {
		t.Errorf("Push() failed, got %v want %v", aq.Len(), 1)
	}

	aq.Push(2)
	if aq.Len() != 2 {
		t.Errorf("Push() failed")
	}

	if got, want := aq.Pop(), 1; got != want {
		t.Errorf("Pop() failed, got %v, want %v", got, want)
	}
	if got, want := aq.Pop(), 2; got != want {
		t.Errorf("Pop() failed, got %v, want %v", got, want)
	}
}
