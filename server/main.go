package main

import (
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/SPSOAFM-IT18/dmp-plant-hub/database"
	"github.com/SPSOAFM-IT18/dmp-plant-hub/graph"
	"github.com/SPSOAFM-IT18/dmp-plant-hub/graph/generated"
	"github.com/SPSOAFM-IT18/dmp-plant-hub/sensors"
	"log"
	"net/http"
	"os"
	"time"
)

const defaultPort = "5000"

var liveMeasurements sensors.Measurements

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	s := sensors.Init()
	for {
		fmt.Println(s.ReadWaterLevel())
		time.Sleep(1 * time.Second)
	}

	/*
		var sens = sensors.Pins()

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
			w.Header().Set("Access-Control-Allow-Origin", "*")
			json.NewEncoder(w).Encode(liveMeasurements)
		})

		http.HandleFunc("/init/measure", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			json.NewEncoder(w).Encode(initMeasurements)
		})
	*/

	var db = database.Connect()

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{DB: db}}))

	http.Handle("/graphql", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/graphql for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
