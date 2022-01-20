package sequences

import (
	"fmt"
	"time"

	"github.com/SPSOAFM-IT18/dmp-plant-hub/rest/model"
	"github.com/SPSOAFM-IT18/dmp-plant-hub/rest/requests"
	"github.com/SPSOAFM-IT18/dmp-plant-hub/utils"
	"github.com/stianeikeland/go-rpio"
)

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

// END TEST

func CheckingSequence(PUMP rpio.Pin, LED rpio.Pin) {
	// get from DB
	// values only for test

	const waterLevelLimit = 75

	if waterLevelMeasure() < waterLevelLimit {
		requests.PostLiveNotify(model.LiveNotify{Title: "Doplňte nádrž", State: "physicalHelpRequired", Action: "Nádrž je prázdná"})

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
		requests.PostLiveNotify(model.LiveNotify{Title: "Kontrola Nádrže", State: "finished", Action: "V nádrži zbýva %vl vody"}, waterLevelMeasure())
	}
}

func IrrigationSequence(PUMP rpio.Pin, LED rpio.Pin, state bool) {
	fmt.Println("Starting irrigation...🌿🤖🚿")

	requests.PostLiveNotify(model.LiveNotify{Title: "Zavlažování", State: "inProgress", Action: "Probíhá zavlažování"})

	// get from DB
	// values only for test
	const moistureLimit = 50
	const waterAmountLimit = 50
	// Definovaný průtok čerpadla
	var pumpFlow float32 = 1.75 // litr/min

	for state {
		if moistureMeasure() < moistureLimit {

			// time passed from running pump will be represented as liters
			var flowMeasure float32
			t0 := time.Now()
			for waterLevelMeasure() < moistureLimit || flowMeasure < utils.TimeToOverdraw(waterAmountLimit, pumpFlow) {
				//var t1 float32 = time.time()
				PUMP.High()
				flowMeasure = float32(time.Since(t0).Seconds())
			}

			requests.PostLiveNotify(model.LiveNotify{Title: "Zavlažování", State: "finished", Action: "Zavlažování dokončeno"})

			time.Sleep(3000 * time.Millisecond)

			requests.PostLiveNotify(model.LiveNotify{Title: "Kontrola Nádrže", State: "inProgress", Action: "Probíhá kontrola nádrže"})

			// after pump stops run Checking sequence
			//CheckingSequence(PUMP rpio.Pin, LED rpio.Pin)
		} else {
			PUMP.Low()

			requests.PostLiveNotify(model.LiveNotify{Title: "", State: "inactive", Action: ""})
		}
	}
}

func InitializationSequence() {
	fmt.Println("Starting initialization sequence...🏁🤖🏁")
	time.Sleep(2000 * time.Millisecond)

	// var waterLevel float32
	// var moistureLevel float32
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

	moistureLevel := utils.ArithmeticMean(moistureAvg)
	waterLevel := utils.ArithmeticMean(waterLevelAvg)

	requests.PostInitMeasured(model.InitMeasured{MoistLimit: moistureLevel, WaterLevelLimit: waterLevel})
}

func MeasurementSequence(PUMP rpio.Pin, LED rpio.Pin) {
	for true {
		moisture := moistureMeasure()
		// dodělat aby DHT measure vracel temp a hum v array nebo objectu
		//var DHTMeasureValues = DHTMeasure()
		// potom tyhle variables odjebat a mrdnout tam přímo ty měřící funkce
		temperature := 20 // DHTMeasureValues[0]
		humidity := 50    // DHTMeasureValues[1]

		requests.PostLiveMeasure(model.LiveMeasure{Moist: moisture, Hum: humidity, Temp: temperature})

		fmt.Println("\nTemperature: %v˚C", temperature)
		fmt.Println("\nHumidity: %v%", humidity)
		fmt.Println("\nSoil moisture: %v%", moisture)
		time.Sleep(1000 * time.Millisecond)
	}
}
