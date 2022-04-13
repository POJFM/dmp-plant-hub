package main

import (
	"github.com/SPSOAFM-IT18/dmp-plant-hub/utils"
	"log"
	"net/http"

	"github.com/SPSOAFM-IT18/dmp-plant-hub/database"
	"github.com/SPSOAFM-IT18/dmp-plant-hub/env"

	sens "github.com/SPSOAFM-IT18/dmp-plant-hub/sensors"
	seq "github.com/SPSOAFM-IT18/dmp-plant-hub/sequences"

	r "github.com/SPSOAFM-IT18/dmp-plant-hub/router"
)

func main() {
	go utils.CatchInterrupt()
	db := database.Connect()
	sensei := sens.Init()
	sensei.StopLED()
	sensei.StopPump()

	// for {
	// 	measurement := sensei.Measure()
	// 	wl := sensei.ReadWaterLevel()
	// 	fmt.Printf("temp: %f\nhum: %f\nmoi: %f\nwl: %f\n", measurement.Temp, measurement.Hum, measurement.Moist, wl)
	// 	time.Sleep(2 * time.Second)
	// }

	go seq.Controller(db, sensei)

	log.Fatal(http.ListenAndServe(":"+env.Process("GO_API_PORT"), r.Router(db, sensei)))
}
