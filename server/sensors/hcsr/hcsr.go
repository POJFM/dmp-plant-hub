package hcsr

import (
	"fmt"
	"go.bug.st/serial"
	"log"
	"strconv"
	"time"
)

// HCSR04 implements DistanceMeter interface
type HCSR04 struct {
	mode     *serial.Mode
	portName string
	port     serial.Port
}

// NewHCSR04 ...
func NewHCSR04(portName string, baudRate int) *HCSR04 {
	mode := &serial.Mode{
		BaudRate: baudRate,
	}
	port, err := serial.Open(portName, mode)
	if err != nil {
		log.Println(err)
	}
	return &HCSR04{
		mode:     mode,
		portName: portName,
		port:     port,
	}
}

// Dist
// Value returns distance in cm to objects
func (hc *HCSR04) Dist() (float64, error) {
	// hc.openPort()
	time.Sleep(2 * time.Second)
	hc.sendSignal()
	log.Println("wait 100 ms")
	//time.Sleep(100 * time.Microsecond)
	log.Println("done waiting, read")
	//go func() {
	//		time.Sleep(2 * time.Second)
	//		hc.sendSignal()
	//	}()
	buff, n := hc.readSerial()
	log.Println("data read")
	//dist, _ := strconv.ParseFloat(strings.TrimSuffix(string(buff[:n]), "\r\n"), 64)
	dist, err := strconv.ParseFloat(string(buff[:n]), 64)
	log.Println("read data: " + fmt.Sprintf("%f", dist))
	if err != nil {
		log.Println(err)
	}
	log.Println("close port")
	// err = hc.port.Close()

	return dist, err
}

func (hc *HCSR04) openPort() {
	var err error
	hc.port, err = serial.Open(hc.portName, hc.mode)
	if err != nil {
		log.Printf("Failed to open serial: %v\n", err)
	}
}

func (hc *HCSR04) sendSignal() {
	n, err := hc.port.Write([]byte("1"))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Sonic signal sent. %v bytes \n", n)
}

func (hc *HCSR04) readSerial() (buff []byte, n int) {
	buff = make([]byte, 100)
	n, err := hc.port.Read(buff)
	if err != nil {
		log.Fatal(err)
	}
	return
}
