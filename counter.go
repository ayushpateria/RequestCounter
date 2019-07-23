package main

import (
	"sync/atomic"
	"time"
)

type counter struct {
	val	uint64
	// q stores the information to calculate qps
	q *qps
}

// inc increments the value of counter by 1. 
func (c *counter) inc() {
	atomic.AddUint64(&c.val, 1)
}

// value returns the current value of the counter, it does
// so in a safe manner.
func (c *counter) value() uint64 {
	return atomic.LoadUint64(&c.val)
}

// update the counts for qps calculation
func (c *counter) updateQps() {
	currCount := c.value()
	c.q.updateCounts(currCount)
}

// qps returns the number of requests served in the last second.
func (c *counter) qps() uint64 {
	return c.q.getQps()
}

func newCounter() *counter {
	c := &counter{q: &qps{}}
	// Start a go routine which will capture the current
	// value of counter for qps calculation.
	go func() {
		for {
			time.Sleep(1 * time.Second)
			c.updateQps()
		}
	}()
	return c
}

// qps tracks the value of current and last second counter values.
type qps struct {
	currCount uint64
	lastSecCount uint64
}

// updateCounts updates the curr and last count, this
// is called every second from counter.
func (q *qps) updateCounts(currCount uint64) {
	q.lastSecCount = q.currCount
	q.currCount = currCount
}

func (q *qps) getQps() uint64 {
	return q.currCount - q.lastSecCount 
}