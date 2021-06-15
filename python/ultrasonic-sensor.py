#https://www.hackster.io/shafin-kothia/water-level-monitor-with-raspberry-pi-d509a2
import RPi.GPIO as GPIO
import time


GPIO.setmode(GPIO.BCM)

TRIG = 2
ECHO = 3
LED = 27
i=0

GPIO.setup(TRIG ,GPIO.OUT)
GPIO.setup(ECHO,GPIO.IN)
GPIO.setup(LED ,GPIO.OUT)

GPIO.output(TRIG, False)
print("Starting.....")
time.sleep(2)

while True:
   GPIO.output(TRIG, True)
   time.sleep(0.00001)
   GPIO.output(TRIG, False)

   while GPIO.input(ECHO)==0:
      pulse_start = time.time()

   while GPIO.input(ECHO)==1:
      pulse_stop = time.time()

   pulse_time = pulse_stop - pulse_start

   distance = pulse_time * 17150
   print(round(distance, 2));

   time.sleep(1)
   
   if distance < 4:
       print("Water will overflow")
       GPIO.output(LED, True);
       time.sleep(0.5)
       GPIO.output(LED, False);
       time.sleep(0.5)
       GPIO.output(LED, True);
       time.sleep(0.5)
       GPIO.output(LED, False);
       time.sleep(0.5)
   else:
       GPIO.output(LED, False);