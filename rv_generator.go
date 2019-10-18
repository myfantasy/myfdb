package main

import (
	"sync"
	"syscall"
	"time"
)

var rvc int64
var rvgl int64
var rvmx sync.Mutex

// RvGet2 - Generate Next RV (sync)
func RvGet2() int64 {
	t := time.Now().UnixNano() / 10000 * 10000

	rvmx.Lock()

	if rvgl == t {
		rvc = rvc + 1
		t = t + rvc
	} else {
		rvc = 0
		rvgl = t
	}

	rvmx.Unlock()

	return t
}

// RvGet - Generate Next RV (sync)
func RvGet() int64 {
	t := GoTime() / 10000 * 10000

	rvmx.Lock()

	if rvgl == t {
		rvc = rvc + 1
		t = t + rvc
	} else {
		rvc = 0
		rvgl = t
	}

	rvmx.Unlock()

	return t
}

// GoTime - fast get time
func GoTime() int64 {
	a := syscall.Timeval{}
	syscall.Gettimeofday(&a)
	return syscall.TimevalToNsec(a)
}
