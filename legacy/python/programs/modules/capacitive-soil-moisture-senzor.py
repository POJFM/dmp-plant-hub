# https://www.electroniclinic.com/soil-moisture-sensor-with-raspberry-pi-circuit-diagram-and-python-code/
#! / usr / bin / python3
# soil moisture .py
import spidev
import time
max = 460.0 # Maximum value at full humidity
spi = spidev.SpiDev()
spi.open(0, 1)
answer = spi.xfer([1, 128, 0])
if 0 <= answer[1] <= 3:
  value = 1023 - ((answer[1] * 256) + answer[2])
percentage = ((value / max) * 100)
print("Soil moisture:", percentage, "%")