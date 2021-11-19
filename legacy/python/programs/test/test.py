import RPi.GPIO as GPIO
import time

GPIO.setmode(GPIO.BCM)

LED = 27
GPIO.setup(LED ,GPIO.OUT)

try:
  while(True):
    GPIO.output(LED, True)
    print("blink") 
    time.sleep(1)
    GPIO.output(LED, False)
    time.sleep(1)        
except KeyboardInterrupt: # If CTRL+C is pressed, exit cleanly:
  print("Keyboard interrupt")

except:
  print("some error") 

finally:
  print("clean up") 
  GPIO.cleanup() # cleanup all GPIO 