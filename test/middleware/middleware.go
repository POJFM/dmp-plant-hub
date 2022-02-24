package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/SPSOAFM-IT18/dmp-plant-hub/test/env"
	"github.com/SPSOAFM-IT18/dmp-plant-hub/test/model"
)

var restart bool = false
var pumpState bool = false

func HandleGetInitMeasured(w http.ResponseWriter, r *http.Request) {
	// TEST
	data := model.InitMeasured{MoistLimit: 53.5, WaterLevelLimit: 50}

	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Accept", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", env.Process("CORS"))

	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.WriteHeader(http.StatusOK)

	res, err := json.Marshal(data)
	if err != nil {
		return
	}

	w.Write(res)
}

func HandlePostInitMeasured(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Accept", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", env.Process("CORS"))

	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.WriteHeader(http.StatusOK)

	var data model.InitMeasured
	_ = json.NewDecoder(r.Body).Decode(&data)
	fmt.Print("POST INIT MEASURED: ", data)
}

func HandleGetLiveMeasure(w http.ResponseWriter, r *http.Request) {
	// TEST
	data := model.LiveMeasure{Moist: 50.5, Hum: 45, Temp: 20}

	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Accept", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", env.Process("CORS"))

	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.WriteHeader(http.StatusOK)

	res, err := json.Marshal(data)
	if err != nil {
		return
	}

	w.Write(res)
}

func HandlePostLiveMeasure(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Accept", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", env.Process("CORS"))

	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.WriteHeader(http.StatusOK)

	var data model.LiveMeasure
	_ = json.NewDecoder(r.Body).Decode(&data)
	fmt.Print("POST MEASURE: ", data)
}

func HandleGetLiveNotify(w http.ResponseWriter, r *http.Request) {
	// actually default values, just haven't figured out how to pass them
	data := model.LiveNotify{Title: "", State: "inactive", Action: ""}

	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Accept", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", env.Process("CORS"))

	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.WriteHeader(http.StatusOK)

	res, err := json.Marshal(data)
	if err != nil {
		return
	}

	w.Write(res)
}

func HandlePostLiveNotify(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Accept", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", env.Process("CORS"))

	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.WriteHeader(http.StatusOK)

	var data model.LiveNotify
	_ = json.NewDecoder(r.Body).Decode(&data)
	fmt.Print("POST NOTIFY: ", data)
}

func HandleGetLiveControl(w http.ResponseWriter, r *http.Request) {
	// actually default values, just haven't figured out how to pass them
	data := model.LiveControl{Restart: false, PumpState: false}

	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Accept", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", env.Process("CORS"))

	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.WriteHeader(http.StatusOK)

	res, err := json.Marshal(data)
	if err != nil {
		return
	}

	w.Write(res)
}

func HandlePostLiveControl(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Accept", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", env.Process("CORS"))
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.WriteHeader(http.StatusOK)

	var data model.LiveControl
	_ = json.NewDecoder(r.Body).Decode(&data)

	fmt.Println("POST LIVE CONTROL: ")
	fmt.Println(data.Restart)
	fmt.Println(data.PumpState)

	if data.Restart {
		restart = true
	}

	if data.PumpState {
		pumpState = true
	}
}

func GetLiveControl(cRestart, cPumpState chan bool) {
	cRestart <- restart
	cPumpState <- pumpState
}
