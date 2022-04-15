package router

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	plg "github.com/99designs/gqlgen/graphql/playground"
	"github.com/SPSOAFM-IT18/dmp-plant-hub/database"
	"github.com/SPSOAFM-IT18/dmp-plant-hub/env"
	"github.com/SPSOAFM-IT18/dmp-plant-hub/graph"
	"github.com/SPSOAFM-IT18/dmp-plant-hub/graph/generated"
	mid "github.com/SPSOAFM-IT18/dmp-plant-hub/rest/middleware"
	sens "github.com/SPSOAFM-IT18/dmp-plant-hub/sensors"
	"github.com/go-chi/chi"
	webs "github.com/gorilla/websocket"
	"github.com/rs/cors"
)

func Router(db *database.DB, sensei *sens.Sensors) *chi.Mux {
	r := chi.NewRouter()

	// Add CORS middleware around every request
	// See https://github.com/rs/cors for full option listing
	r.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{env.Process("CORS")},
		AllowCredentials: true,
		Debug:            false,
	}).Handler)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{DB: db}}))

	srv.AddTransport(&transport.Websocket{
		Upgrader: webs.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				// Check against your desired domains here
				return r.Host == env.Process("CORS")
			},
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
	})

	r.Handle("/graphql", plg.Handler("GraphQL playground", "/query"))
	r.Handle("/query", srv)

	r.MethodFunc("GET", "/init/measured", mid.HandleGetInitMeasured)
	r.MethodFunc("POST", "/init/measured", mid.HandlePostInitMeasured)

	r.MethodFunc("GET", "/live/measure", mid.HandleGetLiveMeasure)
	r.MethodFunc("POST", "/live/measure", mid.HandlePostLiveMeasure)

	r.MethodFunc("GET", "/live/notify", mid.HandleGetLiveNotify)
	r.MethodFunc("POST", "/live/notify", mid.HandlePostLiveNotify)

	r.MethodFunc("GET", "/live/control", mid.HandleGetLiveControl)
	r.MethodFunc("POST", "/live/control", mid.HandlePostLiveControl)

	r.MethodFunc("GET", "/api/weather", mid.HandleGetWeather)

	r.MethodFunc("GET", "/api/geocodes", mid.HandleGetGeocode)
	r.MethodFunc("POST", "/api/geocodes", mid.HandlePostGeocode)

	r.MethodFunc("GET", "/api/google", mid.HandleGetGoogle)
	r.MethodFunc("POST", "/api/google", mid.HandlePostGoogle)

	mid.LoadInstances(db, sensei)

	http.Handle("/", r)

	return r
}
