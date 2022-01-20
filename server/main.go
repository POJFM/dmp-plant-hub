package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/SPSOAFM-IT18/dmp-plant-hub/database"
	"github.com/SPSOAFM-IT18/dmp-plant-hub/env"
	"github.com/SPSOAFM-IT18/dmp-plant-hub/graph"
	"github.com/SPSOAFM-IT18/dmp-plant-hub/graph/generated"
	"github.com/SPSOAFM-IT18/dmp-plant-hub/sensors"
	"github.com/stianeikeland/go-rpio/v4"
)

type kokotak struct {
	*sensors.PinOut
}

type measurements struct {
	waterLevel float32
}

var liveMeasurements sensors.Measurements

const pin_PUMP = 18
const pin_LED = 27

func main() {
	// //@CHECK FOR DATA IN DB
	// if (data in settings table) {
	// 	IrrigationSequence(pin_PUMP, pin_LED)
	// } else {
	// 	InitializationSequence()
	// }

	go MeasurementSequence(pin_PUMP, pin_LED)

	var sens = sensors.Pins()
	//sequences.InitializationSequence()
	//sequences.MeasurementSequence(kokot.PUMP, kokot.LED)

	fmt.Println(sens.ReadMoisture())

	osSignalChannel := make(chan os.Signal, 1)
	signal.Notify(osSignalChannel, os.Interrupt)
	go func() {
		for range osSignalChannel {
			rpio.Close()
		}
	}()

	c := make(chan sensors.Measurements)
	go sens.MeasureAsync(c)

	go func() {
		for {
			liveMeasurements = <-c
		}
	}()

	initMeasurements := liveMeasurements

	http.HandleFunc("/live/measure", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", env.Process("CORS"))

		json.NewEncoder(w).Encode(liveMeasurements)
	})

	http.HandleFunc("/init/measure", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", env.Process("CORS"))

		json.NewEncoder(w).Encode(initMeasurements)
	})

	/*go func() {
		for {
			measurement := <-c
			fmt.Println(measurement)
		}
	}()*/

	//var foo = kokot.ReadWaterLevel()

	var db = database.Connect()

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{DB: db}}))

	http.Handle("/graphql", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", env.Process("GO_API_PORT"))
	log.Fatal(http.ListenAndServe(":"+env.Process("GO_API_PORT"), nil))
}
