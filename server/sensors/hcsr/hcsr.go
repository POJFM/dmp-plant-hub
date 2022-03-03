// Package hcsr
// https://github.com/shanghuiyang/rpi-devices/blob/master/dev/hcsr04.go
/*
HC-SR04 is an ultrasonic distance meter used to measure the distance to objects.
Spec:
  - power supply:	+5V DC
  - range:			2 - 450cm
  - resolution:		0.3cm
	 ___________________________
    |                           |
    |          HC-SR04          |
    |                           |
    |___________________________|
         |     |     |     |
        vcc  trig   echo  gnd
Connect to Raspberry Pi:
  - vcc:	any 5v pin
  - gnd:	any gnd pin
  - trig:	any data pin
  - echo:	any data pin
*/
package hcsr

import (
	"time"

	"github.com/stianeikeland/go-rpio/v4"
)

const (
	timeout    = 3600
	voiceSpeed = 34000.0
)

// HCSR04 implements DistanceMeter interface
type HCSR04 struct {
	trig rpio.Pin
	echo rpio.Pin
}

// NewHCSR04 ...
func NewHCSR04(trig int8, echo int8) *HCSR04 {
	hc := &HCSR04{
		trig: rpio.Pin(trig),
		echo: rpio.Pin(echo),
	}
	hc.trig.Output()
	hc.trig.Low()
	hc.echo.Input()
	return hc
}

// Dist
// Value returns distance in cm to objects
func (hc *HCSR04) Dist() (float64, error) {
	hc.trig.Low()
	time.Sleep(100 * time.Microsecond)
	hc.trig.High()
	time.Sleep(15 * time.Microsecond)

	for n := 0; n < timeout && hc.echo.Read() != rpio.High; n++ {
		time.Sleep(1 * time.Microsecond)
	}
	start := time.Now()

	for n := 0; n < timeout && hc.echo.Read() != rpio.Low; n++ {
		time.Sleep(1 * time.Microsecond)
	}
	return time.Since(start).Seconds() * voiceSpeed / 2.0, nil
}
