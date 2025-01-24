package arraystack_test

import (
	"testing"

	"github.com/mi-wada/open_data_structures/arraystack"
)

func TestArrayStack(t *testing.T) {
	as := arraystack.New[int]()
	if as.Len() != 0 {
		t.Errorf("New() failed")
	}

	as.Push(1)
	as.Push(2)
	if as.Len() != 2 {
		t.Errorf("Push() failed")
	}

	as.Add(1, 3)
	if as[0] != 1 || as[1] != 3 || as[2] != 2 {
		t.Errorf("Set() failed got: %v", as)
	}

	as.Remove(0)
	if as[0] != 3 || as[1] != 2 {
		t.Errorf("Remove() failed got: %v", as)
	}

	as.Pop()
	if as[0] != 3 {
		t.Errorf("Pop() failed got: %v", as)
	}

	as.Pop()
	if as.Len() != 0 {
		t.Errorf("Pop() failed got: %v", as)
	}
}
