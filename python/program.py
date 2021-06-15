#https://www.hackster.io/shafin-kothia/water-level-monitor-with-raspberry-pi-d509a2
import RPi.GPIO as GPIO
import spidev
import Adafruit_DHT
import time
import array

GPIO.setmode(GPIO.BCM)

# Pins
TRIG = 2
ECHO = 3
MOIST = 22
DHT = 23
PUMP = 18
LED = 27

# Web
# pump flow => 0.6 l/min
manualWaterOverdrawn, manualWaterLevel, manualTemp, manualHum
initializationState=False # btn on web after init is completed

# Code
initialization=False
# add reserve => not from bottom but from low water level
waterLevel=0 # on init measures 5 times, appends the values into an array and then averages the values into single value
moistureLevel=0
WaterOverdrawnLevel=0

# Functions
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
        humidity, temperature = Adafruit_DHT.read(Adafruit_DHT.DHT11, DHT)
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
    percentage value = ((value / max) * 100)
    return percentage

def arithmeticMean(list):
    return sum(list) / len(list)

# IO
GPIO.setup(TRIG ,GPIO.OUT)
GPIO.setup(ECHO,GPIO.IN)
GPIO.setup(MOIST ,GPIO.IN)
GPIO.setup(DHT ,GPIO.IN)
GPIO.setup(PUMP ,GPIO.OUT)
GPIO.setup(LED ,GPIO.OUT)

GPIO.output(TRIG, False)

# Initialization sequence
print("Starting initialization sequence...")
time.sleep(2)

# init measurement
waterLevelAvg= []
moistureAvg= []

while count < 5:
  waterLevelAvg.append(waterLevelMeasure())
  moistureAvg.append(moistureMeasure())
  count += 1
  time.sleep(1)

moistureLevel = arithmeticMean(moistureAvg)



# wait for initializationState from web then get values




if initializationState = True:
  if bool(manualWaterLevel):
    waterLevel = manualWaterLevel
  else:
    waterLevel = arithmeticMean(waterLevelAvg) - 2

  if bool(manualWaterOverdrawn):
    WaterOverdrawnLevel = manualWaterOverdrawn

  print(f"Water level is set to: {round(waterLevel, 2)}")
  print(f"Measured moisture level: {round(moistureLevel, 2)}")
  time.sleep(3)
  initialization = True
  print("GardenBot is coming to life...🥬🤖...")

# Measurement sequence
while initialization:
  if waterLevelMeasure() < moistureLevel:
    print("Soil is drying out, starting irrigation...")
    # time passed from running pump will be represented as liters
    tic = time.perf_counter()
    while waterLevelMeasure() < moistureLevel or flowMeasure < WaterOverdrawnLevel:
      toc = time.perf_counter()
      GPIO.output(PUMP, True);
    

    # after pump stops run checking sequence


    # Checking sequence
    if waterLevelMeasure() < waterLevel:
      print("!Water tank limit level reached!")
      while waterLevelMeasure() < waterLevel:
        GPIO.output(LED, True);
        time.sleep(0.5)
        GPIO.output(LED, False);
        time.sleep(0.5)
    else:
      GPIO.output(LED, False);

  else:
    GPIO.output(PUMP, False);

