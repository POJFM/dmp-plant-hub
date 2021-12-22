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

type kokotak struct {
	*sensors.PinOut
}

func main() {
	//Pins()
	//InitializationSequence()

	// Wait for init to finish

	//MeasurementSequence(Pins.PUMP, Pins.LED)

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
