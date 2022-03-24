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

	d "code.cloudfoundry.org/go-diodes"
)

func saveOnFourHoursPeriod(db *db.DB, dMoist, dTemp, dHum *d.OneToOne) {
	utils.WaitTillWholeHour()

	gocron.Every(4).Hours().Do(func() {
		pdMoist := d.NewPoller(dMoist)
		pdTemp := d.NewPoller(dTemp)
		pdHum := d.NewPoller(dHum)

		moist := *(*float64)(pdMoist.Next())
		hum := *(*float64)(pdTemp.Next())
		temp := *(*float64)(pdHum.Next())
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

func Controller(db *db.DB, sensei *sens.Sensors, dMoist, dTemp, dHum *d.OneToOne) {
	fmt.Println("Hello welome to PlantHub...🌿🤖🚿")
	if db.CheckSettings() {
		go measurementSequence(sensei, dMoist, dTemp, dHum)
		go saveOnFourHoursPeriod(db, dMoist, dTemp, dHum)
		go irrigationSequence(db, sensei, dMoist, dTemp, dHum)
	} else {
		go initializationSequence(sensei)
		go measurementSequence(sensei, dMoist, dTemp, dHum)
		initializationFinished := true
		for initializationFinished {
			// stopLED := make(chan bool)
			// go func() {
			// 	for {
			// 		select {
			// 		case <-stopLED:
			// 			return
			// 		default:
			// 			// for i := 0; i < 2; i++ {
			// 			// 	sensei.StartLED()
			// 			// 	time.Sleep(500 * time.Millisecond)
			// 			// 	sensei.StopLED()
			// 			// 	time.Sleep(500 * time.Millisecond)
			// 			// }
			// 			time.Sleep(1500 * time.Millisecond)
			// 		}
			// 	}
			// }()
			if db.CheckSettings() {
				initializationFinished = false
				//stopLED <- true
				go measurementSequence(sensei, dMoist, dTemp, dHum)
				go saveOnFourHoursPeriod(db, dMoist, dTemp, dHum)
				go irrigationSequence(db, sensei, dMoist, dTemp, dHum)
			}
			time.Sleep(1000 * time.Millisecond)
		}
	}
}

func checkingSequence(db *db.DB, sensei *sens.Sensors) {
	settings := db.GetSettingByColumn([]string{"water_level_limit"})

	mid.LoadLiveNotify("Kontrola Nádrže", "inProgress", "Probíhá kontrola nádrže")
	//req.PostLiveNotify(model.LiveNotify{Title: "Kontrola Nádrže", State: "inProgress", Action: "Probíhá kontrola nádrže"})

	time.Sleep(3000 * time.Millisecond)

	if sensei.ReadWaterLevel() < *settings.WaterLevelLimit {
		mid.LoadLiveNotify("Doplňte nádrž", "physicalHelpRequired", "Nádrž je prázdná")
		//req.PostLiveNotify(model.LiveNotify{Title: "Doplňte nádrž", State: "physicalHelpRequired", Action: "Nádrž je prázdná"})

		fmt.Println("Water tank limit level reached...🚫🤖🚫")

		for sensei.ReadWaterLevel() < *settings.WaterLevelLimit {
			sensei.StartLED()
			time.Sleep(1000 * time.Millisecond)
			sensei.StopLED()
			time.Sleep(1000 * time.Millisecond)
		}
	}

	waterLevel := fmt.Sprintf("V nádrži zbývá %fl vody", sensei.ReadWaterLevel())
	// Dodělat na water amount v litrech
	mid.LoadLiveNotify("Kontrola Nádrže", "finished", waterLevel)
	//req.PostLiveNotify(model.LiveNotify{Title: "Kontrola Nádrže", State: "finished", Action: waterLevel})

	time.Sleep(3000 * time.Millisecond)

	mid.LoadLiveNotify("", "inactive", "")
	//req.PostLiveNotify(model.LiveNotify{Title: "", State: "inactive", Action: ""})
}

func irrigationSequenceMode(db *db.DB, sensei *sens.Sensors, dMoist, dTemp, dHum *d.OneToOne, limitsTrigger, scheduledTrigger bool, moistureLimit, waterAmountLimit, pumpFlow *float64, irrigationDuration *int) {
	pdMoist := d.NewPoller(dMoist)
	pdTemp := d.NewPoller(dTemp)
	pdHum := d.NewPoller(dHum)

	moist := *(*float64)(pdMoist.Next())
	hum := *(*float64)(pdTemp.Next())
	temp := *(*float64)(pdHum.Next())
	irr := true

	measurement := &graphmodel.NewMeasurement{
		Hum:            &hum,
		Temp:           &temp,
		Moist:          &moist,
		WithIrrigation: &irr,
	}
	ctx := context.Background()

	db.CreateMeasurement(ctx, measurement)

	if scheduledTrigger {
		// time passed from running pump will be represented as liters
		var flowMeasure float64
		t0 := time.Now()

		fmt.Println("Starting irrigation...🌿🤖🚿")
		mid.LoadLiveNotify("Zavlažování", "inProgress", "Probíhá zavlažování")

		for flowMeasure < *waterAmountLimit/(*pumpFlow) || int(time.Since(t0).Seconds()) > *irrigationDuration {
			sensei.StartPump()
			flowMeasure = float64(time.Since(t0).Seconds())
		}

		sensei.StopPump()

		mid.LoadLiveNotify("Zavlažování", "finished", "Zavlažování dokončeno")

		checkingSequence(db, sensei)
	}

	if limitsTrigger {
		for {
			// time passed from running pump will be represented as liters
			var flowMeasure float64
			t0 := time.Now()

			if moist < *moistureLimit {
				fmt.Println("Starting irrigation...🌿🤖🚿")

				mid.LoadLiveNotify("Zavlažování", "inProgress", "Probíhá zavlažování")
				//req.PostLiveNotify(model.LiveNotify{Title: "Zavlažování", State: "inProgress", Action: "Probíhá zavlažování"})

				// TimeToOverdraw is calculated by deviding amount by flow
				for moist < *moistureLimit || flowMeasure < *waterAmountLimit/(*pumpFlow) || int(time.Since(t0).Seconds()) > *irrigationDuration {
					sensei.StartPump()
					flowMeasure = float64(time.Since(t0).Seconds())
				}

				//req.PostLiveNotify(model.LiveNotify{Title: "Zavlažování", State: "finished", Action: "Zavlažování dokončeno"})

				sensei.StopPump()

				checkingSequence(db, sensei)
			}
			time.Sleep(1 * time.Minute)
		}
	}
}

func irrigationSequence(db *db.DB, sensei *sens.Sensors, dMoist, dTemp, dHum *d.OneToOne) {
	fmt.Println("Starting PlantHub...🌿🤖🚿")

	settings := db.GetSettingByColumn([]string{"limits_trigger", "scheduled_trigger", "moist_limit", "water_amount_limit", "irrigation_duration", "hour_range"})

	// Definovaný průtok čerpadla
	var pumpFlow float64 = 1.75 // litr/min

	if *settings.LimitsTrigger && !(*settings.ScheduledTrigger) {
		irrigationSequenceMode(db, sensei, dMoist, dTemp, dHum, true, false, settings.MoistLimit, settings.WaterAmountLimit, &pumpFlow, settings.IrrigationDuration)
	}

	if *settings.ScheduledTrigger && !(*settings.LimitsTrigger) {
		utils.WaitTillWholeHour()

		gocron.Every(uint64(*settings.HourRange)).Hours().Do(func() {
			irrigationSequenceMode(db, sensei, dMoist, dTemp, dHum, false, true, settings.MoistLimit, settings.WaterAmountLimit, &pumpFlow, settings.IrrigationDuration)
		})
		<-gocron.Start()
	}

	if *settings.ScheduledTrigger && *settings.LimitsTrigger {
		irrigationSequenceMode(db, sensei, dMoist, dTemp, dHum, true, false, settings.MoistLimit, settings.WaterAmountLimit, &pumpFlow, settings.IrrigationDuration)

		utils.WaitTillWholeHour()

		gocron.Every(uint64(*settings.HourRange)).Hours().Do(func() {
			irrigationSequenceMode(db, sensei, dMoist, dTemp, dHum, false, true, settings.MoistLimit, settings.WaterAmountLimit, &pumpFlow, settings.IrrigationDuration)
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

	mid.LoadInitMeasured(&moistureLevel, &waterLevel)
	//req.PostInitMeasured(model.InitMeasured{MoistLimit: moistureLevel, WaterLevelLimit: waterLevel})
}

func measurementSequence(sensei *sens.Sensors, dMoist, dTemp, dHum *d.OneToOne) {
	gocron.Every(2).Seconds().Do(func() {
		measurements := sensei.Measure()
		fmt.Printf("temp: %f\nhum: %f\nmoi: %f\n", measurements.Temp, measurements.Hum, measurements.Moist)

		moist := math.Floor(measurements.Moist*100) / 100
		temp := math.Floor(measurements.Temp*100) / 100
		hum := math.Floor(measurements.Hum*100) / 100

		dMoist.Set(d.GenericDataType(&moist))
		dTemp.Set(d.GenericDataType(&temp))
		dHum.Set(d.GenericDataType(&hum))

		mid.LoadLiveMeasure(dMoist, dHum, dTemp)
	},
	)
	<-gocron.Start()
}
