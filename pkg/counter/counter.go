package counter

import (
	"sync/atomic"
	"time"
)

type Counter struct {
	val uint64
	// q stores the information to calculate qps
	q *qps
}

// Inc increments the value of Counter by 1.
func (c *Counter) Inc() {
	atomic.AddUint64(&c.val, 1)
}

// Value returns the current value of the Counter, it does
// so in a safe manner.
func (c *Counter) Value() uint64 {
	return atomic.LoadUint64(&c.val)
}

// update the counts for qps calculation
func (c *Counter) updateQps() {
	currCount := c.Value()
	c.q.updateCounts(currCount)
}

// Qps returns the number of requests served in the last second.
func (c *Counter) Qps() uint64 {
	return c.q.getQps()
}

func NewCounter() *Counter {
	c := &Counter{q: &qps{}}
	// Start a go routine which will capture the current
	// value of Counter for qps calculation.
	go func() {
		for {
			time.Sleep(1 * time.Second)
			c.updateQps()
		}
	}()
	return c
}

// qps tracks the value of current and last second Counter values.
type qps struct {
	currCount    uint64
	lastSecCount uint64
}

// updateCounts updates the curr and last count, this
// is called every second from Counter.
func (q *qps) updateCounts(currCount uint64) {
	q.lastSecCount = q.currCount
	q.currCount = currCount
}

func (q *qps) getQps() uint64 {
	return q.currCount - q.lastSecCount
}
