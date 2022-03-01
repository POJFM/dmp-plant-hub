package sequences

import (
	"fmt"
	"math"
	"time"

	mid "github.com/SPSOAFM-IT18/dmp-plant-hub/rest/middleware"

	db "github.com/SPSOAFM-IT18/dmp-plant-hub/database"
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

func SaveOnFourHoursPeriod(cMoist, cTemp, cHum chan float64) {
	utils.WaitTillWholeHour()

	gocron.Every(4).Hours().Do(func() {
		moist := <-cMoist
		temp := <-cTemp
		hum := <-cHum

		// Save to DB

		// TEST
		fmt.Println("Cron: ", moist, temp, hum)
		// END TEST
	})
	<-gocron.Start()
}

func Controller(db *db.DB, sensei *sens.Sensors, cMoist chan float64, cPumpState chan bool) {
	if db.CheckSettings() {
		go irrigationSequence(cMoist, cPumpState)
	} else {
		go initializationSequence(sensei, cMoist)
		initializationFinished := true
		for initializationFinished {
			stopLED := make(chan bool)
			go func() {
				for {
					select {
					case <-stopLED:
						return
					default:
						for i := 0; i < 2; i++ {
							sens.LED.High()
							time.Sleep(500 * time.Millisecond)
							sens.LED.Low()
							time.Sleep(500 * time.Millisecond)
						}
						time.Sleep(1500 * time.Millisecond)
					}
				}
			}()
			if db.CheckSettings() {
				initializationFinished = false
				stopLED <- true
				go irrigationSequence(cMoist, cPumpState)
			}
			time.Sleep(1000 * time.Millisecond)
		}
	}
}

func checkingSequence() {
	// get from DB
	// values only for test
	const waterLevelLimit = 75

	req.PostLiveNotify(model.LiveNotify{Title: "Kontrola Nádrže", State: "inProgress", Action: "Probíhá kontrola nádrže"})

	time.Sleep(3000 * time.Millisecond)

	if waterLevelMeasure() < waterLevelLimit {
		req.PostLiveNotify(model.LiveNotify{Title: "Doplňte nádrž", State: "physicalHelpRequired", Action: "Nádrž je prázdná"})

		fmt.Println("Water tank limit level reached...🚫🤖🚫")

		for waterLevelMeasure() < waterLevelLimit {
			sens.LED.High()
			time.Sleep(1000 * time.Millisecond)
			sens.LED.Low()
			time.Sleep(1000 * time.Millisecond)
		}
	}

	waterLevel := fmt.Sprintf("V nádrži zbývá %fl vody", waterLevelMeasure())
	// Dodělat na water amount v litrech
	req.PostLiveNotify(model.LiveNotify{Title: "Kontrola Nádrže", State: "finished", Action: waterLevel})

	time.Sleep(3000 * time.Millisecond)

	req.PostLiveNotify(model.LiveNotify{Title: "", State: "inactive", Action: ""})
}

func irrigationSequenceMode(limitsTrigger, scheduledTrigger bool, cMoist chan float64, moistureLimit, waterAmountLimit, pumpFlow float64, irrigationDuration int) {
	if <-cMoist < moistureLimit {
		fmt.Println("Starting irrigation...🌿🤖🚿")

		req.PostLiveNotify(model.LiveNotify{Title: "Zavlažování", State: "inProgress", Action: "Probíhá zavlažování"})

		// time passed from running pump will be represented as liters
		var flowMeasure float64
		t0 := time.Now()
		// TimeToOverdraw is calculated by deviding amount by flow
		if limitsTrigger {
			for <-cMoist < moistureLimit || flowMeasure < waterAmountLimit/pumpFlow || int(time.Since(t0).Seconds()) > irrigationDuration {
				sens.PUMP.High()
				flowMeasure = float64(time.Since(t0).Seconds())
			}
		}

		if scheduledTrigger {
			for flowMeasure < waterAmountLimit/pumpFlow || int(time.Since(t0).Seconds()) > irrigationDuration {
				sens.PUMP.High()
				flowMeasure = float64(time.Since(t0).Seconds())
			}
		}

		req.PostLiveNotify(model.LiveNotify{Title: "Zavlažování", State: "finished", Action: "Zavlažování dokončeno"})

		sens.PUMP.Low()

		checkingSequence()
	}
}

func irrigationSequence(cMoist chan float64, cPumpState chan bool) {
	// get from DB
	// values only for test
	limitsTrigger := true
	scheduledTrigger := false
	moistureLimit := 50.0
	waterAmountLimit := 50.0
	irrigationDuration := 30 // in seconds
	hourRange := 6

	// Definovaný průtok čerpadla
	var pumpFlow float64 = 1.75 // litr/min

	gocron.Every(1).Seconds().Do(func() {
		if <-cPumpState {
			irrigationState := true
			for irrigationState {
				if <-cPumpState {
					sens.PUMP.High()
				} else {
					sens.PUMP.Low()
					irrigationState = false
				}
			}
		}
	})
	<-gocron.Start()

	if limitsTrigger && !scheduledTrigger {
		gocron.Every(1).Seconds().Do(func() {
			irrigationSequenceMode(true, false, cMoist, moistureLimit, waterAmountLimit, pumpFlow, irrigationDuration)
		})
		<-gocron.Start()
	}

	if scheduledTrigger && !limitsTrigger {
		utils.WaitTillWholeHour()

		gocron.Every(uint64(hourRange)).Hours().Do(func() {
			irrigationSequenceMode(false, true, cMoist, moistureLimit, waterAmountLimit, pumpFlow, irrigationDuration)
		})
		<-gocron.Start()
	}

	if scheduledTrigger && limitsTrigger {
		gocron.Every(1).Seconds().Do(func() {
			irrigationSequenceMode(true, false, cMoist, moistureLimit, waterAmountLimit, pumpFlow, irrigationDuration)
		})
		<-gocron.Start()

		utils.WaitTillWholeHour()

		gocron.Every(uint64(hourRange)).Hours().Do(func() {
			irrigationSequenceMode(false, true, cMoist, moistureLimit, waterAmountLimit, pumpFlow, irrigationDuration)
		})
		<-gocron.Start()
	}
}

func initializationSequence(sensei *sens.Sensors, cMoist chan float64) {
	fmt.Println("Starting initialization sequence...🏁🤖🏁")
	time.Sleep(2000 * time.Millisecond)

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

func MeasurementSequence(sensei *sens.Sensors, cMoist, cTemp, cHum chan float64, cPumpState chan bool) {
	gocron.Every(1).Seconds().Do(func() {
		measurements := sensei.Measure()
		req.PostLiveMeasure(model.LiveMeasure{Moist: measurements.Moist, Hum: measurements.Hum, Temp: measurements.Temp})

		fmt.Printf("\nTemperature: %v˚C", measurements.Temp)
		fmt.Printf("\nHumidity: %v", measurements.Hum)
		fmt.Printf("\nSoil moisture: %v", measurements.Moist)

		cMoist <- math.Floor(measurements.Moist*100) / 100
		cTemp <- math.Floor(measurements.Temp*100) / 100
		cHum <- math.Floor(measurements.Hum*100) / 100

		mid.GetLiveControl(cPumpState)
		mid.LoadLiveMeasure(cMoist, cHum, cTemp)
	},
	)
	<-gocron.Start()
}
