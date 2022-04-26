package main

import (
	"log"
	"net/http"

	"github.com/SPSOAFM-IT18/dmp-plant-hub/utils"

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

	go seq.Controller(db, sensei)

	log.Fatal(http.ListenAndServe(":"+env.Process("GO_API_PORT"), r.Router(db, sensei)))
}
