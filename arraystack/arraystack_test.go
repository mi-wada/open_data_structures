package arraystack_test

import (
	"testing"

	"github.com/mi-wada/open_data_structures/arraystack"
)

func TestArrayStack(t *testing.T) {
	as := arraystack.New[int]()
	if got, want := as.Size(), 0; got != want {
		t.Errorf("ArrayStack.Size() = %v, want %v", got, want)
	}

	as.Add(0, 0)
	if got, want := as.Size(), 1; got != want {
		t.Errorf("ArrayStack.Size() = %v, want %v", got, want)
	}
	as.Add(1, 1)
	if got, want := as.Remove(1), 1; got != want {
		t.Errorf("ArrayStack.Remove() = %v, want %v", got, want)
	}
	if got, want := as.Remove(0), 0; got != want {
		t.Errorf("ArrayStack.Remove() = %v, want %v", got, want)
	}

	as.Push(0)
	as.Push(1)
	if got, want := as.Pop(), 1; got != want {
		t.Errorf("ArrayStack.Pop() = %v, want %v", got, want)
	}
	if got, want := as.Pop(), 0; got != want {
		t.Errorf("ArrayStack.Pop() = %v, want %v", got, want)
	}
}
