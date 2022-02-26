package utils

import (
	"github.com/stianeikeland/go-rpio/v4"
	"log"
	"os"
	"os/signal"
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
