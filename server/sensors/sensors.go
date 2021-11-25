package sensors

import (
	"fmt"
	"log"

	"github.com/SPSOAFM-IT18/dmp-plant-hub/sensors/dht"
	"github.com/stianeikeland/go-rpio"
)

type pinOut struct {
	TRIG  rpio.Pin
	ECHO  rpio.Pin
	MOIST rpio.Pin
	DHT   rpio.Pin
	PUMP  rpio.Pin
	LED   rpio.Pin
}

func Pins() *pinOut {
	p := pinOut{
		TRIG:  rpio.Pin(2),
		ECHO:  rpio.Pin(3),
		MOIST: rpio.Pin(22),
		DHT:   rpio.Pin(23),
		PUMP:  rpio.Pin(18),
		LED:   rpio.Pin(27),
	}
	var err error = rpio.Open()
	if err != nil {
		panic(fmt.Sprint("unable to open gpio", err.Error()))
	}
	defer rpio.Close()

	// IO
	p.TRIG.Output()
	p.ECHO.Input()
	p.MOIST.Input()
	p.DHT.Input()
	p.PUMP.Output()
	p.LED.Output()

	p.TRIG.Low()

	return &p
}

func ReadDHT() (temp float32, hum float32) {
	temp, hum, err := dht.ReadDHTxx(dht.DHT11, 23, false)
	if err != nil {
		log.Fatal(err)
	}
	return temp, hum
}
