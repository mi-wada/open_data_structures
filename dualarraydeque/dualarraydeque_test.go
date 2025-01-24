package dualarraydeque_test

import (
	"testing"

	"github.com/mi-wada/open_data_structures/dualarraydeque"
)

func TestDualArrayDeque(t *testing.T) {
	d := dualarraydeque.New[int]()
	if got, want := d.Len(), 0; got != want {
		t.Errorf("New fails, got %v, want %v", got, want)
	}

	d.Add(0, 1)
	if got, want := d.Len(), 1; got != want {
		t.Errorf("Add() fails, got %v, want %v", got, want)
	}
	d.Add(1, 2)
	if got, want := d.Len(), 2; got != want {
		t.Errorf("Add() fails, got %v, want %v", got, want)
	}
	d.Add(2, 3)
	if got, want := d.Len(), 3; got != want {
		t.Errorf("Add() fails, got %v, want %v", got, want)
	}
	d.Add(3, 4)
	if got, want := d.Len(), 4; got != want {
		t.Errorf("Add() fails, got %v, want %v", got, want)
	}

	if got, want := d.Remove(0), 1; got != want {
		t.Errorf("Remove() fails, got %v, want %v", got, want)
	}
	if got, want := d.Remove(2), 4; got != want {
		t.Errorf("Remove() fails, got %v, want %v", got, want)
	}
}
