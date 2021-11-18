# https://www.electroniclinic.com/soil-moisture-sensor-with-raspberry-pi-circuit-diagram-and-python-code/
#!/usr/bin/python3
# soil moisture .py

import time
import mcp3008

with mcp3008.MCP3008() as adc:
	#print((adc.read([mcp3008.CH0])[0] - 260) / 260 * 100)
	print(adc.read([mcp3008.CH0])[0])
