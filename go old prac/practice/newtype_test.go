package main

import "testing"

func TestNewtype(t *testing.T) {
	m := newtype()
	if len(m) != 16 {
		t.Errorf("expected length of 16,but got %d", len(m))
	}
}
