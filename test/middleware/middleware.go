package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/SPSOAFM-IT18/dmp-plant-hub/test/env"
	"github.com/SPSOAFM-IT18/dmp-plant-hub/test/model"
)

func GetInitMeasured(w http.ResponseWriter, r *http.Request) {
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

func PostInitMeasured(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Accept", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", env.Process("CORS"))

	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.WriteHeader(http.StatusOK)

	var data model.InitMeasured
	_ = json.NewDecoder(r.Body).Decode(&data)
	fmt.Print("POST INIT MEASURED from Web app: ", data)
}

func GetLiveMeasure(w http.ResponseWriter, r *http.Request) {
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

func PostLiveMeasure(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Accept", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", env.Process("CORS"))

	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.WriteHeader(http.StatusOK)

	var data model.LiveMeasure
	_ = json.NewDecoder(r.Body).Decode(&data)
	fmt.Print("POST MEASURE from Web app: ", data)
}

func GetLiveNotify(w http.ResponseWriter, r *http.Request) {
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

func PostLiveNotify(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Accept", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", env.Process("CORS"))

	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.WriteHeader(http.StatusOK)

	var data model.LiveNotify
	_ = json.NewDecoder(r.Body).Decode(&data)
	fmt.Print("POST NOTIFY from Web app: ", data)
}

func GetLiveControl(w http.ResponseWriter, r *http.Request) {
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

func PostLiveControl(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Accept", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", env.Process("CORS"))
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.WriteHeader(http.StatusOK)

	var data model.LiveControl
	_ = json.NewDecoder(r.Body).Decode(&data)
	fmt.Print("POST CONTROL from Web app: ", data)
	//requests.PostLiveControl(data)
}
