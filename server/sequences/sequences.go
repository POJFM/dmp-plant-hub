package sequences

import (
	"fmt"
	"time"

	"github.com/SPSOAFM-IT18/dmp-plant-hub/rest/model"
	"github.com/SPSOAFM-IT18/dmp-plant-hub/rest/requests"
	"github.com/SPSOAFM-IT18/dmp-plant-hub/utils"
	"github.com/jasonlvhit/gocron"
	"github.com/stianeikeland/go-rpio/v4"
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

// on start waits for time to be hour o'clock
// then starts chron routine that is timed on every 4 hours
func SaveOnFourHoursPeriod(moisture, temperature, humidity float32) {
	for time.Now().Format("04") != "00" {
		// TEST
		fmt.Println(time.Now().Format("04"))
		// END TEST
		time.Sleep(1 * time.Minute)
	}
	// needs to access values from measurements sequence
	gocron.Every(4).Hours().Do( /* save data to db */ )
	<-gocron.Start()
}

func CheckingSequence(PUMP, LED rpio.Pin, moisture float32) {
	// get from DB
	// values only for test

	const waterLevelLimit = 75

	if moisture < waterLevelLimit {
		requests.PostLiveNotify(model.LiveNotify{Title: "Doplňte nádrž", State: "physicalHelpRequired", Action: "Nádrž je prázdná"})

		fmt.Println("Water tank limit level reached...🚫🤖🚫")

		for moisture < waterLevelLimit {
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

func IrrigationSequence(PUMP, LED rpio.Pin, state bool, moisture float32) {
	fmt.Println("Starting irrigation...🌿🤖🚿")

	requests.PostLiveNotify(model.LiveNotify{Title: "Zavlažování", State: "inProgress", Action: "Probíhá zavlažování"})

	// get from DB
	// values only for test
	const moistureLimit = 50
	const waterAmountLimit = 50
	// Definovaný průtok čerpadla
	var pumpFlow float32 = 1.75 // litr/min

	for state {
		if moisture < moistureLimit {

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

func InitializationSequence(moisture float32) {
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
		moistureAvg = append(moistureAvg, moisture)
		waterLevelAvg = append(waterLevelAvg, waterLevelMeasure())
		count++
		time.Sleep(1000 * time.Millisecond)
	}

	moistureLevel := utils.ArithmeticMean(moistureAvg)
	waterLevel := utils.ArithmeticMean(waterLevelAvg)

	requests.PostInitMeasured(model.InitMeasured{MoistLimit: moistureLevel, WaterLevelLimit: waterLevel})
}

func MeasurementSequence(PUMP, LED rpio.Pin, cMoisture, cTemperature, cHumidity chan float32) {
	for true {
		moisture := moistureMeasure()
		// dodělat aby DHT measure vracel temp a hum v array nebo objectu
		//var DHTMeasureValues = DHTMeasure()
		// potom tyhle variables odjebat a mrdnout tam přímo ty měřící funkce
		temperature := float32(20) // DHTMeasureValues[0]
		humidity := float32(5)     // DHTMeasureValues[1]

		requests.PostLiveMeasure(model.LiveMeasure{Moist: moisture, Hum: humidity, Temp: temperature})

		fmt.Println("\nTemperature: %v˚C", temperature)
		fmt.Println("\nHumidity: %v%", humidity)
		fmt.Println("\nSoil moisture: %v%", moisture)

		cMoisture <- moisture
		cTemperature <- temperature
		cHumidity <- humidity

		time.Sleep(1000 * time.Millisecond)
	}
}
