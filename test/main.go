package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/SPSOAFM-IT18/dmp-plant-hub/test/env"
	mid "github.com/SPSOAFM-IT18/dmp-plant-hub/test/middleware"
	"github.com/SPSOAFM-IT18/dmp-plant-hub/test/router"
	//mid "github.com/SPSOAFM-IT18/dmp-plant-hub/test/middleware"
)

// func kokoti() {
// 	requests.PostLiveNotify(model.LiveNotify{Title: "Zavlažování", State: "inProgress", Action: "Probíhá zavlažování"})
// 	requests.PostLiveControl(model.LiveControl{Restart: false, PumpState: false})
// }

// func myTask() {
// 	hours := time.Now().Format("04")

// 	fmt.Println(hours)
// }

// func executeCronJob() {
// 	for time.Now().Format("04") != "39" {
// 		fmt.Println(time.Now().Format("04") != "13")
// 		time.Sleep(1 * time.Minute)
// 	}
// 	gocron.Every(4).Second().Do(myTask)
// 	<-gocron.Start()
// }

func kokot(cRestart, cPumpState chan bool) {
	//requests.PostLiveControl(model.LiveControl{Restart: false, PumpState: false})

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
		//requests.PostLiveNotify(model.LiveNotify{Title: "Zavlažování", State: "inProgress", Action: "Probíhá zavlažování"})
		//requests.GetLiveControl()

		mid.GetLiveControl(cRestart, cPumpState)

		fmt.Println("\n", <-cRestart)
		fmt.Println("", <-cPumpState)

		//go kokoti()
		//go requests.PostLiveNotify(model.LiveNotify{Title: "kokot jsi", State: "active", Action: "debil"})
	}
}

func main() {
	cRestart := make(chan bool, 1)
	cPumpState := make(chan bool, 1)

	r := router.Router()

	t0 := time.Now()
	time.Sleep(5 * time.Second)

	fmt.Println(int(time.Since(t0).Seconds()))

	//go sequences.MeasurementSequence(temp)

	// adc.Adc()

	// port := fmt.Sprint(":" + env.Process("GO_API_PORT"))
	// fmt.Print(port)
	fmt.Println("Starting server on the port", env.Process("GO_API_PORT"))
	go kokot(cRestart, cPumpState)
	//go executeCronJob()
	log.Fatal(http.ListenAndServe(":"+env.Process("GO_API_PORT"), r))
}
