package main

import "sync/atomic"

type counter struct {
	val	uint64
}

// inc increments the value of counter of 1. To make sure
// multiple threads are able to operate on the counter at the
// same time, atomic increments are used.
func (c *counter) inc() {
	atomic.AddUint64(&c.val, 1)
}

// value returns the current value of the counter, it does
// so in a safe manner.
func (c *counter) value() uint64 {
	return atomic.LoadUint64(&c.val)
}

func newCounter() *counter {
	return &counter{}
}