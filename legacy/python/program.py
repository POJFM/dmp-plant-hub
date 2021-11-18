#https://www.hackster.io/shafin-kothia/water-level-monitor-with-raspberry-pi-d509a2
import RPi.GPIO as GPIO
import spidev
import Adafruit_DHT
import time
import array
from termcolor import colored
import flask
from flask import request, jsonify
import logging

GPIO.setwarnings(False)
GPIO.setmode(GPIO.BCM)

app = flask.Flask(__name__)
app.config["DEBUG"] = True
log = logging.getLogger('werkzeug')
log.setLevel(logging.ERROR)
log.disabled = True
app.logger.disabled = True
@app.route('/', methods=['GET'])
def home():
    return '''<h1>Welcome to Plant Hub</h1>'''

@app.route('/api/books', methods=['GET'])
def api_all():
    return jsonify(books)

app.run()

# Pins
TRIG = 2
ECHO = 3
MOIST = 22
DHT = 23
PUMP = 18
LED = 27

# Web
# Objemový limit
manualWaterOverdrawn=0
# Hladinový limit
manualWaterLevel=0
manualTemp=0
manualHum=0
initializationState=False # btn on web after init is completed

# Code
initialization=False
# add reserve => not from bottom but from low water level
waterLevel=0 # on init measures 5 times, appends the values into an array and then averages the values into single value
moistureLevel=0
waterOverdrawnLevel=0
pumpFlow=0.6 # litr/min

# Sensors constants
def waterLevelMeasure():
  GPIO.output(TRIG, True)
  time.sleep(0.00001)
  GPIO.output(TRIG, False)

  while GPIO.input(ECHO)==0:
    pulse_start = time.time()

  while GPIO.input(ECHO)==1:
    pulse_stop = time.time()

  pulse_time = pulse_stop - pulse_start
  distance = pulse_time * 17150
  return distance

def DHTMeasure():
  try:
    humidity, temperature = Adafruit_DHT.read_retry(Adafruit_DHT.DHT11, DHT)
    if humidity is not None and temperature is not None:
      print("Temp={0:0.1f}C Humidity={1:0.1f}%".format(temperature, humidity))
      result = [temperature, humidity]
      return result
    else:
      return None
  except RuntimeError as error:     # Errors happen fairly often, DHT's are hard to read, just keep going
    return error.args[0]

def moistureMeasure():
  max = 460.0 # Maximum value at full humidity
  spi = spidev. SpiDev ()
  spi. open (0, 1)
  answer = spi. xfer ([1, 128, 0])
  if 0 <= answer [1] <= 3:
    value = 1023 - ((answer [1] * 256) + answer [2])

  value = ((value / max) * 100)
  return value

def arithmeticMean(list):
  return sum(list) / len(list)

def timeToOverdraw():
  return manualWaterOverdrawn / pumpFlow

# IO
GPIO.setup(TRIG ,GPIO.OUT)
GPIO.setup(ECHO,GPIO.IN)
GPIO.setup(MOIST ,GPIO.IN)
GPIO.setup(DHT ,GPIO.IN)
GPIO.setup(PUMP ,GPIO.OUT)
GPIO.setup(LED ,GPIO.OUT)

GPIO.output(TRIG, False)

# Initialization sequence
print(colored('Starting initialization sequence...🏁🤖🏁', 'green'))
time.sleep(2)

# init measurement
waterLevelAvg = []
moistureAvg = []

try:
  # calculating average value
  count = 0
  while count < 5:
    waterLevelAvg.append(waterLevelMeasure())
    moistureAvg.append(moistureMeasure())
    count += 1
    time.sleep(1)

  moistureLevel = arithmeticMean(moistureAvg)

  # send limit values to web

  # wait for initializationState from web then get values
  printWait=True
  while initializationState:
    if printWait:
      print(colored('Waiting for initialization sequence to finish...📝🤖📝', 'green'))
      printWait = False
    time.sleep(1)

  if bool(manualWaterLevel):
    waterLevel = manualWaterLevel
  else:
    waterLevel = arithmeticMean(waterLevelAvg) - 2

  if bool(manualWaterOverdrawn):
    waterOverdrawnLevel = manualWaterOverdrawn

  print(colored(f'Water level is set to: {round(waterLevel, 2)}cm', 'blue'))
  print(colored(f'Water overdrawn level is set to: {round(waterOverdrawnLevel, 2)}l', 'blue'))
  print(colored(f'Measured moisture level: {round(moistureLevel, 2)}%', 'blue'))
  time.sleep(3)
  initialization = True
  print(colored('GardenBot is coming to life...✅🤖✅', 'green'))
  time.sleep(1)

  # Measurement sequence

  while initialization:
    if moistureMeasure() < moistureLevel:
      print(colored('Soil is drying out, starting irrigation...🌿🤖🚿', 'green'))
      
      # time passed from running pump will be represented as liters
      flowMeasure=0
      t0 = time.time()
      while waterLevelMeasure() < moistureLevel or flowMeasure < timeToOverdraw():
        t1 = time.time()
        GPIO.output(PUMP, True);
        flowMeasure = t1 - t0

      # after pump stops run Checking sequence
      if waterLevelMeasure() < waterLevel:
        print(colored('!Water tank limit level reached!...🚫🤖🚫', 'red'))
        # send notification to web
        while waterLevelMeasure() < waterLevel:
          GPIO.output(LED, True);
          time.sleep(0.5)
          GPIO.output(LED, False);
          time.sleep(0.5)
      else:
        GPIO.output(LED, False);

    else:
      GPIO.output(PUMP, False);

    DHTMeasureValues = DHTMeasure()
    print(colored(f'Temperature: {DHTMeasureValues[0]}˚C', 'blue'))
    print(colored(f'Humidity: {DHTMeasureValues[1]}%', 'blue'))
    print(colored(f'Soil moisture: {moistureMeasure()}%', 'blue'))
    time.sleep(1)
    
except KeyboardInterrupt: # If CTRL+C is pressed, exit cleanly:
  print("Keyboard interrupt")

finally:
  print("clean up") 
  GPIO.cleanup() # cleanup all GPIO 