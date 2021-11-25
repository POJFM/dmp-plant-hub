package sensors

import (
	"github.com/SPSOAFM-IT18/dmp-plant-hub/sensors/dht"
	"log"
)

func ReadDHT() (temp float32, hum float32) {
	temp, hum, err := dht.ReadDHTxx(dht.DHT11, 23, false)
	if err != nil {
		log.Fatal(err)
	}
	return temp, hum
}
