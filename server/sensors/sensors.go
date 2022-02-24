package sensors

import (
	"github.com/SPSOAFM-IT18/dmp-plant-hub/sensors/dht"
	"github.com/shanghuiyang/rpi-devices/dev"
	"github.com/stianeikeland/go-rpio/v4"
	"log"
	"time"
)

// Pins
const TRIG = 2
const ECHO = 3
const DHT = 23
const PUMP = rpio.Pin(18)
const LED = rpio.Pin(27)

type Sensors struct {
	sonic *dev.HCSR04
	dht   *dht.DHT11
}

type Measurements struct {
	Hum            float64 `json:"hum"`
	Temp           float64 `json:"temp"`
	Moist          float64 `json:"moist"`
	WithIrrigation float64 `json:"with_irrigation"`
}

func Init() *Sensors {
	if err := rpio.Open(); err != nil {
		log.Fatalf("Failed to open GPIO: %v", err)
	}

	PUMP.Output()

	return &Sensors{
		sonic: dev.NewHCSR04(TRIG, ECHO),
		dht:   dht.NewDHT11(DHT),
	}
}

func (s *Sensors) MeasureAsync(c chan<- Measurements) {
	for range time.Tick(time.Second * 1) {
		c <- s.Measure()
	}
}

func (s *Sensors) Measure() Measurements {
	hum, temp := s.ReadDHT()
	return Measurements{
		Hum:            hum,
		Temp:           temp,
		Moist:          s.ReadMoisture(),
		WithIrrigation: 0,
	}
}

func (s *Sensors) StartPump() {
	PUMP.High()
	log.Printf("PUMP started")
}

func (s *Sensors) StopPump() {
	PUMP.Low()
	log.Printf("PUMP stopped")
}

func (s *Sensors) ReadDHT() (hum, temp float64) {
	temp, hum, err := s.dht.ReadData()
	if err != nil {
		log.Printf("DHT11 Error: %v", err)
	}
	return
}

func (s *Sensors) ReadMoisture() (moisture float64) {
	if err := rpio.SpiBegin(rpio.Spi0); err != nil {
		log.Printf("MOISTURE: failed to start SPI: %v\n", err)
	}

	rpio.SpiSpeed(1000000)
	rpio.SpiChipSelect(0)

	channel := byte(0)

	data := []byte{1, (8 + channel) << 4, 0}

	rpio.SpiExchange(data)

	res := int(data[1]&3)<<8 + int(data[2])

	moisture = float64(res)
	// TODO: map moisture value to percentage
	//Vdd and Vref are at 5v. Change *5 to *3.3 if you are
	//powering the chip with 3.3v
	//voltage := (float32(code) * 5) / 1024

	rpio.SpiEnd(rpio.Spi0)
	return
}

func (s *Sensors) ReadWaterLevel() (waterLevel float64) {
	waterLevel, err := s.sonic.Dist()
	if err != nil {
		log.Printf("SONIC Error: %v", err)
	}
	return
}
