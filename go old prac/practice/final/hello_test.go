package main

import (
	"testing"
)

func TestNew(t *testing.T) {
	a := new()
	if len(a) != 6 {
		t.Errorf("length is %d", len(a))
	}
}
