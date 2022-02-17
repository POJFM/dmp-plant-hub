package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/SPSOAFM-IT18/dmp-plant-hub/test/adc"
	"github.com/SPSOAFM-IT18/dmp-plant-hub/test/env"
	"github.com/SPSOAFM-IT18/dmp-plant-hub/test/model"
	"github.com/SPSOAFM-IT18/dmp-plant-hub/test/requests"
	"github.com/SPSOAFM-IT18/dmp-plant-hub/test/router"
	"github.com/SPSOAFM-IT18/dmp-plant-hub/test/sequences"
	"github.com/jasonlvhit/gocron"
)

func kokoti() {
	requests.PostLiveNotify(model.LiveNotify{Title: "Zavlažování", State: "inProgress", Action: "Probíhá zavlažování"})
	requests.PostLiveControl(model.LiveControl{Restart: false, PumpState: false})
}

func myTask() {
	hours := time.Now().Format("04")

	fmt.Println(hours)
}

func executeCronJob() {
	for time.Now().Format("04") != "39" {
		fmt.Println(time.Now().Format("04") != "13")
		time.Sleep(1 * time.Minute)
	}
	gocron.Every(4).Second().Do(myTask)
	<-gocron.Start()
}

func kokot() {
	for {
		time.Sleep(1 * time.Second)
		// var kokotismus = true

		// if kokotismus == true {
		// 	kokotismus = false
		// 	time.AfterFunc(3*time.Second, func() {
		// 		hours := time.Now().format
		// 		fmt.Println(hours)
		// 		kokotismus = true
		// 	})
		// }
		go kokoti()
		//go requests.PostLiveNotify(model.LiveNotify{Title: "kokot jsi", State: "active", Action: "debil"})
	}
}

func main() {
	temp := make(chan float32)

	go sequences.MeasurementSequence(temp)

	sequences.SaveOnFourHoursPeriod(temp)

	r := router.Router()

	adc.Adc()

	// port := fmt.Sprint(":" + env.Process("GO_API_PORT"))
	// fmt.Print(port)
	fmt.Println("Starting server on the port", env.Process("GO_API_PORT"))
	go kokot()
	go executeCronJob()
	log.Fatal(http.ListenAndServe(":"+env.Process("GO_API_PORT"), r))
}
