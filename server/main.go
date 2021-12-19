package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/SPSOAFM-IT18/dmp-plant-hub/database"
	"github.com/SPSOAFM-IT18/dmp-plant-hub/graph"
	"github.com/SPSOAFM-IT18/dmp-plant-hub/graph/generated"
	"github.com/SPSOAFM-IT18/dmp-plant-hub/sensors"
)

const defaultPort = "5000"

// Web
// Objemový limit
var manualWaterOverdrawn float32

// Hladinový limit
var manualWaterLevel float32
var manualTemp float32
var manualHum float32
var initializationState bool = false // btn on web after init is completed

// Code
var initialization bool = false

// add reserve => not from bottom but from low water level
var waterLevel float32 // on init measures 5 times, appends the values into an array and then averages the values into single value
var moistureLevel float32
var waterOverdrawnLevel float32
var pumpFlow float32 // liter/min

type kokotak struct {
	*sensors.PinOut
}

func main() {
	//Pins()
	//InitializationSequence(manualWaterOverdrawn, manualWaterLevel, waterLevel, moistureLevel, waterOverdrawnLevel, initializationState, initialization)
	//MeasurementSequence(Pins.PUMP, Pins.LED, manualWaterOverdrawn, manualWaterLevel, waterLevel, moistureLevel, waterOverdrawnLevel, pumpFlow, initializationState, initialization)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	/*temp, hum, retried := sensors.ReadDHT()
	log.Printf("temp: %s, hum: %s, retried: %s", temp, hum, retried)
	kokot := 1
	for kokot < 100 {
		kokot += kokot
		kokotinec := sensors.Pins()
		log.Printf("jsi kokot: %v", sensors.PinOut.ReadWaterLevel(kokotinec))
	}*/

	var kokot = sensors.Pins()

	var foo = kokot.ReadWaterLevel()
	log.Println(foo)

	var db = database.Connect()

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{DB: db}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
