package sequences

import (
	"context"
	"fmt"
	"math"
	"strconv"
	"time"

	db "github.com/SPSOAFM-IT18/dmp-plant-hub/database"
	graphmodel "github.com/SPSOAFM-IT18/dmp-plant-hub/graph/model"

	mid "github.com/SPSOAFM-IT18/dmp-plant-hub/rest/middleware"

	sens "github.com/SPSOAFM-IT18/dmp-plant-hub/sensors"
	"github.com/SPSOAFM-IT18/dmp-plant-hub/utils"

	"github.com/jasonlvhit/gocron"
)

var (
	gMoist float64
	gHum   float64
	gTemp  float64
)

func saveOnFourHoursPeriod(db *db.DB) {
	utils.WaitTillWholeHour()

	gocron.Every(4).Hours().Do(func() {
		moist := gMoist
		temp := gTemp
		hum := gHum
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

func Controller(db *db.DB, sensei *sens.Sensors) {
	fmt.Println("Hello welome to PlantHub...🌿🤖🚿")
	if db.CheckSettings() {
		go measurementSequence(sensei)
		go saveOnFourHoursPeriod(db)
		go irrigationSequence(db, sensei)
	} else {
		go initializationSequence(sensei)
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
				go measurementSequence(sensei)
				go saveOnFourHoursPeriod(db)
				go irrigationSequence(db, sensei)
			}
			time.Sleep(1000 * time.Millisecond)
		}
	}
}

func CheckingSequence(db *db.DB, sensei *sens.Sensors) {
	fmt.Println("Starting Checking Sequence...🌿🤖🚿")

	//settings := db.GetSettingByColumn([]string{"water_level_limit"})

	mid.LoadLiveNotify("Kontrola Nádrže", "inProgress", "Probíhá kontrola nádrže")
	//req.PostLiveNotify(model.LiveNotify{Title: "Kontrola Nádrže", State: "inProgress", Action: "Probíhá kontrola nádrže"})

	time.Sleep(3000 * time.Millisecond)

	// if sensei.ReadWaterLevel() < *settings.WaterLevelLimit {
	// 	mid.LoadLiveNotify("Doplňte nádrž", "physicalHelpRequired", "Nádrž je prázdná")
	// 	//req.PostLiveNotify(model.LiveNotify{Title: "Doplňte nádrž", State: "physicalHelpRequired", Action: "Nádrž je prázdná"})

	// 	fmt.Println("Water tank limit level reached...🚫🤖🚫")

	// 	fmt.Println("namerena nadrz: ", sensei.ReadWaterLevel())
	// 	fmt.Println("limit nadrze: ", *settings.WaterLevelLimit)

	// 	for sensei.ReadWaterLevel() < *settings.WaterLevelLimit {
	// 		// sensei.StartLED()
	// 		// time.Sleep(1000 * time.Millisecond)
	// 		// sensei.StopLED()
	// 		fmt.Println("dopln nadrz chuju")
	// 		time.Sleep(1000 * time.Millisecond)
	// 	}
	// }

	waterLevel := fmt.Sprintf("V nádrži zbývá %fl vody", sensei.ReadWaterLevel())
	// Dodělat na water amount v litrech
	mid.LoadLiveNotify("Kontrola Nádrže", "finished", waterLevel)
	//req.PostLiveNotify(model.LiveNotify{Title: "Kontrola Nádrže", State: "finished", Action: waterLevel})

	time.Sleep(3000 * time.Millisecond)

	mid.LoadLiveNotify("", "inactive", "")
	//req.PostLiveNotify(model.LiveNotify{Title: "", State: "inactive", Action: ""})
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

		fmt.Println("Starting irrigation...🌿🤖🚿")

		mid.LoadLiveNotify("Zavlažování", "inProgress", "Probíhá zavlažování")

		sensei.StartPump()

		for flowMeasure < *waterAmountLimit/(*pumpFlow) || int(time.Since(t0).Seconds()) > *irrigationDuration {
			flowMeasure = float64(time.Since(t0).Seconds())
		}

		sensei.StopPump()

		mid.LoadLiveNotify("Zavlažování", "finished", "Zavlažování dokončeno")

		CheckingSequence(db, sensei)
	}

	if limitsTrigger {
		for {
			moist := gMoist
			temp := gTemp
			hum := gHum

			// time passed from running pump will be represented as liters
			var flowMeasure float64
			t0 := time.Now()

			// fmt.Println("kokot debil")

			// fmt.Println("moist: ", moistt)
			// fmt.Println("moistt: ", tempt)
			// fmt.Println("moisttt: ", humt)
			// fmt.Println("moistlimit: ", *moistureLimit)
			// fmt.Println("irrDur: ", *irrigationDuration)

			if (moist > *moistureLimit) || mid.GetLiveControl() {
				fmt.Println("Starting irrigation...🌿🤖🚿")

				measurement := &graphmodel.NewMeasurement{
					Hum:            &hum,
					Temp:           &temp,
					Moist:          &moist,
					WithIrrigation: &irr,
				}
				ctx := context.Background()

				db.CreateMeasurement(ctx, measurement)

				mid.LoadLiveNotify("Zavlažování", "inProgress", "Probíhá zavlažování")
				//req.PostLiveNotify(model.LiveNotify{Title: "Zavlažování", State: "inProgress", Action: "Probíhá zavlažování"})

				sensei.StartPump()

				// fmt.Println("moist: ", moistt)
				// fmt.Println("moistlimit: ", *moistureLimit)
				// fmt.Println("irrDur: ", *irrigationDuration)

				if moist > *moistureLimit {
					// TimeToOverdraw is calculated by dividing amount by flow
					for i := 0; i < *irrigationDuration; i++ {
						flowMeasure = float64(time.Since(t0).Seconds())
						fmt.Println(i)
						//fmt.Println(int(time.Since(t0).Seconds()))
						if moist < *moistureLimit || flowMeasure < *waterAmountLimit/(*pumpFlow) {
							i = *irrigationDuration
						}
						time.Sleep(1 * time.Second)
					}
				}

				fmt.Println("přejeb")

				//req.PostLiveNotify(model.LiveNotify{Title: "Zavlažování", State: "finished", Action: "Zavlažování dokončeno"})

				sensei.StopPump()

				CheckingSequence(db, sensei)
			}
			time.Sleep(2 * time.Second)
		}
	}
}

