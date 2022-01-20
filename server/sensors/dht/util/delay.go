package util

import (
	"syscall"
	"time"
)

// this make the CPU busy, but there is no other better way
func DelayMicroseconds(us int64) {
	var tv syscall.Timeval
	_ = syscall.Gettimeofday(&tv)

	stratTick := int64(tv.Sec)*int64(1000000) + int64(tv.Usec) + us
	endTick := int64(0)
	for endTick < stratTick {
		_ = syscall.Gettimeofday(&tv)
		endTick = int64(tv.Sec)*int64(1000000) + int64(tv.Usec)
	}
}

func Delay(ms int64) {
	time.Sleep(time.Millisecond * time.Duration(ms))
}
