// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Measurement struct {
	Moisture       float64 `json:"moisture"`
	Temperature    float64 `json:"temperature"`
	Humidity       float64 `json:"humidity"`
	WaterLevel     float64 `json:"waterLevel"`
	WaterOverdrawn float64 `json:"waterOverdrawn"`
}

type NewMeasurement struct {
	Moisture       float64 `json:"moisture"`
	Temperature    float64 `json:"temperature"`
	Humidity       float64 `json:"humidity"`
	WaterLevel     float64 `json:"waterLevel"`
	WaterOverdrawn float64 `json:"waterOverdrawn"`
}