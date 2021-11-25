package utils

func ArithmeticMean(list []float32) float32 {
	// maybe make it into list map function
	var total float32
	for _, v := range list {
		total += v
	}
	return total / float32(len(list))
}

func TimeToOverdraw(manualWaterOverdrawn float32, pumpFlow float32) float32 {
	return manualWaterOverdrawn / pumpFlow
}
