package database

import (
	sens "github.com/SPSOAFM-IT18/dmp-plant-hub/sensors"
)

func measurementsAvg(measurements []*sens.Measurements) (measurementsAvg sens.Measurements) {
	// this isn't pretty to look at
	for _, m := range measurements {
		measurementsAvg.Moist += m.Moist
		measurementsAvg.Temp += m.Temp
		measurementsAvg.Hum += m.Hum
	}
	measurementsAvg.Moist = measurementsAvg.Moist / float64(len(measurements))
	measurementsAvg.Temp = measurementsAvg.Temp / float64(len(measurements))
	measurementsAvg.Hum = measurementsAvg.Hum / float64(len(measurements))
	return
}
