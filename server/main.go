package main

import (
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/SPSOAFM-IT18/dmp-plant-hub/database"
	"github.com/SPSOAFM-IT18/dmp-plant-hub/env"
	"github.com/SPSOAFM-IT18/dmp-plant-hub/graph"
	"github.com/SPSOAFM-IT18/dmp-plant-hub/graph/generated"
	drivers2 "github.com/SPSOAFM-IT18/dmp-plant-hub/sensors/drivers"
	"github.com/go-chi/chi"
	"github.com/gorilla/websocket"
	"github.com/rs/cors"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/SPSOAFM-IT18/dmp-plant-hub/sensors"
)

var liveMeasurements sensors.Measurements

const pin_PUMP = 18
const pin_LED = 27

func doDHT11() {

	drivers2.OpenRPi()
	// 1、init
	dht11, err := drivers2.NewDHT11(23)
	if err != nil {
		fmt.Println("----------------------------------")
		fmt.Println(err)
		return
	}

	// 2、read
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second * 2)
		rh, tmp, err := dht11.ReadData()
		if err != nil {
			fmt.Println("----------------------------------")
			fmt.Println(err)
			continue
		}

		fmt.Println("----------------------------------")
		fmt.Println("RH:", rh)

		tmpStr := strconv.FormatFloat(tmp, 'f', 1, 64)
		fmt.Println("TMP:", tmpStr)

	}

	// 3、close
	err = dht11.Close()
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	sensei := sensors.Init()
	go func() {
		for {
			fmt.Println(sensei.ReadMoisture())
			time.Sleep(time.Second * 1)
		}
	}()

	/*moisture := make(chan float32)
	temperature := make(chan float32)
	humidity := make(chan float32)

	go sequences.MeasurementSequence(sensors.PUMP, sensors.LED, moisture, temperature, humidity)

	cMoisture := <-moisture
	cTemperature := <-temperature
	cHumidity := <-humidity

	sequences.SaveOnFourHoursPeriod(cMoisture, cTemperature, cHumidity)*/

	// //@CHECK FOR DATA IN DB
	// if (data in settings table) {
	// 	sequences.IrrigationSequence(pin_PUMP, pin_LED, cMoisture, cTemperature, cHumidity)
	// } else {
	// 	sequences.InitializationSequence(cMoisture, cTemperature, cHumidity)
	/*	initializationFinished := true
		for initializationFinished {
			// BLINK LED
			// if (data in settings table) {
			initializationFinished = false
			// 	sequences.IrrigationSequence(pin_PUMP, pin_LED, cMoisture, cTemperature, cHumidity)
			// }
		}*/
	// }

	/*
		c := make(chan os.Signal)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		go func() {
			<-c
			err := rpio.Close()
			if err != nil {
				log.Fatalf("failed to clean ")
			}
			os.Exit(1)
		}()
	*/

	var db = database.Connect()

	router := chi.NewRouter()

	// Add CORS middleware around every request
	// See https://github.com/rs/cors for full option listing
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"http://4.2.0.126:3000", "http://localhost:3000", "http://4.2.0.126", "http://4.2.0.225:5000"},
		AllowCredentials: true,
		Debug:            true,
	}).Handler)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{DB: db}}))

	srv.AddTransport(&transport.Websocket{
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				// Check against your desired domains here
				return r.Host == "http://4.2.0.126:3000"
			},
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
	})

	router.Handle("/graphql", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", env.Process("GO_API_PORT"))
	log.Fatal(http.ListenAndServe(":"+env.Process("GO_API_PORT"), router))
}
