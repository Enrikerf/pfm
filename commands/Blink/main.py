import RPi.GPIO as GPIO    # Import Raspberry Pi GPIO library
from time import sleep     # Import the sleep function from the time module
GPIO.setwarnings(False)    # Ignore warning for now
GPIO.setmode(GPIO.BOARD)   # Use physical pin numbering
GPIO.setup(12, GPIO.OUT, initial=GPIO.LOW)  
GPIO.setup(22, GPIO.OUT, initial=GPIO.LOW)  
GPIO.setup(24, GPIO.OUT, initial=GPIO.LOW)  
GPIO.setup(26, GPIO.OUT, initial=GPIO.LOW)  
GPIO.setup(32, GPIO.OUT, initial=GPIO.LOW)  


while True: # Run forever
 GPIO.output(12, GPIO.HIGH) # Turn on
 GPIO.output(22, GPIO.HIGH) # Turn on
 GPIO.output(24, GPIO.HIGH) # Turn on
 GPIO.output(26, GPIO.HIGH) # Turn on
 GPIO.output(32, GPIO.HIGH) # Turn on
 sleep(1) # Sleep for 1 second
 GPIO.output(12, GPIO.LOW) # Turn off
 GPIO.output(22, GPIO.LOW) # Turn off
 GPIO.output(24, GPIO.LOW) # Turn off
 GPIO.output(26, GPIO.LOW) # Turn off
 GPIO.output(32, GPIO.LOW) # Turn off
 sleep(1) # Sleep for 1 second