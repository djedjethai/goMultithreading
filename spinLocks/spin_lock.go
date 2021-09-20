package main

import (
	"runtime"
	"sync"
	"sync/atomic"
)

type SpinLock int32

func (s *SpinLock) Lock() {
	for !atomic.CompareAndSwapInt32((*int32)(s), 0, 1) {
		runtime.Gosched()
	}
}

func (s *SpinLock) Unlock() {
	atomic.StoreInt32((*int32)(s), 0)
}

// func NewSpinLock() *SpinLock {
// the sync.Locker interface is matched by SpinLock
// as the interface is Lock() and Unlock().
// which our SpinLock implement it
func NewSpinLock() sync.Locker {
	var lock SpinLock
	return &lock
}
