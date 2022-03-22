package requests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/Jeffail/gabs"
	"github.com/SPSOAFM-IT18/dmp-plant-hub/env"
	"github.com/SPSOAFM-IT18/dmp-plant-hub/rest/model"
)

func PostInitMeasured(rawData model.InitMeasured) {
	data := model.InitMeasured{
		MoistLimit:      rawData.MoistLimit,
		WaterLevelLimit: rawData.WaterLevelLimit,
	}

	json_data, err := json.Marshal(data)

	if err != nil {
		log.Print(err)
	}

	resp, err := http.Post(env.Process("GO_REST_API_URL")+"/init/measured", "application/json",
		bytes.NewBuffer(json_data))

	if err != nil {
		log.Print(err)
	}

	var res map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&res)

	fmt.Println(res["json"])
}

func PostLiveMeasure(rawData model.LiveMeasure) {
	data := model.LiveMeasure{
		Moist: rawData.Moist,
		Hum:   rawData.Hum,
		Temp:  rawData.Temp,
	}

	jsonData, err := json.Marshal(data)

	if err != nil {
		log.Print(err)
	}

	resp, err := http.Post(env.Process("GO_REST_API_URL")+"/live/measure", "application/json",
		bytes.NewBuffer(jsonData))

	if err != nil {
		log.Print(err)
	}

	var res map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&res)

	fmt.Println(res["json"])
}

func PostLiveNotify(rawData model.LiveNotify) {
	data := model.LiveNotify{
		Title:  rawData.Title,
		State:  rawData.State,
		Action: rawData.Action,
	}

	json_data, err := json.Marshal(data)

	if err != nil {
		log.Print(err)
	}

	resp, err := http.Post(env.Process("GO_REST_API_URL")+"/live/notify", "application/json",
		bytes.NewBuffer(json_data))

	if err != nil {
		log.Print(err)
	}

	var res map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&res)

	fmt.Println(res["json"])
}

func PostLiveControl(rawData model.LiveControl) {
	data := model.LiveControl{
		Restart:   rawData.Restart,
		PumpState: rawData.PumpState,
	}

	json_data, err := json.Marshal(data)

	if err != nil {
		log.Print(err)
	}

	resp, err := http.Post(env.Process("GO_REST_API_URL")+"/live/control", "application/json",
		bytes.NewBuffer(json_data))

	if err != nil {
		log.Print(err)
	}

	var res map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&res)

	fmt.Println(res["json"])
}

func GetWeatherForecast() {
	url := "http://api.open-notify.org/astros.json"

	spaceClient := http.Client{
		Timeout: time.Second * 2, // Timeout after 2 seconds
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "spacecount-tutorial")

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

	jsonParsed, err := gabs.ParseJSON(body)
	if err != nil {
		panic(err)
	}

	fmt.Println("Get value of Protected:\t", jsonParsed.Path("employees.protected").Data())
}
