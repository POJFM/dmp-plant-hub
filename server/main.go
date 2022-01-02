package main

import (
	"fmt"
	"github.com/SPSOAFM-IT18/dmp-plant-hub/sensors"
	"gobot.io/x/gobot/drivers/spi"
	"gobot.io/x/gobot/platforms/raspi"
)

const defaultPort = "5000"

type kokotak struct {
	*sensors.PinOut
}

type measurements struct {
	waterLevel float32
}

var liveMeasurements sensors.Measurements

func main() {

	a := raspi.NewAdaptor()
	adc := spi.NewMCP3008Driver(a)
	fmt.Println(adc.Read(0))

	/*var sens = sensors.Pins()
	//sequences.InitializationSequence()
	//sequences.MeasurementSequence(kokot.PUMP, kokot.LED)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

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
		w.Header().Set("Access-Control-Allow-Origin", "*")
		json.NewEncoder(w).Encode(liveMeasurements)
	})

	http.HandleFunc("/init/measure", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		json.NewEncoder(w).Encode(initMeasurements)
	})

	//var foo = kokot.ReadWaterLevel()

	var db = database.Connect()

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{DB: db}}))

	http.Handle("/graphql", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))*/
}
