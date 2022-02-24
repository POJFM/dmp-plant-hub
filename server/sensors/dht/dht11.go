// taken and modified
// https://github.com/bosima/go-pidriver/blob/7149880fa03edc7206b58d783ce8ad9882391e00/drivers/dht11.go

package dht

import (
	"errors"
	"strconv"
	"time"

	"github.com/stianeikeland/go-rpio/v4"
)

type DHT11 struct {
	pin          rpio.Pin
	val          []uint8
	closeFunc    func() error
	lastReadTime time.Time
}

func NewDHT11(pin int) *DHT11 {
	return &DHT11{pin: rpio.Pin(pin)}
}

func (d *DHT11) ReadData() (tmp float64, hum float64, err error) {
	if !d.lastReadTime.IsZero() && time.Since(d.lastReadTime).Milliseconds() <= 1000 {
		return -1, -1, errors.New("read interval must be greater than 1 seconds")
	}

	d.val = []uint8{0, 0, 0, 0, 0}

	resetDht(d.pin)

	status := checkDhtStatus(d.pin)
	if !status {
		return -1, -1, errors.New("device is not ready")
	}

	// dht output data: 40bit
	for i := 0; i < 5; i++ {
		for k := 0; k < 8; k++ {
			v := readBit(d.pin)
			if v == rpio.High {
				leftLength := 8 - k - 1
				if leftLength > 0 {
					d.val[i] = d.val[i] | 1<<leftLength
				} else {
					d.val[i] = d.val[i] | 1
				}
			}
		}
	}

	d.lastReadTime = time.Now()

	if !checkData(d.val) {
		return -1, -1, errors.New("data verification failed")
	}

	hum, err = strconv.ParseFloat(strconv.Itoa(int(d.val[0]))+"."+strconv.Itoa(int(d.val[1])), 32)
	tmp, err = strconv.ParseFloat(strconv.Itoa(int(d.val[2]))+"."+strconv.Itoa(int(d.val[3])), 32)

	return
}

func (d *DHT11) Close() error {
	return d.closeFunc()
}

func resetDht(p rpio.Pin) {
	p.Output()
	p.High()
	Delay(2)

	// send start signal, must great than 18ms
	p.Low()
	Delay(25)

	// then over the signal
	p.High()

	// ready to read data
	p.Input()
	p.PullUp()

	// wait 20-40us
	DelayMicroseconds(30)
}

func checkDhtStatus(p rpio.Pin) bool {
	// dht response start: first 80us low, then 80us high
	wait := 0
	for wait < 100 {
		if v := p.Read(); v == rpio.Low {
			break
		}
		DelayMicroseconds(1)
		wait += 1
	}
	if wait >= 100 {
		return false
	}

	wait = 0
	for wait < 100 {
		if v := p.Read(); v == rpio.High {
			break
		}
		DelayMicroseconds(1)
		wait += 1
	}
	return wait < 100
}

func readBit(p rpio.Pin) rpio.State {
	// for per bit: first 50us low, then 26-70us high
	// 26-28us high represents 0
	// 70us high represents 1
	wait := 0
	for wait < 100 {
		if v := p.Read(); v == rpio.Low {
			break
		}
		DelayMicroseconds(1)
		wait += 1
	}

	wait = 0
	for wait < 100 {
		if v := p.Read(); v == rpio.High {
			break
		}
		DelayMicroseconds(1)
		wait += 1
	}

	DelayMicroseconds(40)

	return p.Read()
}

func checkData(data []uint8) bool {
	sum := 0
	for _, v := range data[:4] {
		sum += int(v)
	}

	return sum == int(data[4])
}
