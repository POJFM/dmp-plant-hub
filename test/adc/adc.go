// https://cs.amen-technologies.com/how-does-successive-approximation-adc-work
package adc

import "fmt"

func Adc() {
	fmt.Println("\nADC: ")

	const vin = 7.5
	const bits = 10
	vlsb := vin + (vin / 2)
	vdiff := vin - (vin / 2)
	binCode := 10

	// Default binary code
	for i := 1; i < bits-1; i++ {
		binCode *= 10
	}

	// Succesive approximation register
	// N Comparator
	for i := 0; i < bits-1; i++ {
		vdiff /= 2
		vlsb -= vdiff
		if vdiff == vin {
			i = bits - 1
		}
	}

	fmt.Println("VDAC: ", vlsb)
	fmt.Println("VDAC binary code: ", binCode)
}
