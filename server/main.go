package main

import (
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/SPSOAFM-IT18/dmp-plant-hub/database"
	"github.com/SPSOAFM-IT18/dmp-plant-hub/graph"
	"github.com/SPSOAFM-IT18/dmp-plant-hub/graph/generated"
	"github.com/SPSOAFM-IT18/dmp-plant-hub/sensors"
	"github.com/SPSOAFM-IT18/dmp-plant-hub/sensors/dht/drivers"
	"github.com/go-chi/chi"
	"github.com/gorilla/websocket"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

const defaultPort = "5000"

var liveMeasurements sensors.Measurements

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

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	doDHT11()
	/*	//s := sensors.Init()
		err := dht.HostInit()
		if err != nil {
			fmt.Println("HostInit error:", err)
			return
		}

		thisdht, err := dht.NewDHT("GPIO23", dht.Fahrenheit, "DHT11")
		if err != nil {
			fmt.Println("NewDHT error:", err)
			return
		}

		humidity, temperature, err := thisdht.ReadRetry(11)
		if err != nil {
			fmt.Println("Read error:", err)
			return
		}

		fmt.Printf("humidity: %v\n", humidity)
		fmt.Printf("temperature: %v\n", temperature)*/
	for {
		/*		temperature, humidity, retried, err :=
					dht.ReadDHTxxWithRetry(dht.DHT11, 23, false, 10)
				if err != nil {
					log.Fatal(err)
				}
				// Print temperature and humidity
				fmt.Printf("Temperature = %v*C, Humidity = %v%% (retried %d times)\n",
					temperature, humidity, retried)
				//fmt.Println(s.ReadDHT())
				time.Sleep(1 * time.Second)*/
		/*		err := dht.HostInit()
				if err != nil {
					fmt.Println("HostInit error:", err)
					return
				}

				dht, err := dht.NewDHT("GPIO23", dht.Fahrenheit, "DHT11")
				if err != nil {
					fmt.Println("NewDHT error:", err)
					return
				}

				humidity, temperature, err := dht.Read()
				if err != nil {
					fmt.Println("Read error:", err)
					return
				}

				fmt.Printf("humidity: %v\n", humidity)
				fmt.Printf("temperature: %v\n", temperature)*/

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

	log.Printf("connect to http://localhost:%s/graphql for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
