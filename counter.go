package main

import (
	"sync/atomic"
	"time"
)

type counter struct {
	val	uint64
	// byTs is a mapping from timestamp to count, it
	// stores the count of the counter at the end of that
	// second, and is used to calculate qps. We only store
	// latest three timestamps in byTs, as only last two
	// seconds are required to calculate qps.
	byTs map[int64]uint64
}

// inc increments the value of counter by 1. It also
// stores the current count to `byTs` for qps calculation
// To make sure multiple threads are able to operate on the
// counter at the same time, atomic increments are used.
func (c *counter) inc() {
	currVal := atomic.AddUint64(&c.val, 1)
	// Storing latest three seconds in byTs for qps calculation.
	ts := time.Now().Unix()
	_, ok := c.byTs[ts-1]
	// If there was no increment in the previous second,
	// assign previous two seconds from the current value.
	if !ok {
		prevVal := currVal - 1
		c.byTs[ts-1] = prevVal
		c.byTs[ts-2] = prevVal
	}
	// Reassigning map with the latest three seconds.
	c.byTs = map[int64]uint64 {
		ts: currVal,
		ts-1: c.byTs[ts-1],
		ts-2: c.byTs[ts-2],
	}
}

// value returns the current value of the counter, it does
// so in a safe manner.
func (c *counter) value() uint64 {
	return atomic.LoadUint64(&c.val)
}

// qps returns the number of requests served in the last second.
func (c *counter) qps() uint64 {
	ts := time.Now().Unix()
	if c.byTs[ts-1] == 0 {
		return 0
	}
	return c.byTs[ts-1] - c.byTs[ts-2]
}

func newCounter() *counter {
	return &counter{byTs: make(map[int64]uint64)}
}
