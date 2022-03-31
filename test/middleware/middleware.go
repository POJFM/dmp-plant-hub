package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/SPSOAFM-IT18/dmp-plant-hub/test/env"
	"github.com/SPSOAFM-IT18/dmp-plant-hub/test/model"
)

var (
	pumpState bool = false
	moist          = 0.0
	hum            = 0.0
	temp           = 0.0
	LNtitle   string
	LNstate   = "inactive"
	LNaction  string
)

func LoadLiveMeasure(cMoist, cHum, cTemp float64) {
	moist = cMoist
	hum = cHum
	temp = cTemp
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
	data := model.LiveNotify{Title: LNtitle, State: LNstate, Action: LNaction}

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

	pumpState = data.PumpState

	if data.PumpState {
		LNtitle = "Zavlažování"
		LNstate = "inProgress"
		LNaction = "Probíhá zavlažování"
	} else {
		LNtitle = ""
		LNstate = "inactive"
		LNaction = ""
	}
}

func GetLiveControl(cPumpState chan bool) {
	cPumpState <- pumpState
}
