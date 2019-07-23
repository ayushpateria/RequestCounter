package main

import (
	"sync/atomic"
	"time"
)

type qps struct {
	currCount uint64
	countLastSec uint64
	ts int64
}

type counter struct {
	val	uint64
	// byTs is a mapping from timestamp to count, it
	// stores the count of the counter at the end of that
	// second, and is used to calculate qps. We only store
	// latest three timestamps in byTs, as only last two
	// seconds are required to calculate qps.
	q qps
}

// inc increments the value of counter by 1. It also
// stores the current count to `byTs` for qps calculation
// To make sure multiple threads are able to operate on the
// counter at the same time, atomic increments are used.
func (c *counter) inc() {
	atomic.AddUint64(&c.val, 1)
	// Storing latest three seconds in byTs for qps calculation.
	ts := time.Now().Unix()
	if atomic.LoadInt64(&c.q.ts) == ts {
		atomic.AddUint64(&c.q.currCount, 1)
	} else {
		atomic.StoreInt64(&c.q.ts, ts)
		currCount := atomic.LoadUint64(&c.q.currCount)
		atomic.StoreUint64(&c.q.countLastSec, currCount)
		atomic.StoreUint64(&c.q.currCount, 0)
	}
}

// value returns the current value of the counter, it does
// so in a safe manner.
func (c *counter) value() uint64 {
	return atomic.LoadUint64(&c.val)
}

// qps returns the number of requests served in the last second.
func (c *counter) qps() uint64 {
	return c.q.countLastSec
}

func newCounter() *counter {
	return &counter{q: qps{}}
}
