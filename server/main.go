package main

import (
	"log"
	"net/http"

	"github.com/SPSOAFM-IT18/dmp-plant-hub/database"
	"github.com/SPSOAFM-IT18/dmp-plant-hub/env"

	sens "github.com/SPSOAFM-IT18/dmp-plant-hub/sensors"
	seq "github.com/SPSOAFM-IT18/dmp-plant-hub/sequences"

	r "github.com/SPSOAFM-IT18/dmp-plant-hub/router"
)

func main() {
	cMoist := make(chan float64)
	cTemp := make(chan float64)
	cHum := make(chan float64)

	db := database.Connect()
	sensei := sens.Init()
	sensei.StopLED()
	sensei.StopPump()

	/* sensors test
	sensei := sens.Init()
	fmt.Println("Sensors initialized!")
	for {
		measurement := sensei.Measure()
		wl := sensei.ReadWaterLevel()
		fmt.Printf("temp: %f\nhum: %f\nmoi: %f\nwl: %f\n", measurement.Temp, measurement.Hum, measurement.Moist, wl)
		time.Sleep(2 * time.Second)
	}
	*/

	go seq.Controller(db, sensei, cMoist, cTemp, cHum)

	log.Fatal(http.ListenAndServe(":"+env.Process("GO_API_PORT"), r.Router(db, sensei)))
}
