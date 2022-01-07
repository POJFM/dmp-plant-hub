package router

import (
	"net/http"

	"github.com/SPSOAFM-IT18/dmp-plant-hub/rest/middleware"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/init/measured", middleware.GetInitMeasured).Methods("GET", "OPTIONS")
	router.HandleFunc("/init/measured", middleware.PostInitMeasured).Methods("POST", "OPTIONS")

	router.HandleFunc("/live/measure", middleware.GetLiveMeasure).Methods("GET", "OPTIONS")
	router.HandleFunc("/live/measure", middleware.PostLiveMeasure).Methods("POST", "OPTIONS")

	router.HandleFunc("/live/notify", middleware.GetLiveNotify).Methods("GET", "OPTIONS")
	router.HandleFunc("/live/notify", middleware.PostLiveNotify).Methods("POST", "OPTIONS")

	router.HandleFunc("/live/control", middleware.GetLiveControl).Methods("GET", "OPTIONS")
	router.HandleFunc("/live/control", middleware.PostLiveControl).Methods("POST", "OPTIONS")

	http.Handle("/", router)
	return router
}