func irrigationSequence(db *db.DB, sensei *sens.Sensors) {
	fmt.Println("Starting PlantHub...🌿🤖🚿")

	settings := db.GetSettingByColumn([]string{"limits_trigger", "scheduled_trigger", "moist_limit", "water_amount_limit", "irrigation_duration", "hour_range"})

	fmt.Println("hook 1")
	// Definovaný průtok čerpadla
	var pumpFlow float64 = 1.75 // litr/min

	if *settings.LimitsTrigger && !(*settings.ScheduledTrigger) {
		fmt.Println("hook 2")

		irrigationSequenceMode(db, sensei, true, false, settings.MoistLimit, settings.WaterAmountLimit, &pumpFlow, settings.IrrigationDuration)
		fmt.Println("hook 3")
	}

	if *settings.ScheduledTrigger && !(*settings.LimitsTrigger) {
		fmt.Println("hook 4")

		utils.WaitTillWholeHour()

		gocron.Every(uint64(*settings.HourRange)).Hours().Do(func() {
			irrigationSequenceMode(db, sensei, false, true, settings.MoistLimit, settings.WaterAmountLimit, &pumpFlow, settings.IrrigationDuration)
		})
		<-gocron.Start()
	}

	if *settings.ScheduledTrigger && *settings.LimitsTrigger {
		fmt.Println("hook 5")

		irrigationSequenceMode(db, sensei, true, false, settings.MoistLimit, settings.WaterAmountLimit, &pumpFlow, settings.IrrigationDuration)

		utils.WaitTillWholeHour()

		gocron.Every(uint64(*settings.HourRange)).Hours().Do(func() {
			irrigationSequenceMode(db, sensei, false, true, settings.MoistLimit, settings.WaterAmountLimit, &pumpFlow, settings.IrrigationDuration)
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

func measurementSequence(sensei *sens.Sensors) {
	gocron.Every(2).Seconds().Do(func() {
		measurements := sensei.Measure()

		// 550 as highest limit, therefore 100%
		moiststr, _ := strconv.ParseFloat(strconv.FormatFloat(100*(math.Floor(measurements.Moist*100)/100)/500, 'f', -2, 64), 64)
		moist, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", moiststr), 64)

		temp := math.Floor(measurements.Temp*100) / 100
		hum := math.Floor(measurements.Hum*100) / 100

		fmt.Printf("temp: %f\nhum: %f\nmoi: %f\n", temp, hum, moist)

		fmt.Printf("moist floor: %f\n", math.Floor(measurements.Moist*100)/100)

		gMoist = moist
		gTemp = temp
		gHum = hum

		mid.LoadLiveMeasure(&moist, &hum, &temp)
	},
	)
	<-gocron.Start()
}
