package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/SPSOAFM-IT18/dmp-plant-hub/env"
	"github.com/SPSOAFM-IT18/dmp-plant-hub/rest/model"
	sens "github.com/SPSOAFM-IT18/dmp-plant-hub/sensors"
)

var (
	pumpState = false
	moist     float64
	hum       float64
	temp      float64
	WLL       float64
	LNtitle   string
	LNstate   = "inactive"
	LNaction  string
	sensPump  *sens.Sensors
)

func LoadInitMeasured(initM, initWLL *float64) {
	moist = *initM
	WLL = *initWLL
}

func LoadLiveMeasure(cMoist, cHum, cTemp chan float64) {
	moist = <-cMoist
	hum = <-cHum
	temp = <-cTemp
}

func LoadLiveNotify(title, state, action string) {
	LNtitle = title
	LNstate = state
	LNaction = action
}

func GetLiveControl(cPumpState chan bool) {
	cPumpState <- pumpState
}

func LoadPumpState(sensei *sens.Sensors) {
	sensPump = sensei
}

func HandleGetInitMeasured(w http.ResponseWriter, _ *http.Request) {
	data := model.InitMeasured{MoistLimit: moist, WaterLevelLimit: WLL}

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
}

func HandleGetLiveControl(w http.ResponseWriter, _ *http.Request) {
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
		os.Exit(0)
	}

	if data.PumpState {
		sensPump.StartPump()
	}

	if !data.PumpState {
		sensPump.StopPump()
	}

	pumpState = data.PumpState
}
