package middleware

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	db "github.com/SPSOAFM-IT18/dmp-plant-hub/test/database"
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
	Idb       *db.DB
	lat       float64
	lon       float64
)

func LoadLiveMeasure(cMoist, cHum, cTemp float64) {
	moist = cMoist
	hum = cHum
	temp = cTemp
}

func LoadInstances(db *db.DB) {
	Idb = db
}

func setPostHeader(w http.ResponseWriter) http.ResponseWriter {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Accept", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", env.Process("CORS"))
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.WriteHeader(http.StatusOK)
	return w
}

func setGetHeader(w http.ResponseWriter) http.ResponseWriter {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Accept", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", env.Process("CORS"))
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	return w
}

func HandleGetInitMeasured(w http.ResponseWriter, _ *http.Request) {
	data := model.GetInitMeasured{MoistLimit: 53.5, WaterLevelLimit: 50}

	w = setGetHeader(w)

	res, err := json.Marshal(data)
	if err != nil {
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(res)
}

func HandlePostInitMeasured(w http.ResponseWriter, r *http.Request) {
	w = setPostHeader(w)
	w.WriteHeader(http.StatusOK)

	var data model.PostInitMeasured
	_ = json.NewDecoder(r.Body).Decode(&data)

	lat = data.Lat
	lon = data.Lon
}

func HandleGetLiveMeasure(w http.ResponseWriter, _ *http.Request) {
	data := model.LiveMeasure{Moist: moist, Hum: hum, Temp: temp}

	w = setGetHeader(w)
	w.WriteHeader(http.StatusOK)

	res, err := json.Marshal(data)
	if err != nil {
		return
	}

	w.Write(res)
}

func HandlePostLiveMeasure(w http.ResponseWriter, r *http.Request) {
	w = setPostHeader(w)
	w.WriteHeader(http.StatusOK)

	var data model.LiveMeasure
	_ = json.NewDecoder(r.Body).Decode(&data)
	fmt.Print("POST MEASURE: ", data)
}

func HandleGetLiveNotify(w http.ResponseWriter, _ *http.Request) {
	data := model.LiveNotify{Title: LNtitle, State: LNstate, Action: LNaction}

	w = setGetHeader(w)
	w.WriteHeader(http.StatusOK)

	res, err := json.Marshal(data)
	if err != nil {
		return
	}

	w.Write(res)
}

func HandlePostLiveNotify(w http.ResponseWriter, r *http.Request) {
	w = setPostHeader(w)
	w.WriteHeader(http.StatusOK)

	var data model.LiveNotify
	_ = json.NewDecoder(r.Body).Decode(&data)
	fmt.Print("POST NOTIFY: ", data)
}

func HandleGetLiveControl(w http.ResponseWriter, _ *http.Request) {
	// actually default values, just haven't figured out how to pass them
	data := model.LiveControl{Restart: false, PumpState: false}

	w = setGetHeader(w)
	w.WriteHeader(http.StatusOK)

	res, err := json.Marshal(data)
	if err != nil {
		return
	}

	w.Write(res)
}

func HandlePostLiveControl(w http.ResponseWriter, r *http.Request) {
	w = setPostHeader(w)
	w.WriteHeader(http.StatusOK)

	var data model.LiveControl
	_ = json.NewDecoder(r.Body).Decode(&data)

	fmt.Println("POST LIVE CONTROL: ")
	fmt.Println(data.Restart)
	fmt.Println(data.PumpState)

	pumpState = data.PumpState

	if data.PumpState {
		LNtitle = "irrigation"
		LNstate = "inProgress"
		LNaction = "irrigationInProgress"
	} else {
		LNtitle = ""
		LNstate = "inactive"
		LNaction = ""
	}
}

func GetLiveControl(cPumpState chan bool) {
	cPumpState <- pumpState
}

func HandleGetWeather(w http.ResponseWriter, _ *http.Request) {
	w = setGetHeader(w)
	w.WriteHeader(http.StatusOK)
	//w = setGetHeader(w)
	if Idb.CheckSettings() {
		geocodes := Idb.GetSettingByColumn([]string{"lat", "lon"})
		lat = *geocodes.Lat
		lon = *geocodes.Lon
	}

	res, err := http.Get("https://api.openweathermap.org/data/2.5/onecall?lat=" + fmt.Sprintf("%f", lat) + "&lon=" + fmt.Sprintf("%f", lon) + "&exclude=daily,minutely,alerts&units=metric&appid=" + env.Process("WEATHER_API_KEY")) //nolint:bodyclose

	if err != nil {
		w.WriteHeader(res.StatusCode)
	}
	defer res.Body.Close()
	log.Println(res.Body)

	w.WriteHeader(http.StatusOK)
	_, _ = io.Copy(w, res.Body)
}

func HandleGetGeocode(w http.ResponseWriter, _ *http.Request) {
	w = setGetHeader(w)
	w.WriteHeader(http.StatusOK)

	if Idb.CheckSettings() {
		geocodes := Idb.GetSettingByColumn([]string{"lat", "lon"})
		lat = *geocodes.Lat
		lon = *geocodes.Lon
	}

	res, err := http.Get("https://api.opencagedata.com/geocode/v1/json?q=" + fmt.Sprintf("%f", lat) + "+" + fmt.Sprintf("%f", lon) + "&key=" + env.Process("GEOCODE_API_KEY")) //nolint:bodyclose

	if err != nil {
		w.WriteHeader(res.StatusCode)
	}
	defer res.Body.Close()

	w.WriteHeader(http.StatusOK)
	_, _ = io.Copy(w, res.Body)
}
func HandlePostGeocode(w http.ResponseWriter, r *http.Request) {

}

func HandleGetGoogle(w http.ResponseWriter, _ *http.Request) {

}

func HandlePostGoogle(w http.ResponseWriter, r *http.Request) {

}
