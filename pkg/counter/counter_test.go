package counter

import (
	"testing"
	"time"
)

func TestCounter(t *testing.T) {
	c := NewCounter()
	if expected, got := uint64(0), c.Value(); expected != got {
		t.Errorf("Expected %d, got %d.", expected, got)
	}
	c.Inc()
	if expected, got := uint64(1), c.Value(); expected != got {
		t.Errorf("Expected %d, got %d.", expected, got)
	}
}

func TestQps(t *testing.T) {
	c := NewCounter()
	if expected, got := uint64(0), c.Qps(); expected != got {
		t.Errorf("Expected %d, got %d.", expected, got)
	}
	c.Inc()
	c.Inc()
	time.Sleep(1 * time.Second)
	if expected, got := uint64(2), c.Qps(); expected != got {
		t.Errorf("Expected %d, got %d.", expected, got)
	}
}
