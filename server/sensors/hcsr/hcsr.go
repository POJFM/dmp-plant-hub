package hcsr

import (
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
	port, err := serial.Open(hc.portName, hc.mode)
	sendSignal(port)
	time.Sleep(100 * time.Microsecond)
	buff, n := readSerial(port)
	//dist, _ := strconv.ParseFloat(strings.TrimSuffix(string(buff[:n]), "\r\n"), 64)
	dist, jsikokot := strconv.ParseFloat(string(buff[:n]), 64)
	if jsikokot != nil {
		log.Fatal(jsikokot)
	}
	_ = port.Close()
	return dist, err
}

func sendSignal(port serial.Port) {
	n, err := port.Write([]byte("1"))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Sonic signal sent. %v bytes \n", n)
}

func readSerial(port serial.Port) (buff []byte, n int) {
	buff = make([]byte, 100)
	n, err := port.Read(buff)
	if err != nil {
		log.Fatal(err)
	}
	return
}
