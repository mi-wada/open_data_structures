package uset_test

import (
	"testing"

	"github.com/mi-wada/open_data_structures/uset"
)

func TestUSet(t *testing.T) {
	us := uset.NewUSet[string]()
	if got, want := us.Len(), 0; got != want {
		t.Errorf("got = %v, want = %v", got, want)
	}

	us.Add("hello")
	if got, want := us.Len(), 1; got != want {
		t.Errorf("got = %v, want = %v", got, want)
	}
	v, ok := us.Find("hello")
	wantV, wantOk := "hello", true
	if v != wantV {
		t.Errorf("v = %v, wantV = %v", v, wantV)
	}
	if ok != wantOk {
		t.Errorf("ok = %v, wantOk = %v", ok, wantOk)
	}

	us.Remove("hello")
	if got, want := us.Len(), 0; got != want {
		t.Errorf("got = %v, want = %v", got, want)
	}
}
