package main

import (
	"sync/atomic"
	"time"
	"sync"
)

type counter struct {
	val	uint64
	q *qps
}

// inc increments the value of counter by 1.
func (c *counter) inc() {
	atomic.AddUint64(&c.val, 1)
	// Update qps
	c.q.update()
}

// value returns the current value of the counter, it does
// so in a safe manner.
func (c *counter) value() uint64 {
	return atomic.LoadUint64(&c.val)
}

// qps returns the number of requests served in the last second.
func (c *counter) qps() uint64 {
	return c.q.getQps()
}

func newCounter() *counter {
	return &counter{q: &qps{}}
}


type qps struct {
	currCount uint64
	countLastSec uint64
	ts int64
	lock sync.RWMutex
}

func (q *qps) update() {
	q.lock.Lock()
	defer q.lock.Unlock()
	ts := time.Now().Unix()
	if q.ts == ts {
		q.currCount += 1
	} else {
		q.ts = ts
		currCount := q.currCount
		q.countLastSec = currCount
		q.currCount = 0
	}
}

func (q *qps) getQps() uint64 {
	q.lock.RLock()
	defer q.lock.RUnlock()
	return q.countLastSec
}