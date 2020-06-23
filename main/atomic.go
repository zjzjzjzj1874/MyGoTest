package main

import "sync/atomic"

// package atomic use
type AtomicUser struct {
	name atomic.Value
	age  atomic.Value
}

func NewAtomicUser() *AtomicUser {
	a := &AtomicUser{}
	a.name.Store("")
	a.age.Store(uint64(0))
	return a
}
