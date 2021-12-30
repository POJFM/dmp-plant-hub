package sequences

import (
	"encoding/json"
	"fmt"
	"github.com/SPSOAFM-IT18/dmp-plant-hub/sensors"
	"net/http"
	"time"
)

type Measurements struct {
	WaterLevel float32 `json:"waterLevel"`
	Moist      float32 `json:"moist"`
	Hum        float32 `json:"hum"`
	Temp       float32 `json:"temp"`
}

var liveMeasurements Measurements

/*func buildInitMeasured(w http.ResponseWriter, r *http.Request) {
	fmt.Println("init measured")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(initMeasuredData)
}*/

func BuildLiveMeasure(w http.ResponseWriter, r *http.Request) {
	fmt.Println("live measured data")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(liveMeasurements)
}

/*func buildLiveNotify(w http.ResponseWriter, r *http.Request) {
	fmt.Println("live notify")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(liveNotifyData)
}*/

func Measure(sens *sensors.PinOut) {
	for range time.Tick(time.Second * 1) {
		liveMeasurements = sens.Measure()
	}
}
