package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/Jeffail/gabs"
	"github.com/SPSOAFM-IT18/dmp-plant-hub/test/env"
	"github.com/SPSOAFM-IT18/dmp-plant-hub/test/router"
	//mid "github.com/SPSOAFM-IT18/dmp-plant-hub/test/middleware"
)

func getWeatherForecast() {
	url := "https://api.openweathermap.org/data/2.5/onecall?lat=49.68333&lon=18.35&exclude=daily,minutely,alerts&units=metric&appid=db1c8036cb8ff62962107df3c0ea3171"

	spaceClient := http.Client{
		Timeout: time.Second * 2, // Timeout after 2 seconds
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	//req.Header.Set("User-Agent", "spacecount-tutorial")

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	var resM map[string]interface{}

	json.NewDecoder(res.Body).Decode(&resM)

	jsonParsed, err := gabs.ParseJSON(body)
	if err != nil {
		panic(err)
	}

	fmt.Println(resM["json"])

	fmt.Println("Get value of Protected:\t", jsonParsed.Path("hourly").Data())
}

func kokot(cPumpState chan bool) {
	//requests.PostLiveControl(model.LiveControl{Restart: false, PumpState: false})

	for {
		time.Sleep(1 * time.Second)

		// var kokotismus = true

		// if kokotismus == true {
		// 	kokotismus = false
		// 	time.AfterFunc(3*time.Second, func() {
		// 		hours := time.Now().format
		// 		fmt.Println(hours)
		// 		kokotismus = true
		// 	})
		// }
		//requests.PostLiveNotify(model.LiveNotify{Title: "Zavlažování", State: "inProgress", Action: "Probíhá zavlažování"})
		//requests.GetLiveControl()

		fmt.Println("", <-cPumpState)

		//go kokoti()
		//go requests.PostLiveNotify(model.LiveNotify{Title: "kokot jsi", State: "active", Action: "debil"})
	}
}

func main() {
	//cPumpState := make(chan bool, 1)

	r := router.Router()

	getWeatherForecast()

	//go seq.MeasurementSequence()

	// adc.Adc()

	// port := fmt.Sprint(":" + env.Process("GO_API_PORT"))
	// fmt.Print(port)
	fmt.Println("Starting server on the port", env.Process("GO_API_PORT"))
	//go kokot(cPumpState)
	//go executeCronJob()
	log.Fatal(http.ListenAndServe(":"+env.Process("GO_API_PORT"), r))
}
