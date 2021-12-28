package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// IDK how to make it work
// needs to reload when data changes

type LiveMeasure struct {
	Moist float32 `json:"moist"`
	Hum   float32 `json:"hum"`
	Temp  float32 `json:"temp"`
}

type InitMeasured struct {
	MoistLimit      float32 `json:"moistLimit"`
	WaterLevelLimit float32 `json:"waterLevelLimit"`
}

var initMeasuredData = InitMeasured{MoistLimit: 53.5, WaterLevelLimit: 50}

type LiveMeasureArray []LiveMeasure

var liveData = LiveMeasureArray{
	LiveMeasure{Moist: 50.5, Hum: 45, Temp: 20},
	LiveMeasure{Moist: 40, Hum: 44.4, Temp: 14},
}

func CreateLiveMeasure(w http.ResponseWriter, r *http.Request) {
	//liveMeasureArray = append(liveMeasureArray, LiveMeasure{Moist: 50.5, Hum: 45, Temp: 20})
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(liveData)
	fmt.Println("live measured data")
	liveData = append(liveData, LiveMeasure{Moist: 40, Hum: 44.4, Temp: 14})
	// for true {
	// 	time.Sleep(3000 * time.Millisecond)
	// 	liveData = append(liveData, LiveMeasure{Moist: 40, Hum: 44.4, Temp: 14})
	// 	fmt.Println("live measured data")

	// }

}

//random koment

// func handleRequests() {
// 	http.HandleFunc("/live/measure", CreateLiveMeasure)
// 	log.Fatal(http.ListenAndServe(":5000", nil))
// }

func main() {
	//handleRequests()
	// liveData = append(liveData, LiveMeasure{Moist: 50.5, Hum: 45, Temp: 20})
	// liveData = append(liveData, LiveMeasure{Moist: 40, Hum: 44.4, Temp: 14})
	// http.HandleFunc("/live/measure", CreateLiveMeasure)
	// time.Sleep(3000 * time.Millisecond)
	// fmt.Println("kokot")
	// log.Fatal(http.ListenAndServe(":5000", nil))

	// time.Sleep(3000 * time.Millisecond)
	// liveData = append(liveData, LiveMeasure{Moist: 40, Hum: 44.4, Temp: 14})
	// liveData = append(liveData, LiveMeasure{Moist: 20, Hum: 43.4, Temp: 10})
	// http.HandleFunc("/live/measure", CreateLiveMeasure)
	// log.Fatal(http.ListenAndServe(":5000", nil))
	// for true {
	// 	time.Sleep(3000 * time.Millisecond)
	// 	fmt.Println("kokot")
	// }
	http.HandleFunc("/live/measure", CreateLiveMeasure)
	log.Fatal(http.ListenAndServe(":5000", nil))
}
