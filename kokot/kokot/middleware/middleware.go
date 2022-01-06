package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type LiveControl struct {
	Restart   bool `json:"restart,omitempty"`
	PumpState bool `json:"pumpState,omitempty"`
}

// CreateTask create task route
func PostLiveControl(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Accept", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.WriteHeader(http.StatusOK)

	var task LiveControl
	_ = json.NewDecoder(r.Body).Decode(&task)
	fmt.Print("POST from Web app: ", task)
}
