package drivers

import (
	"errors"
	"github.com/SPSOAFM-IT18/dmp-plant-hub/sensors/dht/util"
	"strconv"
	"time"

	"github.com/stianeikeland/go-rpio/v4"
)

type DHT11 struct {
	PinNo        uint8
	pin          *rpio.Pin
	val          []uint8
	closeFunc    func() error
	lastReadTime time.Time
}

func NewDHT11(pinNo uint8) (*DHT11, error) {
	dht := &DHT11{PinNo: pinNo}

	err := OpenRPi()
	if err != nil {
		return nil, errors.New("init error: " + err.Error())
	}

	dht.closeFunc = func() error {
		return CloseRPi()
	}

	pin := rpio.Pin(dht.PinNo)
	dht.pin = &pin

	return dht, nil
}

func (d *DHT11) ReadData() (rh float64, tmp float64, err error) {
	if !d.lastReadTime.IsZero() && time.Since(d.lastReadTime).Milliseconds() <= 1000 {
		return -1, -1, errors.New("read interval must be greater than 1 seconds")
	}

	p := d.pin
	d.val = []uint8{0, 0, 0, 0, 0}

	resetDht(p)

	status := checkDhtStatus(p)
	if !status {
		return -1, -1, errors.New("device is not ready")
	}

	// dht output data: 40bit
	for i := 0; i < 5; i++ {
		for k := 0; k < 8; k++ {
			v := readBit(p)
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

	rh, err = strconv.ParseFloat(strconv.Itoa(int(d.val[0]))+"."+strconv.Itoa(int(d.val[1])), 32)
	tmp, err = strconv.ParseFloat(strconv.Itoa(int(d.val[2]))+"."+strconv.Itoa(int(d.val[3])), 32)

	return rh, tmp, err
}

func (d *DHT11) Close() error {
	return d.closeFunc()
}

func resetDht(p *rpio.Pin) {
	p.Output()
	p.High()
	util.Delay(2)

	// send start signal, must great than 18ms
	p.Low()
	util.Delay(25)

	// then over the signal
	p.High()

	// ready to read data
	p.Input()
	p.PullUp()

	// wait 20-40us
	util.DelayMicroseconds(30)
}

func checkDhtStatus(p *rpio.Pin) bool {
	// dht response start: first 80us low, then 80us high
	wait := 0
	for wait < 100 {
		if v := p.Read(); v == rpio.Low {
			break
		}
		util.DelayMicroseconds(1)
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
		util.DelayMicroseconds(1)
		wait += 1
	}
	return wait < 100
}

func readBit(p *rpio.Pin) rpio.State {
	// for per bit: first 50us low, then 26-70us high
	// 26-28us high represents 0
	// 70us high represents 1
	wait := 0
	for wait < 100 {
		if v := p.Read(); v == rpio.Low {
			break
		}
		util.DelayMicroseconds(1)
		wait += 1
	}

	wait = 0
	for wait < 100 {
		if v := p.Read(); v == rpio.High {
			break
		}
		util.DelayMicroseconds(1)
		wait += 1
	}

	util.DelayMicroseconds(40)

	return p.Read()
}

func checkData(data []uint8) bool {
	sum := 0
	for _, v := range data[:4] {
		sum += int(v)
	}

	return sum == int(data[4])
}
