package drivers

import (
	"errors"
	"math"
	"time"
)

// fetch from mq135 data sheet:
// temperature: x=[-10 0 10 20 30 40 50]
// humidity: y=[30 60 85]
// z=[[1.71,1.44,1.25];[1.59,1.35,1.17];[1.42,1.2,1.05];[1.25,1.07,0.93];[1.15,0.98,0.85];[1.0,0.85,0.73];[0.87,0.74,0.63]]
// z = p00 + p10*x + p01*y + p20*x^2 + p11*x*y
// p00 =       1.789  (1.761, 1.818)
// p10 =     -0.0165  (-0.01782, -0.01518)
// p01 =   -0.007536  (-0.007996, -0.007076)
// p20 =   9.921e-06  (-1.118e-05, 3.102e-05)
// p11 =   6.723e-05  (5.097e-05, 8.349e-05)
var thCorrection = []float64{1.789, -0.0165, -0.007536, 9.921e-06, 6.723e-05}

// fetch from mq135 data sheet: {Rs/Ro}=a{ppm}^b
var gasParas = map[string][]float64{
	"CO":  {4.8972, -0.2392},
	"Eth": {3.9593, -0.3031},
	"NH3": {6.5320, -0.3982},
	"Ace": {3.0679, -0.3131},
	"Tol": {3.4880, -0.3203},
	"CO2": {5.3892, -0.3535},
}

const readSampleInterval = 50
const readSampleTimes = 5

type MQ135 struct {
	Mcp         *MCP3008 // MCP3008
	Vcc         float64  // for mq135 voltage
	Gas         string   // for mq135: CO2、NH3、CO、EtOH、Tol、Ace
	Ro          float64  // for mq135
	Rl          float64  // for mq135
	Temperature float64  // for mq135
	Humidity    float64  // for mq135
}

func NewMQ135(mcp *MCP3008, vcc float64, gas string, ro float64, rl float64) *MQ135 {
	mq135 := &MQ135{Mcp: mcp, Vcc: vcc, Gas: gas, Ro: ro, Rl: rl}
	return mq135
}

func (mq *MQ135) MeasureRo(temperature float64, humidity float64, currentGasConcentration float64) float64 {
	val := mq.MeasureResistance(temperature, humidity)

	//R0 = RS * (1 / A * c)-1/B
	mq.Ro = val *
		math.Pow(currentGasConcentration/
			math.Pow(1.0/gasParas[mq.Gas][0], 1.0/gasParas[mq.Gas][1]),
			-1.0/(1.0/gasParas[mq.Gas][1]))

	return mq.Ro
}

func (mq *MQ135) MeasureGasConcentration(temperature float64, humidity float64) (float64, error) {

	rs := 0.0
	for i := 0; i < readSampleTimes; i++ {
		rs += mq.MeasureResistance(temperature, humidity)
		time.Sleep(time.Millisecond * readSampleInterval)
	}
	rs = rs / float64(readSampleTimes)

	if mq.Ro == 0 {
		return 0, errors.New("ro not set")
	}

	ratio := rs / mq.Ro
	// c =  A * (RS / R0)B
	return math.Pow(1.0/gasParas[mq.Gas][0], 1.0/gasParas[mq.Gas][1]) *
		math.Pow(ratio, 1.0/gasParas[mq.Gas][1]), nil
}

func (mq *MQ135) MeasureResistance(temperature float64, humidity float64) float64 {
	var voltage float64 = mq.MeasureVoltage()
	return (mq.Vcc/voltage - 1) * mq.Rl / mq.correctResistance(temperature, humidity)
}

func (mq *MQ135) correctResistance(temperature float64, humidity float64) float64 {
	p00 := thCorrection[0]
	p10 := thCorrection[1]
	p01 := thCorrection[2]
	p20 := thCorrection[3]
	p11 := thCorrection[4]
	return p00 + p10*temperature + p01*humidity + p20*math.Pow(temperature, 2) + p11*temperature*humidity
}

func (mq *MQ135) MeasureVoltage() float64 {
	return mq.Mcp.ReadAdc()
}
