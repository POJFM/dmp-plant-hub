package utils

import (
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/stianeikeland/go-rpio/v4"
)

func ArithmeticMean(list []float64) float64 {
	// maybe make it into list map function
	var total float64
	for _, v := range list {
		total += v
	}
	return total / float64(len(list))
}

func CatchInterrupt() {
	sigchan := make(chan os.Signal)
	signal.Notify(sigchan, os.Interrupt)
	<-sigchan
	log.Println("Program killed.. cleaning GPIO")
	err := rpio.Close()
	if err != nil {
		log.Fatalln("Unable to clean GPIO")
	}
	os.Exit(0)
}

func WaitTillWholeHour() {
	for time.Now().Format("04") != "00" {
		time.Sleep(1 * time.Minute)
	}
}
