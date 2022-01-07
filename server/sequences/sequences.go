package sequences

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"time"

	"github.com/SPSOAFM-IT18/dmp-plant-hub/rest/model"
	"github.com/SPSOAFM-IT18/dmp-plant-hub/utils"
	"github.com/stianeikeland/go-rpio"
)

// TRY THIS FIRST
// requests.PostLiveNotify(model.LiveNotify{Title: "Zavlažování", State: "inProgress", Action: "Probíhá zavlažování"})
// requests.PostLiveControl(model.LiveControl{Restart: false, PumpState: false})
// 	fmt.Println("LIVE CONTROL DATA: %v", requests.GetLiveControl())

// TEST
func waterLevelMeasure() float32 {
	return 1
}
func moistureMeasure() float32 {
	return 1
}
func DHTMeasure() float32 {
	return 1
}

var initMeasuredData = model.InitMeasured{MoistLimit: 53.5, WaterLevelLimit: 50}

var liveMeasureData = model.LiveMeasure{Moist: 50.5, Hum: 45, Temp: 20}

var liveNotifyData = model.LiveNotify{Title: "", State: "inactive", Action: ""}

// END TEST

// maybe a good thing to put these into separate file but I haven't figured out a way to pass data values in it
func buildInitMeasured(w http.ResponseWriter, r *http.Request) {
	fmt.Println("init measured")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(initMeasuredData)
}

func buildLiveMeasure(w http.ResponseWriter, r *http.Request) {
	fmt.Println("live measured data")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(liveMeasureData)
}

func buildLiveNotify(w http.ResponseWriter, r *http.Request) {
	fmt.Println("live notify")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(liveNotifyData)
}

// func buildInitMeasuredData(moistLimit float32, waterLevelLimit float32)  {
// 	initMeasuredData := InitMeasured{moistLimit, waterLevelLimit}
// 	initMeasuredDataEnc := json.Data(initMeasuredData)
// 	return initMeasuredDataEnc
// }

// func buildLiveMeasureData (moist float32, hum float32, temp float32)  {
// 	liveMeasureData := liveMeasure{moist, hum, temp}
//   liveMeasureDataEnc := json.Data(liveMeasureData)
// 	return liveMeasureDataEnc
// }

// func buildLiveNotifyData (title string, state string, action string)  {
// 	liveNotifyData := LiveNotify{title, state, action}
//   liveNotifyDataEnc := json.Data(liveNotifyData)
// 	return liveNotifyDataEnc
//}

func InitializationSequence() {
	fmt.Println("Starting initialization sequence...🏁🤖🏁")
	time.Sleep(2000 * time.Millisecond)

	var waterLevel float32
	var moistureLevel float32
	var waterLevelAvg []float32
	waterLevelAvg = make([]float32, 5)
	var moistureAvg []float32
	moistureAvg = make([]float32, 5)

	// calculating average value
	var count int = 0
	for count < 5 {
		moistureAvg = append(moistureAvg, moistureMeasure())
		waterLevelAvg = append(waterLevelAvg, waterLevelMeasure())
		count += 1
		time.Sleep(1000 * time.Millisecond)
	}

	moistureLevel = utils.ArithmeticMean(moistureAvg)
	waterLevel = utils.ArithmeticMean(waterLevelAvg)

	initMeasuredData = model.InitMeasured{MoistLimit: moistureLevel, WaterLevelLimit: waterLevel}

	// send limit values to web api
	http.HandleFunc("/init/measured", buildInitMeasured)

	// saving logic in frontend
	log.Fatal(http.ListenAndServe(":5000", nil))
}

func MeasurementSequence(PUMP rpio.Pin, LED rpio.Pin) {
	// get from DB
	// values only for test
	const waterLevelLimit = 75
	const moistureLimit = 50
	const waterAmountLimit = 50
	// Definovaný průtok čerpadla
	var pumpFlow float32 = 1.75 // litr/min

	// nonstop reading from sensors
	var moisture float32
	var humidity float32
	var temperature float32
	time.Sleep(2000 * time.Millisecond)
	fmt.Printf("\nWater level limit: %vcm", math.Round(float64(waterLevelLimit)))
	fmt.Println("\nWater amount limit: %vl", math.Round(float64(moistureLimit)))
	fmt.Println("\nMoisture limit: %v%", math.Round(float64(waterAmountLimit)))
	time.Sleep(3000 * time.Millisecond)

	http.HandleFunc("/live/notify", buildLiveNotify)

	for true {
		moisture = moistureMeasure()
		// dodělat aby DHT measure vracel temp a hum v array nebo jsonu
		//var DHTMeasureValues = DHTMeasure()
		temperature = 20 // DHTMeasureValues[0]
		humidity = 50    // DHTMeasureValues[1]

		// live measure needs to be an array of json data
		// Idk how but it needs to be done or the whole world is doomed
		liveMeasureData = append(liveMeasureData, model.LiveMeasure{Moist: moisture, Hum: humidity, Temp: temperature})
		http.HandleFunc("/live/measure", buildLiveMeasure)

		if moistureMeasure() < moistureLimit {
			liveNotifyData = model.LiveNotify{Title: "Zavlažování", State: "inProgress", Action: "Probíhá zavlažování"}
			http.HandleFunc("/live/notify", buildLiveNotify)
			fmt.Println("Starting irrigation...🌿🤖🚿")

			// time passed from running pump will be represented as liters
			var flowMeasure float32
			t0 := time.Now()
			for waterLevelMeasure() < moistureLimit || flowMeasure < utils.TimeToOverdraw(waterAmountLimit, pumpFlow) {
				//var t1 float32 = time.time()
				PUMP.High()
				flowMeasure = float32(time.Since(t0).Seconds())
			}

			liveNotifyData = model.LiveNotify{Title: "Zavlažování", State: "finished", Action: "Zavlažování dokončeno"}
			http.HandleFunc("/live/notify", buildLiveNotify)

			time.Sleep(3000 * time.Millisecond)

			liveNotifyData = model.LiveNotify{Title: "Kontrola Nádrže", State: "inProgress", Action: "Probíhá kontrola nádrže"}
			http.HandleFunc("/live/notify", buildLiveNotify)

			// after pump stops run Checking sequence
			if waterLevelMeasure() < waterLevelLimit {
				liveNotifyData = model.LiveNotify{Title: "Doplňte nádrž", State: "physicalHelpRequired", Action: "Nádrž je prázdná"}
				http.HandleFunc("/live/notify", buildLiveNotify)

				fmt.Println("Water tank limit level reached...🚫🤖🚫")

				for waterLevelMeasure() < waterLevelLimit {
					LED.High()
					time.Sleep(1000 * time.Millisecond)
					LED.Low()
					time.Sleep(1000 * time.Millisecond)
				}
			} else {
				LED.Low()

				// Dodělat na water amount v litrech
				liveNotifyData = model.LiveNotify{Title: "Kontrola Nádrže", State: "finished", Action: "V nádrži zbýva 3l vody"}
				http.HandleFunc("/live/notify", buildLiveNotify)
			}
		} else {
			PUMP.Low()

			liveNotifyData = model.LiveNotify{Title: "", State: "inactive", Action: ""}
			http.HandleFunc("/live/notify", buildLiveNotify)
		}

		// dodělat aby DHT measure vracel temp a hum v array nebo jsonu
		fmt.Println("\nTemperature: %v˚C", 15)
		fmt.Println("\nHumidity: %v%", 45)
		fmt.Println("\nSoil moisture: %v%", moisture)
		time.Sleep(1000 * time.Millisecond)
	}
	log.Fatal(http.ListenAndServe(":5000", nil))
}
