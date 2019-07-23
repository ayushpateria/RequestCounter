package main

import (
	"testing"
	"time"
)
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

func TestQps(t *testing.T) {
	c := newCounter()
	if expected, got := uint64(0), c.qps(); expected != got {
		t.Errorf("Expected %d, got %d.", expected, got)
	}
	c.inc()
	c.inc()
	time.Sleep(1 * time.Second)
	if expected, got := uint64(2), c.qps(); expected != got {
		t.Errorf("Expected %d, got %d.", expected, got)
	}
}
