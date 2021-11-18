''' # https://peppe8o.com/using-raspberry-pi-with-dht11-temperature-and-humidity-sensor-and-python/
import time
import board
import adafruit_dht
# Initial the dht device, with data pin connected to:
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
# pin = 23

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
import time
from flask import Flask, render_template, request
app = Flask(__name__)

DHT = 23

def exception_handler(func):
    def wrapper(*args, **kwargs):
        try:
            return func(*args, **kwargs)
        except Exception as e:
            error_code = getattr(e, "code", 500)
            logger.exception("Service exception: %s", e)
            r = dict_to_json(
              {"message": e.message, "matches": e.message, "error_code": error_code})
            return Response(r, status=error_code, mimetype='application/json')
    wrapper.__name__ = func.__name__
    return wrapper


''' app.add_url_rule('/time',
    view_func=Main.as_view('time'),
    methods=['GET'])

app.add_url_rule('/init/measured',
    view_func=Main.as_view('init_measured'),
    methods=['GET'])

app.add_url_rule('/measure',
    view_func=Main.as_view('measure'),
    methods=['GET']) '''


@app.route('/time')
@exception_handler
def get_current_time():
    return {'time': time.time()}

# send limit values to web


@app.route('/init/measured', methods=['GET'])
@exception_handler
def api_init_measure():
    return {
        'waterLevel': 44444,
        'moistureLevel': 55555,
    }


def DHTMeasure():
    try:
        humidity, temperature = Adafruit_DHT.read_retry(
            Adafruit_DHT.DHT11, DHT)
        if humidity is not None and temperature is not None:
            return [temperature, humidity]
        else:
            return None
    except RuntimeError as error:     # Errors happen fairly often, DHT's are hard to read, just keep going
        return error.args[0]


# DHTMeasureValues = DHTMeasure()
# @app.route('/measure')
# @exception_handler
# def api_measure_data():
#   return {
#     'temperature': DHTMeasureValues[0],
#     'humidity': DHTMeasureValues[1],
#   }
DHTMeasureValues = DHTMeasure()
@app.route('/measure')
@exception_handler
def api_measure_data():
	while True:
		return {
			'temperature': DHTMeasureValues[0],
			'humidity': DHTMeasureValues[1],
		}
''' while True:
  DHTMeasureValues = DHTMeasure()
  print(colored(f'Temperature: {DHTMeasureValues[0]}˚C', 'blue') + colored(f' Humidity: {DHTMeasureValues[1]}%', 'blue'))
 '''
if __name__ == '__main__':
    app.run(host='0.0.0.0')
