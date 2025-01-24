package sllist_test

import (
	"testing"

	"github.com/mi-wada/open_data_structures/sllist"
)

func TestSLList(t *testing.T) {
	sl := sllist.New[int]()

	sl.Push(1)
	sl.Push(2)
	if got, want := sl.Len(), 2; got != want {
		t.Errorf("Push failed, got %v, want %v", got, want)
	}

	if got, want := sl.Pop(), 2; got != want {
		t.Errorf("Pop failed, got %v, want %v", got, want)
	}
	if got, want := sl.Pop(), 1; got != want {
		t.Errorf("Pop failed, got %v, want %v", got, want)
	}

	sl.Add(1)
	sl.Add(2)
	if got, want := sl.Pop(), 1; got != want {
		t.Errorf("Pop failed, got %v, want %v", got, want)
	}
	if got, want := sl.Pop(), 2; got != want {
		t.Errorf("Pop failed, got %v, want %v", got, want)
	}
}
