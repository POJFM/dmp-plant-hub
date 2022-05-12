package sequences

import (
	"context"
	"fmt"
	"log"
	"math"
	"time"

	db "github.com/SPSOAFM-IT18/dmp-plant-hub/database"
	graphmodel "github.com/SPSOAFM-IT18/dmp-plant-hub/graph/model"

	mid "github.com/SPSOAFM-IT18/dmp-plant-hub/rest/middleware"

	sens "github.com/SPSOAFM-IT18/dmp-plant-hub/sensors"
	"github.com/SPSOAFM-IT18/dmp-plant-hub/utils"

	"github.com/jasonlvhit/gocron"
)

var (
	gMoist     float64
	gLastMoist float64
	gHum       float64
	gLastHum   float64
	gTemp      float64
	gLastTemp  float64
)

func saveOnFourHoursPeriod(db *db.DB) {
	utils.WaitTillWholeHour()

	err := gocron.Every(4).Hours().Do(func() {
		hum := 0.0
		temp := 0.0
		moist := 0.0
		// average data from a 10-minute interval
		for i := 0; i < 100; i++ {
			hum += gHum
			temp += gTemp
			moist += gMoist
			time.Sleep(6 * time.Second)
		}
		avgHum := hum / 100
		avgTemp := temp / 100
		avgMoist := moist / 100
		hum = math.Floor(avgHum*100) / 100
		temp = math.Floor(avgTemp*100) / 100
		moist = math.Floor(avgMoist*100) / 100
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
	if err != nil {
		log.Printf("Gocron error: %v", err)
	}
	<-gocron.Start()
}

func Controller(db *db.DB, sensei *sens.Sensors) {
	log.Println("Hello welcome to PlantHub...🌿🤖🚿")
	if db.CheckSettings() {
		go measurementSequence(sensei)
		go saveOnFourHoursPeriod(db)
		go irrigationSequence(db, sensei)
	} else {
		go initializationSequence(sensei)
		initializationFinished := true
		for initializationFinished {
			if db.CheckSettings() {
				initializationFinished = false
				go measurementSequence(sensei)
				go saveOnFourHoursPeriod(db)
				go irrigationSequence(db, sensei)
			}
			time.Sleep(10 * time.Second)
		}
	}
}

func CheckingSequence(db *db.DB, sensei *sens.Sensors, flowMeasure, pumpFlow *float64) {
	log.Println("Starting Checking Sequence...🌿🤖🚿")

	settings := db.GetSettingByColumn([]string{"water_level_limit", "default_water_amount"})
	ctx := context.Background()
	waterOverdrawn := 0.0
	waterAmount := 0.0
	waterOverdrawn = float64(*pumpFlow * (*flowMeasure / 1))
	irrigationHistory := db.GetIrrigation(ctx)

	if len(irrigationHistory) == 0 {
		waterAmount = *settings.DefaultWaterAmount - waterOverdrawn
	} else {
		waterAmount = float64(*irrigationHistory[len(irrigationHistory)-1].WaterAmount - waterOverdrawn)
	}

	mid.LoadLiveNotify("Kontrola Nádrže", "inProgress", "Probíhá kontrola nádrže")

	time.Sleep(3000 * time.Millisecond)

	log.Println("namerena nadrz: ", sensei.ReadWaterLevel())

	if sensei.ReadWaterLevel() > *settings.WaterLevelLimit {
		mid.LoadLiveNotify("Doplňte nádrž", "physicalHelpRequired", "Nádrž je prázdná")

		log.Println("Water tank limit level reached...🚫🤖🚫")

		log.Println("namerena nadrz: ", sensei.ReadWaterLevel())
		log.Println("limit nadrze: ", *settings.WaterLevelLimit)

		for sensei.ReadWaterLevel() > *settings.WaterLevelLimit {
			log.Println("doplnit nadrz")
			time.Sleep(1000 * time.Millisecond)
		}

		waterAmount = *settings.DefaultWaterAmount
	}

	waterLevel := fmt.Sprintf("V nádrži zbývá %fl vody", waterAmount)
	// Dodělat na water amount v litrech
	mid.LoadLiveNotify("Kontrola Nádrže", "finished", waterLevel)

	newWaterLevel := sensei.ReadWaterLevel()

	NewirrigationHistory := &graphmodel.NewIrrigation{
		WaterLevel:     &newWaterLevel,
		WaterAmount:    &waterAmount,
		WaterOverdrawn: &waterOverdrawn,
	}

	db.CreateIrrigation(ctx, NewirrigationHistory)

	time.Sleep(3000 * time.Millisecond)

	mid.LoadLiveNotify("", "inactive", "")
}

func irrigationSequenceMode(db *db.DB, sensei *sens.Sensors, limitsTrigger, scheduledTrigger bool, moistureLimit, waterAmountLimit, pumpFlow *float64, irrigationDuration *int) {
	moist := gMoist
	temp := gTemp
	hum := gHum
	irr := true

	if scheduledTrigger {
		measurement := &graphmodel.NewMeasurement{
			Hum:            &hum,
			Temp:           &temp,
			Moist:          &moist,
			WithIrrigation: &irr,
		}
		ctx := context.Background()

		db.CreateMeasurement(ctx, measurement)

		// time passed from running pump will be represented as liters
		var flowMeasure float64
		t0 := time.Now()

		log.Println("Starting irrigation...🌿🤖🚿")

		mid.LoadLiveNotify("Zavlažování", "inProgress", "Probíhá zavlažování")

		sensei.StartPump()

		for flowMeasure < *waterAmountLimit/(*pumpFlow) || int(time.Since(t0).Seconds()) > *irrigationDuration {
			flowMeasure = time.Since(t0).Seconds()
		}

		sensei.StopPump()

		mid.LoadLiveNotify("Zavlažování", "finished", "Zavlažování dokončeno")

		CheckingSequence(db, sensei, &flowMeasure, pumpFlow)
	}

	//if limitsTrigger {
	//	for {
	//		moist := gMoist
	//		temp := gTemp
	//		hum := gHum
	//		// time passed from running pump will be represented as liters
	//		var flowMeasure float64
	//		t0 := time.Now()
	//		if (moist > *moistureLimit) || mid.GetLiveControl() {
	//			log.Println("Starting irrigation...🌿🤖🚿")
	//
	//			measurement := &graphmodel.NewMeasurement{
	//				Hum:            &hum,
	//				Temp:           &temp,
	//				Moist:          &moist,
	//				WithIrrigation: &irr,
	//			}
	//			ctx := context.Background()
	//			db.CreateMeasurement(ctx, measurement)
	//			mid.LoadLiveNotify("Zavlažování", "inProgress", "Probíhá zavlažování")
	//			sensei.StartPump()
	//			if moist > *moistureLimit {
	//				// TimeToOverdraw is calculated by dividing amount by flow
	//				for i := 0; i < *irrigationDuration; i++ {
	//					flowMeasure = time.Since(t0).Seconds()
	//					if moist < *moistureLimit || flowMeasure < *waterAmountLimit/(*pumpFlow) {
	//						i = *irrigationDuration
	//					}
	//					time.Sleep(1 * time.Second)
	//				}
	//			}
	//			sensei.StopPump()
	//			CheckingSequence(db, sensei, &flowMeasure, pumpFlow)
	//		}
	//		time.Sleep(2 * time.Second)
	//	}
	//}
}

func irrigationSequence(db *db.DB, sensei *sens.Sensors) {
	log.Println("Starting PlantHub...🌿🤖🚿")

	settings := db.GetSettingByColumn([]string{"limits_trigger", "scheduled_trigger", "moist_limit", "water_amount_limit", "irrigation_duration", "hour_range"})

	// Definovaný průtok čerpadla
	var pumpFlow = 0.05 // liter / second

	if *settings.LimitsTrigger && !(*settings.ScheduledTrigger) {
		irrigationSequenceMode(db, sensei, true, false, settings.MoistLimit, settings.WaterAmountLimit, &pumpFlow, settings.IrrigationDuration)
	}

	if *settings.ScheduledTrigger && !(*settings.LimitsTrigger) {
		utils.WaitTillWholeHour()

		gocron.Every(uint64(*settings.HourRange)).Hours().Do(func() {
			irrigationSequenceMode(db, sensei, false, true, settings.MoistLimit, settings.WaterAmountLimit, &pumpFlow, settings.IrrigationDuration)
		})
		<-gocron.Start()
	}

	if *settings.ScheduledTrigger && *settings.LimitsTrigger {
		irrigationSequenceMode(db, sensei, true, false, settings.MoistLimit, settings.WaterAmountLimit, &pumpFlow, settings.IrrigationDuration)

		utils.WaitTillWholeHour()

		gocron.Every(uint64(*settings.HourRange)).Hours().Do(func() {
			irrigationSequenceMode(db, sensei, false, true, settings.MoistLimit, settings.WaterAmountLimit, &pumpFlow, settings.IrrigationDuration)
		})
		<-gocron.Start()
	}
}

func initializationSequence(sensei *sens.Sensors) {
	log.Println("Starting initialization sequence...🏁🤖🏁")
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
	//waterLevel := utils.ArithmeticMean(waterLevelAvg)
	waterLevel := sensei.ReadWaterLevel()

	mid.LoadInitMeasured(&moistureLevel, &waterLevel)
}

func measurementSequence(sensei *sens.Sensors) {
	err := gocron.Every(2).Seconds().Do(func() {
		measurements := sensei.Measure()

		temp := math.Floor(measurements.Temp*100) / 100
		hum := math.Floor(measurements.Hum*100) / 100
		moist := math.Floor(measurements.Moist*100) / 100

		// filter out bad data
		if temp <= 0 {
			temp = gLastTemp
		}
		if hum <= 0 {
			hum = gLastHum
		}
		if moist == 100 {
			moist = gLastMoist
		}

		gTemp = temp
		gLastTemp = temp
		gHum = hum
		gLastHum = hum
		gMoist = moist
		gLastMoist = moist

		// fmt.Printf("temp: %f\nhum: %f\nmoi: %f\n", temp, hum, moist)
		// go fmt.Printf("sonic: %f\n", sensei.ReadWaterLevel())

		mid.LoadLiveMeasure(&moist, &hum, &temp)
	},
	)
	if err != nil {
		log.Printf("Gocron error: %v", err)
	}
	<-gocron.Start()
}
