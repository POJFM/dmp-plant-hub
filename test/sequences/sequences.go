package sequences

import (
	"fmt"
	"math"
	"math/rand"

	mid "github.com/SPSOAFM-IT18/dmp-plant-hub/test/middleware"
	"github.com/jasonlvhit/gocron"
)

// on start waits for time to be hour o'clock
// then starts chron routine that is timed on every 4 hours
func SaveOnFourHoursPeriod(temp chan float64) {
	//fmt.Println(temp)
	// for time.Now().Format("04") != "00" {
	// 	// TEST
	// 	fmt.Println(time.Now().Format("04"))
	// 	fmt.Println("Cron\nTemperature: %v˚C", temp)
	// 	// END TEST
	// 	time.Sleep(1 * time.Second)
	// }
	// needs to access values from measurements sequence
	gocron.Every(1).Seconds().Do(func() {
		cTemp := <-temp
		fmt.Printf("Cron\nTemperature: %v˚C", cTemp)
	})
	<-gocron.Start()
}

func MeasurementSequence() {
	gocron.Every(1).Seconds().Do(func() {
		moist := math.Floor(float64(rand.Float64()*3*10)*100) / 100
		hum := math.Floor(float64(rand.Float64()*3*10)*100) / 100
		temp := math.Floor(float64(rand.Float64()*3*10)*100) / 100

		fmt.Printf("\nTemperature: %v˚C", temp)
		fmt.Printf("\nHumidity: %v", hum)
		fmt.Printf("\nSoil moisture: %v", moist)

		mid.LoadLiveMeasure(moist, hum, temp)
		//req.PostLiveMeasure(model.LiveMeasure{Moist: moist, Hum: hum, Temp: temp})
	},
	)
	<-gocron.Start()
}
