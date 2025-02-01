package arrayqueue_test

import (
	"testing"

	"github.com/mi-wada/open_data_structures/arrayqueue"
)

func TestArrayQueue(t *testing.T) {
	aq := arrayqueue.New[int]()
	if got, want := aq.Size(), 0; got != want {
		t.Errorf("aq.Size() = %v, want %v", got, want)
	}

	aq.Add(0)
	aq.Add(1)
	if got, want := aq.Size(), 2; got != want {
		t.Errorf("aq.Size() = %v, want %v", got, want)
	}
	if got, want := aq.Remove(), 0; got != want {
		t.Errorf("aq.Remove() = %v, want %v", got, want)
	}
	if got, want := aq.Remove(), 1; got != want {
		t.Errorf("aq.Remove() = %v, want %v", got, want)
	}
}
