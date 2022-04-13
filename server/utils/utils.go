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
	log.Println("Interrupt signal caught")
	Exit()
}

func WaitTillWholeHour() {
	for time.Now().Format("04") != "00" {
		time.Sleep(1 * time.Minute)
	}
}

func Exit() {
	log.Println("Shutting down...")
	log.Println("Attempting to clean GPIO...")
	err := rpio.Close()
	if err != nil {
		log.Fatalln("Unable to clean GPIO")
	}
	log.Println("GPIO cleaned successfully! Exiting..")
	os.Exit(0)
}
