package router

import (
	"net/http"

	mid "github.com/SPSOAFM-IT18/dmp-plant-hub/rest/middleware"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/init/measured", mid.HandleGetInitMeasured).Methods("GET", "OPTIONS")
	router.HandleFunc("/init/measured", mid.HandlePostInitMeasured).Methods("POST", "OPTIONS")

	router.HandleFunc("/live/measure", mid.HandleGetLiveMeasure).Methods("GET", "OPTIONS")
	router.HandleFunc("/live/measure", mid.HandlePostLiveMeasure).Methods("POST", "OPTIONS")

	router.HandleFunc("/live/notify", mid.HandleGetLiveNotify).Methods("GET", "OPTIONS")
	router.HandleFunc("/live/notify", mid.HandlePostLiveNotify).Methods("POST", "OPTIONS")

	router.HandleFunc("/live/control", mid.HandleGetLiveControl).Methods("GET", "OPTIONS")
	router.HandleFunc("/live/control", mid.HandlePostLiveControl).Methods("POST", "OPTIONS")

	http.Handle("/", router)

	return router
}
