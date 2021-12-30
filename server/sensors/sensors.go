package sensors

import (
	"fmt"
	"log"
	"time"

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

type Measurements struct {
	Hum            float32 `json:"hum"`
	Temp           float32 `json:"temp"`
	Moist          float32 `json:"moist"`
	WithIrrigation float32 `json:"with_irrigation"`
}

func Pins() *PinOut {
	err := rpio.Open()
	if err != nil {
		panic(fmt.Sprint("unable to open gpio", err.Error()))
	}

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
	//p.MOIST.Input()
	p.DHT.Input()
	p.PUMP.Output()
	p.LED.Output()

	p.TRIG.Low()

	return &p
}

func (p *PinOut) MeasureAsync(c chan<- Measurements) {
	for range time.Tick(time.Second * 1) {
		c <- p.Measure()
	}
}

func (p *PinOut) Measure() Measurements {
	return Measurements{
		Hum:            0,
		Temp:           0,
		Moist:          0,
		WithIrrigation: 0,
	}
}

/*func ReadDHT() (temp float32, hum float32, retried int) {
	temp, hum, retried, err := dht.ReadDHTxxWithRetry(dht.DHT11, 23, false, 10)
	if err != nil {
		log.Fatal(err)
	}
	return temp, hum, retried
}*/

func (p *PinOut) ReadMoisture() (moisture []byte) {
	rpio.SpiBegin(rpio.Spi2)
	bytes := rpio.SpiReceive(10)
	rpio.SpiEnd(rpio.Spi2)
	return bytes
}

func (p *PinOut) ReadWaterLevel() (waterLevel float32) {
	startTime := time.Now().UnixNano()
	stopTime := time.Now().UnixNano()
	p.TRIG.High()
	time.Sleep(10 * time.Microsecond)
	p.TRIG.Low()
	startTime = time.Now().UnixNano()
	//log.Println("startTime is ", startTime)
	for p.ECHO.Read() == 0 {
		log.Println("echo is ", p.ECHO.Read())
	}
	//log.Println("echo is 1")
	stopTime = time.Now().UnixNano()
	//log.Println("stopTime is ", stopTime)
	duration := time.Duration(stopTime - startTime).Seconds()
	//log.Println("duration: ", duration)
	//return (float32(time.Duration(stopTime - startTime) * time.Microsecond) * 34300) / 2
	return float32(duration) * 34300 / 2
}
