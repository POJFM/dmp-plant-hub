package sequences

import (
	"fmt"
	"math"
	"time"

	"github.com/SPSOAFM-IT18/dmp-plant-hub/utils"
	"github.com/TwiN/go-color"
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

// TEST

func InitializationSequence(manualWaterOverdrawn float32, manualWaterLevel float32, waterLevel float32, moistureLevel float32, waterOverdrawnLevel float32, initializationState bool, initialization bool) {
	fmt.Println(color.Ize(color.Green, "Starting initialization sequence...🏁🤖🏁"))
	time.Sleep(2 * time.Millisecond)
	// init measurement
	var waterLevelAvg []float32
	waterLevelAvg = make([]float32, 5)
	var moistureAvg []float32
	moistureAvg = make([]float32, 5)
	// calculating average value
	var count int = 0
	for count < 5 {
		waterLevelAvg = append(waterLevelAvg, waterLevelMeasure())
		moistureAvg = append(moistureAvg, moistureMeasure())
		count += 1
		time.Sleep(1 * time.Millisecond)
	}

	moistureLevel = utils.ArithmeticMean(moistureAvg)

	// send limit values to web api

	// wait for initializationState from web then get values
	var printWait bool = true
	for initializationState {
		if printWait {
			fmt.Println(color.Ize(color.Green, "Waiting for initialization sequence to finish...📝🤖📝"))
			printWait = false
		}
		time.Sleep(1 * time.Millisecond)
	}

	// Check if levels have been manualy set
	if manualWaterLevel > 0 {
		waterLevel = manualWaterLevel
	} else {
		waterLevel = utils.ArithmeticMean(waterLevelAvg) - 2
	}

	if manualWaterOverdrawn > 0 {
		waterOverdrawnLevel = manualWaterOverdrawn
	}

	fmt.Printf(color.Ize(color.Blue, "\nWater level is set to: %vcm"), math.Round(float64(waterLevel)))
	fmt.Println(color.Ize(color.Blue, "\nWater overdrawn level is set to: %vl"), math.Round(float64(waterOverdrawnLevel)))
	fmt.Println(color.Ize(color.Blue, "\nMeasured moisture level: %v%"), math.Round(float64(moistureLevel)))
	time.Sleep(3 * time.Millisecond)
	initialization = true
	fmt.Println(color.Ize(color.Green, "GardenBot is coming to life...✅🤖✅"))
	time.Sleep(1 * time.Millisecond)
}

func MeasurementSequence(PUMP rpio.Pin, LED rpio.Pin, manualWaterOverdrawn float32, manualWaterLevel float32, waterLevel float32, moistureLevel float32, waterOverdrawnLevel float32, pumpFlow float32, initializationState bool, initialization bool) {
	for initialization {
		if moistureMeasure() < moistureLevel {
			fmt.Println(color.Ize(color.Green, "Soil is drying out, starting irrigation...🌿🤖🚿"))

			// time passed from running pump will be represented as liters
			var flowMeasure float32
			t0 := time.Now()
			for waterLevelMeasure() < moistureLevel || flowMeasure < utils.TimeToOverdraw(manualWaterOverdrawn, pumpFlow) {
				//var t1 float32 = time.time()
				PUMP.High()
				flowMeasure = float32(time.Since(t0).Seconds())
			}

			// after pump stops run Checking sequence
			if waterLevelMeasure() < waterLevel {
				fmt.Println(color.Ize(color.Red, "Water tank limit level reached...🚫🤖🚫"))
				// send notification to web and set blinking LED
				for waterLevelMeasure() < waterLevel {
					LED.High()
					time.Sleep(1 * time.Millisecond)
					LED.Low()
					time.Sleep(1 * time.Millisecond)
				}
			} else {
				LED.Low()
			}
		} else {
			PUMP.Low()
		}

		var DHTMeasureValues = DHTMeasure()
		fmt.Println(color.Ize(color.Blue, "\nTemperature: %v˚C"), DHTMeasureValues) //DHTMeasureValues[0]
		fmt.Println(color.Ize(color.Blue, "\nHumidity: %v%"), DHTMeasureValues)     //DHTMeasureValues[1]
		fmt.Println(color.Ize(color.Blue, "\nSoil moisture: %v%"), moistureMeasure())
		time.Sleep(1 * time.Millisecond)
	}
}
