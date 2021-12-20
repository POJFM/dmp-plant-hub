package sensors

import (
	"fmt"
	"log"
	"time"

	"github.com/SPSOAFM-IT18/dmp-plant-hub/sensors/dht"
	"github.com/stianeikeland/go-rpio/v4"
)

type PinOut struct {
	TRIG  rpio.Pin
	ECHO  rpio.Pin
	MOIST rpio.Pin
	DHT   rpio.Pin
	PUMP  rpio.Pin
	LED   rpio.Pin
}

func Pins() *PinOut {
	var err error = rpio.Open()
	if err != nil {
		panic(fmt.Sprint("unable to open gpio", err.Error()))
	}
	//defer rpio.Close()

	p := PinOut{
		TRIG:  rpio.Pin(2),
		ECHO:  rpio.Pin(3),
		MOIST: rpio.Pin(22),
		DHT:   rpio.Pin(23),
		PUMP:  rpio.Pin(18),
		LED:   rpio.Pin(27),
	}

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

func ReadDHT() (temp float32, hum float32, retried int) {
	temp, hum, retried, err := dht.ReadDHTxxWithRetry(dht.DHT11, 23, false, 10)
	if err != nil {
		log.Fatal(err)
	}
	return temp, hum, retried
}

func (p *PinOut) ReadWaterLevel() (waterLevel float32) {
	startTime := time.Now().UnixNano()
	stopTime := time.Now().UnixNano()
	p.TRIG.High()
	time.Sleep(10 * time.Microsecond)
	p.TRIG.Low()
	startTime = time.Now().UnixNano()
	log.Println("startTime is ", startTime)
	for p.ECHO.Read() == 0 {
		log.Println("echo is ", p.ECHO.Read())
	}
	log.Println("echo is 1")
	stopTime = time.Now().UnixNano()
	log.Println("stopTime is ", stopTime)
	duration := time.Duration(stopTime - startTime).Seconds()
	log.Println("duration: ", duration)
	//return (float32(time.Duration(stopTime - startTime) * time.Microsecond) * 34300) / 2
	return float32(duration) * 34300 / 2
}
