package requests

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type LiveControl struct {
	Restart   bool `json:"restart"`
	PumpState bool `json:"pumpState"`
}

func DefaultLiveControl(w http.ResponseWriter, r *http.Request) {
	// data := url.Values{
	// 	"name":       {"John Doe"},
	// 	"occupation": {"gardener"},
	// }

	// resp, err := http.PostForm("https://httpbin.org/post", data)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// var res map[string]interface{}

	// json.NewDecoder(resp.Body).Decode(&res)

	// fmt.Println(res["form"])
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	json.NewEncoder(w).Encode(LiveControl{Restart: false, PumpState: false})
	fmt.Println("live control")
}
