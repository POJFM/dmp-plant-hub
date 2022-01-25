package requests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

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
		log.Fatal(err)
	}

	resp, err := http.Post(env.Process("GO_API_URL")+"/init/measured", "application/json",
		bytes.NewBuffer(json_data))

	if err != nil {
		log.Fatal(err)
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

	json_data, err := json.Marshal(data)

	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post(env.Process("GO_API_URL")+"/live/measure", "application/json",
		bytes.NewBuffer(json_data))

	if err != nil {
		log.Fatal(err)
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
		log.Fatal(err)
	}

	resp, err := http.Post(env.Process("GO_API_URL")+"/live/notify", "application/json",
		bytes.NewBuffer(json_data))

	if err != nil {
		log.Fatal(err)
	}

	var res map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&res)

	fmt.Println(res["json"])
}

func GetLiveControl() []byte {
	resp, err := http.Get(env.Process("GO_API_URL") + "/live/control")

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))
	return body
}

func PostLiveControl(rawData model.LiveControl) {
	data := model.LiveControl{
		Restart:   rawData.Restart,
		PumpState: rawData.PumpState,
	}

	json_data, err := json.Marshal(data)

	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post(env.Process("GO_API_URL")+"/live/control", "application/json",
		bytes.NewBuffer(json_data))

	if err != nil {
		log.Fatal(err)
	}

	var res map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&res)

	fmt.Println(res["json"])
}
