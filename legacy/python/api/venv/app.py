""" 
Plant Hub API
"""
import os
import time
import json
import requests
import geocoder
from flask import Flask, render_template, request
from os.path import join, dirname
from requests import get
from bs4 import BeautifulSoup
from dotenv import dotenv_values, load_dotenv

app = Flask(__name__)

load_dotenv('.env')
WEATHER_API_KEY = os.environ.get("WEATHER_API_KEY")

# Get public IP and return geo dataset
geo = geocoder.ip(get('https://api.ipify.org').text)


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


''' weatherData = BeautifulSoup(requests.get('https://www.google.com/search?q=weather+{}'.format(geo.city), cookies={'privacy-policy': '1,XXXXXXXXXXXXXXXXXXXXXX', 'CONSENT':'YES+cb.20210418-17-p0.it+FX+917; '}).content, 'html.parser')
#weatherData = BeautifulSoup(requests.get('https://www.google.com/search?q=weather+ostrava', cookies={'privacy-policy': '1,XXXXXXXXXXXXXXXXXXXXXX', 'CONSENT':'YES+cb.20210418-17-p0.it+FX+917; '}).content, 'html.parser')


# get the temperature
weatherTemperature = weatherData.find('div', attrs={'class': 'BNeawe iBp4i AP7Wnd'}).text

# this contains time and sky description
weatherTimeSkyData = weatherData.find('div', attrs={'class': 'BNeawe tAd8D AP7Wnd'}).text

weatherTimeSky = weatherTimeSkyData.split('\n')
weatherTime = weatherTimeSky[0]
weatherSky = weatherTimeSky[1]
 '''

''' print(weatherTemperature)
print(weatherTime)
print(weatherSky) '''

weatherDataRaw = requests.get(
    'http://api.openweathermap.org/data/2.5/forecast?appid={}&units=metric&cnt=5&q={}'.format(WEATHER_API_KEY, geo.city))
weatherData = json.loads(weatherDataRaw.text)
# print(weatherData)
# print(weatherData['list'])
# for list in weatherData['list']:
#     print('{}\n'.format(list['main']['temp']))
# print(weatherData.list[1].main.temp)

# time, temp,


@app.route('/weather', methods=['GET'])
@exception_handler
def api_weather():
    return {
        'waterLevel': 44444,
        'moistureLevel': 55555,
    }

# send limit values to web


@app.route('/init/measured', methods=['GET'])
@exception_handler
def api_init_measured():
    return {
        'moistureLevel': 20,
        'waterLevel': 52,
        'waterOverdrawnLevel': 2,
        'location': geo.city
    }


if __name__ == '__main__':
    app.run(host='0.0.0.0')
