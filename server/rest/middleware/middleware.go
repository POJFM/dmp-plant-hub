package middleware

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	db "github.com/SPSOAFM-IT18/dmp-plant-hub/database"
	"github.com/SPSOAFM-IT18/dmp-plant-hub/env"
	"github.com/SPSOAFM-IT18/dmp-plant-hub/rest/model"
	sens "github.com/SPSOAFM-IT18/dmp-plant-hub/sensors"
	"github.com/SPSOAFM-IT18/dmp-plant-hub/utils"
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
	Isens     *sens.Sensors
	Idb       *db.DB
	lat       float64
	lon       float64
)

func LoadInitMeasured(initM, initWLL *float64) {
	moist = *initM
	WLL = *initWLL
}

func LoadLiveMeasure(gMoist, gHum, gTemp *float64) {
	moist = *gMoist
	hum = *gHum
	temp = *gTemp
}

func LoadLiveNotify(title, state, action string) {
	LNtitle = title
	LNstate = state
	LNaction = action
}

func GetLiveControl() bool {
	return pumpState
}

func LoadInstances(db *db.DB, sensei *sens.Sensors) {
	Idb = db
	Isens = sensei
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
	data := model.GetInitMeasured{MoistLimit: moist, WaterLevelLimit: WLL}
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
	var data model.PostInitMeasured
	_ = json.NewDecoder(r.Body).Decode(&data)
	fmt.Print("POST INIT MEASURED: ", data)

	lat = data.Lat
	lon = data.Lon
}

func HandleGetLiveMeasure(w http.ResponseWriter, _ *http.Request) {
	data := model.LiveMeasure{Moist: moist, Hum: hum, Temp: temp}
	w = setGetHeader(w)
	res, err := json.Marshal(data)
	if err != nil {
		return
	}
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(res)
}

func HandlePostLiveMeasure(w http.ResponseWriter, r *http.Request) {
	w = setPostHeader(w)
	var data model.LiveMeasure
	_ = json.NewDecoder(r.Body).Decode(&data)
}

func HandleGetLiveNotify(w http.ResponseWriter, _ *http.Request) {
	data := model.LiveNotify{Title: LNtitle, State: LNstate, Action: LNaction}
	w = setGetHeader(w)
	res, err := json.Marshal(data)
	if err != nil {
		return
	}
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(res)
}

func HandlePostLiveNotify(w http.ResponseWriter, r *http.Request) {
	w = setPostHeader(w)
	var data model.LiveNotify
	_ = json.NewDecoder(r.Body).Decode(&data)
}

func HandleGetLiveControl(w http.ResponseWriter, _ *http.Request) {
	data := model.LiveControl{Restart: false, PumpState: false}
	w = setGetHeader(w)
	res, err := json.Marshal(data)
	if err != nil {
		return
	}
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(res)
}

func HandlePostLiveControl(w http.ResponseWriter, r *http.Request) {
	w = setPostHeader(w)
	var data model.LiveControl
	_ = json.NewDecoder(r.Body).Decode(&data)

	if data.Restart {
		utils.Exit()
	}
	pumpState = data.PumpState
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
	fmt.Println(res.Body)

	w.WriteHeader(http.StatusOK)
	_, _ = io.Copy(w, res.Body)
}

func HandleGetGeocode(w http.ResponseWriter, _ *http.Request) {
	w = setGetHeader(w)
	w.WriteHeader(http.StatusOK)
	//w = setGetHeader(w)

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
	//fmt.Println(res.Body)

	w.WriteHeader(http.StatusOK)
	_, _ = io.Copy(w, res.Body)
}

func HandlePostGeocode(w http.ResponseWriter, r *http.Request) {

}

func HandleGetGoogle(w http.ResponseWriter, _ *http.Request) {

}

func HandlePostGoogle(w http.ResponseWriter, r *http.Request) {

}
