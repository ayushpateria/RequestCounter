package main

import (
	"sync/atomic"
	"time"
)

type counter struct {
	val	uint64
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
	go func() {
		for {
			time.Sleep(1 * time.Second)
			c.updateQps()
		}
	}()
	return c
}

type qps struct {
	currCount uint64
	lastSecCount uint64
}

func (q *qps) updateCounts(currCount uint64) {
	q.lastSecCount = q.currCount
	q.currCount = currCount
}

func (q *qps) getQps() uint64 {
	return q.currCount - q.lastSecCount 
}