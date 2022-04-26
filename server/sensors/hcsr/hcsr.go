package hcsr

import (
	"go.bug.st/serial"
	"log"
	"strconv"
	"strings"
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
	buff := make([]byte, 100)
	n, err := hc.port.Read(buff)
	dist, _ := strconv.ParseFloat(strings.TrimSuffix(string(buff[:n]), "\r\n"), 64)
	return dist, err
}
