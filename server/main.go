package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/SPSOAFM-IT18/dmp-plant-hub/database"
	"github.com/SPSOAFM-IT18/dmp-plant-hub/env"
	"github.com/SPSOAFM-IT18/dmp-plant-hub/graph"
	"github.com/SPSOAFM-IT18/dmp-plant-hub/graph/generated"
	"github.com/SPSOAFM-IT18/dmp-plant-hub/sensors"
	"github.com/SPSOAFM-IT18/dmp-plant-hub/sensors/dht/drivers"
	"github.com/SPSOAFM-IT18/dmp-plant-hub/sequences"
	"github.com/go-chi/chi"
	"github.com/gorilla/websocket"
	"github.com/rs/cors"
)

var liveMeasurements sensors.Measurements

const pin_PUMP = 18
const pin_LED = 27

func doDHT11() {

	drivers.OpenRPi()
	// 1、init
	dht11, err := drivers.NewDHT11(23)
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
	doDHT11()

	cMoist := make(chan float32)
	cTemp := make(chan float32)
	cHum := make(chan float32)

	go sequences.MeasurementSequence(sensors.PUMP, sensors.LED, cMoist, cTemp, cHum)

	sequences.SaveOnFourHoursPeriod(cMoist, cTemp, cHum)

	// //@CHECK FOR DATA IN DB
	// if (data in settings table) {
	// 	sequences.IrrigationSequence(pin_PUMP, pin_LED, cMoisture, cTemperature, cHumidity)
	// } else {
	// 	sequences.InitializationSequence(cMoisture, cTemperature, cHumidity)
	// }

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
	/*http.Handle("/graphql", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)*/

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", env.Process("GO_API_PORT"))
	log.Fatal(http.ListenAndServe(":"+env.Process("GO_API_PORT"), router))
}
