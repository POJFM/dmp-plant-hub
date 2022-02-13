package drivers

import (
	"errors"

	"github.com/stianeikeland/go-rpio/v4"
)

type MCP3008 struct {
	Chip      uint8   // raspberry pi spi chip: CE0 CE1
	CH        uint8   // channel: CH0-CH7
	VRef      float64 // voltage reference
	closeFunc func()
	endFunc   func()
}

func NewMCP3008(chip uint8, ch uint8, vref float64) (*MCP3008, error) {
	mcp := &MCP3008{Chip: chip, CH: ch, VRef: vref}

	if err := OpenRPi(); err != nil {
		return nil, errors.New("init error: " + err.Error())
	}

	if err := rpio.SpiBegin(rpio.Spi0); err != nil {
		return nil, errors.New("init error: " + err.Error())
	}

	rpio.SpiSpeed(1350000)
	rpio.SpiMode(0, 0)
	rpio.SpiChipSelect(mcp.Chip) // Select CE0 slave

	mcp.endFunc = func() {
		rpio.SpiEnd(rpio.Spi0)
	}

	mcp.closeFunc = func() {
		CloseRPi()
	}

	return mcp, nil
}

func (mcp *MCP3008) End() {
	mcp.endFunc()
}

func (mcp *MCP3008) Close() {
	mcp.closeFunc()
}

func (mcp *MCP3008) ReadAdc() float64 {
	cmd := make([]byte, 3)
	// command: 0x01，transmit start
	cmd[0] = 0x1
	// command:（0x08|CH0~CH7）<<4
	cmd[1] = (0x08 | (mcp.CH & 0x07)) << 4
	// command: 0x00
	cmd[2] = 0x00
	rpio.SpiExchange(cmd)
	// take the last two bytes, a total of 10 bits, and convert to uint16
	var adc = uint16(cmd[1]&0x03)<<8 | uint16(cmd[2])
	//fmt.Println(strconv.FormatUint(uint64(adc), 2))
	return (float64(adc) * mcp.VRef) / 1024.0
}
