package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/SPSOAFM-IT18/dmp-plant-hub/env"
	"github.com/SPSOAFM-IT18/dmp-plant-hub/rest/model"
)

var pumpState = false
var moist = 0.0
var hum = 0.0
var temp = 0.0

func LoadLiveMeasure(cMoist, cHum, cTemp chan float64) {
	moist = <-cMoist
	hum = <-cHum
	temp = <-cTemp
}

func HandleGetInitMeasured(w http.ResponseWriter, _ *http.Request) {
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

	fmt.Print("GET INIT MEASURED: ", res)
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

func HandleGetLiveMeasure(w http.ResponseWriter, _ *http.Request) {
	data := model.LiveMeasure{Moist: moist, Hum: hum, Temp: temp}

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

	fmt.Print("GET MEASURE: ", res)
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

func HandleGetLiveNotify(w http.ResponseWriter, _ *http.Request) {
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

	fmt.Print("GET LIVE NOTIFY: ", res)
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

func HandleGetLiveControl(w http.ResponseWriter, _ *http.Request) {
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

	fmt.Print("GET LIVE CONTROL: ", res)
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
		os.Exit(0)
	}

	pumpState = data.PumpState
}

func GetLiveControl(cPumpState chan bool) {
	cPumpState <- pumpState
}
