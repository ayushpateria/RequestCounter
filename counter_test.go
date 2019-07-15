package main

import "testing"

func TestCounter(t *testing.T) {
	c := newCounter()
	if expected, got := uint64(0), c.value(); expected != got {
		t.Errorf("Expected %d, got %d.", expected, got)
	}
	c.inc()
	if expected, got := uint64(1), c.value(); expected != got {
		t.Errorf("Expected %d, got %d.", expected, got)
	}
}
