package sequences

import (
	"fmt"
	"math/rand"

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
		fmt.Println("Cron\nTemperature: %v˚C", cTemp)
	})
	<-gocron.Start()
}

func MeasurementSequence(cTemp chan float64) {
	gocron.Every(1).Seconds().Do(func() {
		temperature := float64(rand.Float64() * 5)

		fmt.Println("Measure\nTemperature: %v˚C", temperature, cTemp)

		cTemp <- temperature

		temp := <-cTemp

		fmt.Println(temp)
	},
	)
	<-gocron.Start()
}
