package model

type InitMeasured struct {
	MoistLimit      float32 `json:"moistLimit"`
	WaterLevelLimit float32 `json:"waterLevelLimit"`
}

type LiveMeasure struct {
	Moist float32 `json:"moist"`
	Hum   float32 `json:"hum"`
	Temp  float32 `json:"temp"`
}

type LiveNotify struct {
	Title  string `json:"title"`
	State  string `json:"state"`
	Action string `json:"action"`
}

type LiveControl struct {
	Restart   bool `json:"restart"`
	PumpState bool `json:"pumpState"`
}
