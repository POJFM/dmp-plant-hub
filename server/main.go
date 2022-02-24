package main

import (
	"github.com/SPSOAFM-IT18/dmp-plant-hub/env"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	plg "github.com/99designs/gqlgen/graphql/playground"
	"github.com/SPSOAFM-IT18/dmp-plant-hub/database"
	"github.com/SPSOAFM-IT18/dmp-plant-hub/graph"
	"github.com/SPSOAFM-IT18/dmp-plant-hub/graph/generated"
	"github.com/go-chi/chi"
	webs "github.com/gorilla/websocket"
	"github.com/rs/cors"
)

func main() {

	//cMoist := make(chan float32)
	//cTemp := make(chan float32)
	//cHum := make(chan float32)

	//go seq.MeasurementSequence(cMoist, cTemp, cHum)

	//go seq.SaveOnFourHoursPeriod(cMoist, cTemp, cHum)

	// //@CHECK FOR DATA IN DB
	// if DATA_IN_DB {
	// 	go seq.IrrigationSequence(cMoist)
	// } else {
	// 	go seq.InitializationSequence(cMoist)
	// 	initializationFinished := true
	// 	for initializationFinished {
	// 		stopLED := make(chan bool)
	// 		go func() {
	// 			for {
	// 				select {
	// 				case <-stopLED:
	// 					return
	// 				default:
	// 					for i := 0; i < 2; i++ {
	// 						sens.LED.High()
	// 						time.Sleep(500 * time.Millisecond)
	// 						sens.LED.Low()
	// 						time.Sleep(500 * time.Millisecond)
	// 					}
	// 					time.Sleep(1500 * time.Millisecond)
	// 				}
	// 			}
	// 		}()
	// 		if DATA_IN_DB {
	// 			initializationFinished = false
	// 			stopLED <- true
	// 			go seq.IrrigationSequence(cMoist)
	// 		}
	// 		time.Sleep(1000 * time.Millisecond)
	// 	}
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
		Upgrader: webs.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				// Check against your desired domains here
				return r.Host == "http://4.2.0.126:3000"
			},
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
	})

	router.Handle("/graphql", plg.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)
	/*http.Handle("/graphql", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)*/

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", env.Process("GO_API_PORT"))
	log.Fatal(http.ListenAndServe(":"+env.Process("GO_API_PORT"), router))
}
