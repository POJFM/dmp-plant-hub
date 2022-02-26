package sequences

import (
	"fmt"
	"time"

	mid "github.com/SPSOAFM-IT18/dmp-plant-hub/rest/middleware"
	"github.com/SPSOAFM-IT18/dmp-plant-hub/rest/model"
	req "github.com/SPSOAFM-IT18/dmp-plant-hub/rest/requests"
	sens "github.com/SPSOAFM-IT18/dmp-plant-hub/sensors"
	"github.com/SPSOAFM-IT18/dmp-plant-hub/utils"

	"github.com/jasonlvhit/gocron"
)

// TEST
func waterLevelMeasure() float64 {
	return 1
}
func moistureMeasure() float64 {
	return 1
}
func DHTMeasure() float64 {
	return 1
}

// END TEST

// on start waits for time to be hour o'clock
// then starts chron routine that is timed on every 4 hours
func SaveOnFourHoursPeriod(cMoist, cTemp, cHum chan float64) {
	for time.Now().Format("04") != "00" {
		// TEST
		fmt.Println(time.Now().Format("04"))
		// END TEST
		time.Sleep(1 * time.Minute)
	}
	gocron.Every(4).Hours().Do(func() {
		moist := <-cMoist
		temp := <-cTemp
		hum := <-cHum
		// TEST
		fmt.Println("Cron: ", moist, temp, hum)
		// END TEST
	})
	<-gocron.Start()
}

func CheckingSequence(cMoist chan float64) {
	// get from DB
	// values only for test

	const waterLevelLimit = 75

	if <-cMoist < waterLevelLimit {
		req.PostLiveNotify(model.LiveNotify{Title: "Doplňte nádrž", State: "physicalHelpRequired", Action: "Nádrž je prázdná"})

		fmt.Println("Water tank limit level reached...🚫🤖🚫")

		for <-cMoist < waterLevelLimit {
			sens.LED.High()
			time.Sleep(1000 * time.Millisecond)
			sens.LED.Low()
			time.Sleep(1000 * time.Millisecond)
		}
	} else {
		sens.LED.Low()

		waterLevel := fmt.Sprintf("V nádrži zbývá %fl vody", waterLevelMeasure())
		// Dodělat na water amount v litrech
		req.PostLiveNotify(model.LiveNotify{Title: "Kontrola Nádrže", State: "finished", Action: waterLevel})
	}
}

func IrrigationSequence(cMoist chan float64, cRestart, cPumpState chan bool) {
	fmt.Println("Starting irrigation...🌿🤖🚿")

	req.PostLiveNotify(model.LiveNotify{Title: "Zavlažování", State: "inProgress", Action: "Probíhá zavlažování"})

	// get from DB
	// values only for test
	moistureLimit := 50.0
	waterAmountLimit := 50.0
	// Definovaný průtok čerpadla
	var pumpFlow float64 = 1.75 // litr/min

	gocron.Every(1).Seconds().Do(func() {
		if <-cRestart {
			// get from DB
			// values only for test
			moistureLimit = 50
			waterAmountLimit = 50

			req.PostLiveControl(model.LiveControl{Restart: false, PumpState: false})
		}

		if <-cMoist < moistureLimit {
			// time passed from running pump will be represented as liters
			var flowMeasure float64
			t0 := time.Now()
			// TimeToOverdraw is calculated by divideing amount by flow
			for waterLevelMeasure() < moistureLimit || flowMeasure < waterAmountLimit/pumpFlow {
				//var t1 float64 = time.time()
				sens.PUMP.High()
				flowMeasure = float64(time.Since(t0).Seconds())
			}

			req.PostLiveNotify(model.LiveNotify{Title: "Zavlažování", State: "finished", Action: "Zavlažování dokončeno"})

			time.Sleep(3000 * time.Millisecond)

			req.PostLiveNotify(model.LiveNotify{Title: "Kontrola Nádrže", State: "inProgress", Action: "Probíhá kontrola nádrže"})

			// after pump stops run Checking sequence
			//CheckingSequence(cMoist)
		} else {
			sens.PUMP.Low()

			req.PostLiveNotify(model.LiveNotify{Title: "", State: "inactive", Action: ""})
		}
	})
	<-gocron.Start()
}

func InitializationSequence(cMoist chan float64) {
	fmt.Println("Starting initialization sequence...🏁🤖🏁")
	time.Sleep(2000 * time.Millisecond)

	// var waterLevel float64
	// var moistureLevel float64
	var waterLevelAvg []float64
	waterLevelAvg = make([]float64, 5)
	var moistureAvg []float64
	moistureAvg = make([]float64, 5)

	// calculating average value
	for i := 0; i < 5; i++ {
		moistureAvg = append(moistureAvg, <-cMoist)
		waterLevelAvg = append(waterLevelAvg, waterLevelMeasure())
		time.Sleep(1000 * time.Millisecond)
	}

	moistureLevel := utils.ArithmeticMean(moistureAvg)
	waterLevel := utils.ArithmeticMean(waterLevelAvg)

	req.PostInitMeasured(model.InitMeasured{MoistLimit: moistureLevel, WaterLevelLimit: waterLevel})
}

func MeasurementSequence(cMoist, cTemp, cHum chan float64, cRestart, cPumpState chan bool) {
	gocron.Every(1).Seconds().Do(func() {
		moisture := float64(moistureMeasure())
		// dodělat aby DHT measure vracel temp a hum v array nebo objectu
		//var DHTMeasureValues = DHTMeasure()
		// potom tyhle variables odjebat a mrdnout tam přímo ty měřící funkce
		temperature := float64(20) // DHTMeasureValues[0]
		humidity := float64(5)     // DHTMeasureValues[1]

		req.PostLiveMeasure(model.LiveMeasure{Moist: moisture, Hum: humidity, Temp: temperature})

		fmt.Println("\nTemperature: %v˚C", temperature)
		fmt.Println("\nHumidity: %v%", humidity)
		fmt.Println("\nSoil moisture: %v%", moisture)

		mid.GetLiveControl(cRestart, cPumpState)

		cMoist <- moisture
		cTemp <- temperature
		cHum <- humidity
	},
	)
	<-gocron.Start()
}
