package sequences

import (
	"context"
	"fmt"
	"math"
	"time"

	db "github.com/SPSOAFM-IT18/dmp-plant-hub/database"
	graphmodel "github.com/SPSOAFM-IT18/dmp-plant-hub/graph/model"

	mid "github.com/SPSOAFM-IT18/dmp-plant-hub/rest/middleware"

	sens "github.com/SPSOAFM-IT18/dmp-plant-hub/sensors"
	"github.com/SPSOAFM-IT18/dmp-plant-hub/utils"

	"github.com/jasonlvhit/gocron"
)

func saveOnFourHoursPeriod(db *db.DB, cMoist, cTemp, cHum chan float64) {
	utils.WaitTillWholeHour()

	gocron.Every(4).Hours().Do(func() {
		hum := <-cMoist
		temp := <-cTemp
		moist := <-cMoist
		irr := false
		measurement := &graphmodel.NewMeasurement{
			Hum:            &hum,
			Temp:           &temp,
			Moist:          &moist,
			WithIrrigation: &irr,
		}
		ctx := context.Background()
		db.CreateMeasurement(ctx, measurement)
	})
	<-gocron.Start()
}

func Controller(db *db.DB, sensei *sens.Sensors, cMoist, cTemp, cHum chan float64, cPumpState chan bool) {
	if db.CheckSettings() {
		go measurementSequence(sensei, cMoist, cTemp, cHum, cPumpState)
		go saveOnFourHoursPeriod(db, cMoist, cTemp, cHum)
		go irrigationSequence(sensei, cMoist, cPumpState)
	} else {
		go initializationSequence(sensei)
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
				go measurementSequence(sensei, cMoist, cTemp, cHum, cPumpState)
				go saveOnFourHoursPeriod(db, cMoist, cTemp, cHum)
				go irrigationSequence(sensei, cMoist, cPumpState)
			}
			time.Sleep(1000 * time.Millisecond)
		}
	}
}

func checkingSequence(sensei *sens.Sensors) {
	// get from DB
	// values only for test
	const waterLevelLimit = 75

	mid.LoadLiveNotify("Kontrola Nádrže", "inProgress", "Probíhá kontrola nádrže")
	//req.PostLiveNotify(model.LiveNotify{Title: "Kontrola Nádrže", State: "inProgress", Action: "Probíhá kontrola nádrže"})

	time.Sleep(3000 * time.Millisecond)

	if sensei.ReadWaterLevel() < waterLevelLimit {
		mid.LoadLiveNotify("Doplňte nádrž", "physicalHelpRequired", "Nádrž je prázdná")
		//req.PostLiveNotify(model.LiveNotify{Title: "Doplňte nádrž", State: "physicalHelpRequired", Action: "Nádrž je prázdná"})

		fmt.Println("Water tank limit level reached...🚫🤖🚫")

		for sensei.ReadWaterLevel() < waterLevelLimit {
			sens.LED.High()
			time.Sleep(1000 * time.Millisecond)
			sens.LED.Low()
			time.Sleep(1000 * time.Millisecond)
		}
	}

	waterLevel := fmt.Sprintf("V nádrži zbývá %fl vody", sensei.ReadWaterLevel())
	// Dodělat na water amount v litrech
	mid.LoadLiveNotify("Kontrola Nádrže", "finishedphysicalHelpRequired", waterLevel)
	//req.PostLiveNotify(model.LiveNotify{Title: "Kontrola Nádrže", State: "finished", Action: waterLevel})

	time.Sleep(3000 * time.Millisecond)

	mid.LoadLiveNotify("", "inactive", "")
	//req.PostLiveNotify(model.LiveNotify{Title: "", State: "inactive", Action: ""})
}

func irrigationSequenceMode(sensei *sens.Sensors, limitsTrigger, scheduledTrigger bool, cMoist chan float64, moistureLimit, waterAmountLimit, pumpFlow float64, irrigationDuration int) {
	if <-cMoist < moistureLimit {
		fmt.Println("Starting irrigation...🌿🤖🚿")

		mid.LoadLiveNotify("Zavlažování", "inProgress", "Probíhá zavlažování")
		//req.PostLiveNotify(model.LiveNotify{Title: "Zavlažování", State: "inProgress", Action: "Probíhá zavlažování"})

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

		mid.LoadLiveNotify("Zavlažování", "finished", "Zavlažování dokončeno")
		//req.PostLiveNotify(model.LiveNotify{Title: "Zavlažování", State: "finished", Action: "Zavlažování dokončeno"})

		sens.PUMP.Low()

		checkingSequence(sensei)
	}
}

func irrigationSequence(sensei *sens.Sensors, cMoist chan float64, cPumpState chan bool) {
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
			irrigationSequenceMode(sensei, true, false, cMoist, moistureLimit, waterAmountLimit, pumpFlow, irrigationDuration)
		})
		<-gocron.Start()
	}

	if scheduledTrigger && !limitsTrigger {
		utils.WaitTillWholeHour()

		gocron.Every(uint64(hourRange)).Hours().Do(func() {
			irrigationSequenceMode(sensei, false, true, cMoist, moistureLimit, waterAmountLimit, pumpFlow, irrigationDuration)
		})
		<-gocron.Start()
	}

	if scheduledTrigger && limitsTrigger {
		gocron.Every(1).Seconds().Do(func() {
			irrigationSequenceMode(sensei, true, false, cMoist, moistureLimit, waterAmountLimit, pumpFlow, irrigationDuration)
		})
		<-gocron.Start()

		utils.WaitTillWholeHour()

		gocron.Every(uint64(hourRange)).Hours().Do(func() {
			irrigationSequenceMode(sensei, false, true, cMoist, moistureLimit, waterAmountLimit, pumpFlow, irrigationDuration)
		})
		<-gocron.Start()
	}
}

func initializationSequence(sensei *sens.Sensors) {
	fmt.Println("Starting initialization sequence...🏁🤖🏁")
	time.Sleep(2000 * time.Millisecond)

	var waterLevelAvg []float64
	waterLevelAvg = make([]float64, 5)
	var moistureAvg []float64
	moistureAvg = make([]float64, 5)

	// calculating average value
	for i := 0; i < 5; i++ {
		moistureAvg = append(moistureAvg, sensei.ReadMoisture())
		waterLevelAvg = append(waterLevelAvg, sensei.ReadWaterLevel())
		time.Sleep(1000 * time.Millisecond)
	}

	moistureLevel := utils.ArithmeticMean(moistureAvg)
	waterLevel := utils.ArithmeticMean(waterLevelAvg)

	mid.LoadInitMeasured(moistureLevel, waterLevel)
	//req.PostInitMeasured(model.InitMeasured{MoistLimit: moistureLevel, WaterLevelLimit: waterLevel})
}

func measurementSequence(sensei *sens.Sensors, cMoist, cTemp, cHum chan float64, cPumpState chan bool) {
	gocron.Every(1).Seconds().Do(func() {
		fmt.Println("kokot")
		measurements := sensei.Measure()

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
