package router

import (
	"github.com/SPSOAFM-IT18/dmp-plant-hub/kokot/kokot/middleware"

	"github.com/gorilla/mux"
)

// Router is exported and used in main.go
func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/live/control", middleware.PostLiveControl).Methods("POST", "OPTIONS")
	return router
}
