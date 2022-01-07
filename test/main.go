package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/SPSOAFM-IT18/dmp-plant-hub/test/env"
	"github.com/SPSOAFM-IT18/dmp-plant-hub/test/model"
	"github.com/SPSOAFM-IT18/dmp-plant-hub/test/requests"
	"github.com/SPSOAFM-IT18/dmp-plant-hub/test/router"
)

func kokoti() {
	requests.PostLiveNotify(model.LiveNotify{Title: "Zavlažování", State: "inProgress", Action: "Probíhá zavlažování"})
	requests.PostLiveControl(model.LiveControl{Restart: false, PumpState: false})
}

func kokot() {
	for {
		time.Sleep(1 * time.Second)
		go kokoti()
		//go requests.PostLiveNotify(model.LiveNotify{Title: "kokot jsi", State: "active", Action: "debil"})
	}
}

func main() {
	// až bude internet
	//go get github.com/joho/godotenv
	r := router.Router()

	// port := fmt.Sprint(":" + env.Process("GO_API_PORT"))
	// fmt.Print(port)
	fmt.Println("Starting server on the port", env.Process("GO_API_PORT"))
	go kokot()
	log.Fatal(http.ListenAndServe(":"+env.Process("GO_API_PORT"), r))
}
