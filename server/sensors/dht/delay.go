// taken and modified
// https://github.com/bosima/go-pidriver/blob/7149880fa03edc7206b58d783ce8ad9882391e00/util/delay.go

package dht

import (
	"syscall"
	"time"
)

// DelayMicroseconds
// this make the CPU busy, but there is no other better way
// not sure if this is good 🤔
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
