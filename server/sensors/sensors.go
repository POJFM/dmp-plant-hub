package sensors

import (
	"fmt"
	"github.com/SPSOAFM-IT18/dmp-plant-hub/sensors/drivers"
	"log"
	"os"
	"time"

	"github.com/shanghuiyang/rpi-devices/dev"
	"github.com/stianeikeland/go-rpio/v4"
)

// Pins
const TRIG = 2
const ECHO = 3
const DHT = 23
const PUMP = rpio.Pin(18)
const LED = rpio.Pin(27)

type PinOut struct {
	TRIG  rpio.Pin
	ECHO  rpio.Pin
	MOIST rpio.Pin
	DHT   rpio.Pin
	PUMP  rpio.Pin
	LED   rpio.Pin
}

type Sensors struct {
	sonic *dev.HCSR04
	dht   *dev.DHT11
	mcp   *drivers.MCP3008
}

type Measurements struct {
	Hum            float32 `json:"hum"`
	Temp           float32 `json:"temp"`
	Moist          float32 `json:"moist"`
	WithIrrigation float32 `json:"with_irrigation"`
}

func Init() *Sensors {

	/*p := PinOut{
		TRIG:  rpio.Pin(2),
		ECHO:  rpio.Pin(3),
		MOIST: rpio.Pin(22),
		DHT:   rpio.Pin(23),
		PUMP:  rpio.Pin(18),
		LED:   rpio.Pin(27),
	}*/

	// TODO: close connections on exit/interrupt
	return &Sensors{
		sonic: dev.NewHCSR04(TRIG, ECHO),
		dht:   dev.NewDHT11(),
		//mcp:   (drivers.NewMCP3008(0, 0, 3.3), err),
	}
}

func (s *Sensors) MeasureAsync(c chan<- Measurements) {
	for range time.Tick(time.Second * 1) {
		c <- s.Measure()
	}
}

func (s *Sensors) Measure() Measurements {
	return Measurements{
		Hum:            0,
		Temp:           0,
		Moist:          0,
		WithIrrigation: 0,
	}
}

func StartPump() {
	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer func() {
		err := rpio.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()

	kokot := rpio.Pin(18)
	kokot.Output()
	kokot.High()

	pin := rpio.Pin(27)
	pin.Output()
	pin.High()
}

func (s *Sensors) StopPump() {
	PUMP.Low()
}

/*func (s *Sensors) ReadDHT() (temp, hum float32) {
	/*	temp, hum, err := s.dht.TempHumidity()
		if err != nil {
			log.Fatalf("failed to read from DHT11, error: %v", err)
		}
		log.Printf("t = %.0f, h = %.0f%%", temp, hum)
	temperature, humidity, retried, err :=
		dht.ReadDHTxxWithRetry(dht.DHT11, 23, false, 10)
	if err != nil {
		log.Fatal(err)
	}
	// Print temperature and humidity
	fmt.Printf("Temperature = %v*C, Humidity = %v%% (retried %d times)\n",
		temperature, humidity, retried)
	return temperature, humidity
}*/

func (s *Sensors) ReadMoisture() (moisture float64) {
	mcp, err := drivers.NewMCP3008(0, 0, 3.3)
	if err != nil {
		log.Fatalf("mcp3008 failed: %s", err)
	}
	moisture = mcp.ReadAdc()
	return moisture
}

func (s *Sensors) ReadMoisture0() (moisture float64) {
	mcp, err := drivers.NewMCP30008(0, 0, mcp.Mode0, 500000)

	if err != nil {
		log.Fatalf("mcp3008 failed: %s", err)
	}

	return mcp.Measure(0)
}

func (s *Sensors) ReadWaterLevel() (waterLevel float64) {
	waterLevel, err := s.sonic.Dist()
	if err != nil {
		log.Fatalf("failed to read waterLevel, error: %v", err)
	}
	return waterLevel
}
