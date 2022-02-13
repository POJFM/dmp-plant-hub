package drivers

import (
	"errors"

	"github.com/stianeikeland/go-rpio/v4"
)

var rpiOpenState bool = false

func OpenRPi() error {
	if !rpiOpenState {
		err := rpio.Open()
		if err != nil {
			return errors.New("open raspberry pi error: " + err.Error())
		}
	}

	return nil
}

func CloseRPi() error {
	if rpiOpenState {
		err := rpio.Close()
		if err != nil {
			return errors.New("close raspberry pi error: " + err.Error())
		}
	}

	return nil
}
