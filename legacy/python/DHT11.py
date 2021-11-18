''' # https://peppe8o.com/using-raspberry-pi-with-dht11-temperature-and-humidity-sensor-and-python/
import time
import board
import adafruit_dht
#Initial the dht device, with data pin connected to:
dhtDevice = adafruit_dht.DHT11(board.D17)
while True:
    try:
         # Print the values to the serial port
         temperature_c = dhtDevice.temperature
         temperature_f = temperature_c * (9 / 5) + 32
         humidity = dhtDevice.humidity
               .format(temperature_f, temperature_c, humidity))
    except RuntimeError as error:     # Errors happen fairly often, DHT's are hard to read, just keep going
         print(error.args[0])
  time.sleep(2.0) 

# https://www.thegeekpub.com/236867/using-the-dht11-temperature-sensor-with-the-raspberry-pi/
import Adafruit_DHT
import time

DHT_SENSOR = Adafruit_DHT.DHT11
DHT_PIN = 4

while True:
  humidity, temperature = Adafruit_DHT.read(DHT_SENSOR, DHT_PIN)
  if humidity is not None and temperature is not None:
    print("Temp={0:0.1f}C Humidity={1:0.1f}%".format(temperature, humidity))
  else:
    print("Sensor failure. Check wiring.");
  time.sleep(3);

# used code
import Adafruit_DHT
import time

DHT = 23

while True:
  try:
    humidity, temperature = Adafruit_DHT.read(Adafruit_DHT.DHT11, DHT)
    if humidity is not None and temperature is not None:
      print("Temp={0:0.1f}C Humidity={1:0.1f}%".format(temperature, humidity))
  except RuntimeError as error:     # Errors happen fairly often, DHT's are hard to read, just keep going
    print(error.args[0])
  time.sleep(3);

#!/usr/bin/python

# Copyright (c) 2014 Adafruit Industries
# Author: Tony DiCola

# Permission is hereby granted, free of charge, to any person obtaining a copy
# of this software and associated documentation files (the "Software"), to deal
# in the Software without restriction, including without limitation the rights
# to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
# copies of the Software, and to permit persons to whom the Software is
# furnished to do so, subject to the following conditions:

# The above copyright notice and this permission notice shall be included in all
# copies or substantial portions of the Software.

# THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
# IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
# FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
# AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
# LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
# OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
# SOFTWARE.
import Adafruit_DHT

# Sensor should be set to Adafruit_DHT.DHT11,
# Adafruit_DHT.DHT22, or Adafruit_DHT.AM2302.
sensor = Adafruit_DHT.DHT22

# Example using a Beaglebone Black with DHT sensor
# connected to pin P8_11.
pin = 'P8_11'

# Example using a Raspberry Pi with DHT sensor
# connected to GPIO23.
#pin = 23

# Try to grab a sensor reading.  Use the read_retry method which will retry up
# to 15 times to get a sensor reading (waiting 2 seconds between each retry).
humidity, temperature = Adafruit_DHT.read_retry(sensor, pin)

# Note that sometimes you won't get a reading and
# the results will be null (because Linux can't
# guarantee the timing of calls to read the sensor).
# If this happens try again!
if humidity is not None and temperature is not None:
    print('Temp={0:0.1f}*C  Humidity={1:0.1f}%'.format(temperature, humidity))
else:
    print('Failed to get reading. Try again!')'''
    
#!/usr/bin/python
import Adafruit_DHT
from termcolor import colored

DHT = 23

def DHTMeasure():
  try:
    humidity, temperature = Adafruit_DHT.read_retry(Adafruit_DHT.DHT11, DHT)
    if humidity is not None and temperature is not None:
      return [temperature, humidity]
    else:
      return None
  except RuntimeError as error:     # Errors happen fairly often, DHT's are hard to read, just keep going
    return error.args[0]

while True:
	DHTMeasureValues = DHTMeasure()
	print(colored(f'Temperature: {DHTMeasureValues[0]}˚C', 'blue') + colored(f' Humidity: {DHTMeasureValues[1]}%', 'blue'))