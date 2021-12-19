import { useEffect, useState, useRef } from 'react'
import axios from 'axios'
import Card from '@material-ui/core/Card'
import CardContent from '@material-ui/core/CardContent'
//import { RefreshButton } from './buttons/RefreshButton'
import { useQuery } from '@apollo/client'
import { dashboard } from '../graphql/queries'
import {
	Chart as ChartJS,
	CategoryScale,
	LinearScale,
	PointElement,
	LineElement,
	BarElement,
	Title,
	Tooltip,
	Legend,
} from 'chart.js'
import { LiveMeasurementsChart, WaterConsumptionChart, IrrigationChart, MeasurementsHistoryChart } from './Chart'
import ReactScrollWheelHandler from 'react-scroll-wheel-handler'
import { getMonths } from 'src/utils'

export default function Dashboard() {
	const [currentTemp, setCurrentTemp] = useState(0),
		[currentMoist, setCurrentMoist] = useState(0),
		[currentHum, setCurrentHum] = useState(0),
		[temp, setTemp] = useState<any>(),
		[moist, setMoist] = useState<any>(),
		[hum, setHum] = useState<any>(),
		[irrigationCount, setIrrigationCount] = useState<number>(),
		[months, setMonths] = useState<any>(),
		[weather, setWeather] = useState<any>()

	const { loading, error, data } = useQuery(dashboard)
	// TEST
	const settings = { chartType: 0 }
	// END TEST

	useEffect(() => {
		document.title = 'Plant Hub | Dashboard'
		
		const currentMonth = new Date().getMonth()
		console.log(currentMonth)
		setMonths(getMonths())
		console.log(months)

		for (let i = currentMonth; i < 12; i++) {
			let month: number
			// data.IrrigationHistory.map((item: any) => {
			// 	item.timestamp === 'regex na jeden měsíc a číslo z loopu podle current month' && month++
			// })
			//irrigationCount.push(month)
			console.log(i)
			i > 10 && (i = 0)
			i === currentMonth - 1 && (i = 12)
		}
	}, [])

	//https://graphql.org/blog/rest-api-graphql-wrapper/
	const fetchWeatherForecast = () => {
		axios
			.request({
				method: 'GET',
				url: `https://api.openweathermap.org/data/2.5/onecall?lat=${49.68333}&lon=${18.35}&exclude=daily,minutely,alerts&units=metric&appid=${
					process.env.REACT_APP_FORECAST_API_KEY
				}`,
				headers: {
					'Content-Type': 'application/json',
				},
			})
			.then((res) => {
				setWeather(res.data.hourly.slice(0, 15)) //
			})
			.catch((error) => {
				console.error(error)
			})
	}

	// Fetch on render then every 30mins
	setTimeout(() => fetchWeatherForecast(), weather ? 300_000 : 1)

	console.log(
		`https://api.openweathermap.org/data/2.5/onecall?lat=${49.68333}&lon=${18.35}&exclude=daily,minutely,alerts&units=metric&appid=${
			process.env.REACT_APP_FORECAST_API_KEY
		}`
	)

	const liveMeasurements = () => {
		axios
			.request({
				method: 'GET',
				url: `${process.env.REACT_APP_GO_API_URL}/live/measure`,
				headers: {
					'Content-Type': 'application/json',
				},
			})
			.then((res) => {
				setCurrentTemp(res.data.temp)
				setCurrentHum(res.data.hum)
				setCurrentMoist(res.data.moist)
				temp.push(res.data.temp)
				hum.push(res.data.hum)
				moist.push(res.data.moist)
			})
			.catch((error) => {
				console.error(error)
			})
	}

	setTimeout(() => liveMeasurements(), 1000)

	const useHorizontalScroll = () => {
		const elRef = useRef()
		useEffect(() => {
			const el = elRef.current
			if (el) {
				const onWheel = (e: any) => {
					if (e.deltaY == 0) return
					e.preventDefault()
					//@ts-ignore
					el.scrollTo({
						//@ts-ignore
						left: el.scrollLeft + e.deltaY,
						behavior: 'smooth',
					})
				}
				//@ts-ignore
				el.addEventListener('wheel', onWheel)
				//@ts-ignore
				return () => el.removeEventListener('wheel', onWheel)
			}
		}, [])
		return elRef
	}

	ChartJS.register(CategoryScale, LinearScale, PointElement, LineElement, BarElement, Title, Tooltip, Legend)

	// // Reduce the animation steps for demo clarity.
	// var myLiveChart = new Chart(ctx).Line(startingData, {animationSteps: 15});

	// setInterval(function(){
	// 	// Get a random index point
	// 	var indexToUpdate = Math.round(Math.random() * startingData.labels.length);

	// 	// Update one of the points in the second dataset
	// 	myLiveChart.datasets[1].points[indexToUpdate].value = Math.random() * 100;

	// 	myLiveChart.update();
	// }, 5000);

	const weatherForecastDataToIcons = (weatherState: any) => {
		if (weatherState.icon.includes('01d')) {
			return 'clearsky'
		} else if (weatherState.icon.includes('01n')) {
			return 'clearskyNight'
		} else if (weatherState.icon.includes('02d')) {
			return 'fair'
		} else if (weatherState.icon.includes('02n')) {
			return 'fairNight'
		} else if (weatherState.description.includes('snow')) {
			return 'snow'
		} else if (weatherState.description.includes('mist')) {
			return 'mist'
		} else if (weatherState.description.includes('rain')) {
			return 'rain'
		} else if (weatherState.description.includes('sleet') || weatherState.description.includes('Sleet')) {
			return 'sleet'
		} else if (weatherState.description.includes('scattered clouds')) {
			return 'partlyCloudy'
		} else if (
			weatherState.description.includes('broken clouds') ||
			weatherState.description.includes('overcast clouds')
		) {
			return 'cloudy'
		} else if (weatherState.description.includes('thunderstorm')) {
			return 'rainAndThunder'
		}
	}

	return (
		<div className="dashboard">
			<Card className="card h-52">
				<CardContent>
					<div className="flex-row">
						<div className="flex-col w-4/12">
							<div className="flex-row">
								<div className="flex-col ml-5">
									<div className="flex-row pt-5px" title="Teplota vzduchu">
										<span className="flex-col w-12 max-h-full">
											<img src="/assets/icons/dashboard/temperature.svg" />
										</span>
										<span className="flex-col">{`${currentTemp}°C`}</span>
									</div>
									<div className="flex-row pt-5px" title="Vlhkost vzduchu">
										<span className="flex-col w-12 max-h-full">
											<img src="/assets/icons/dashboard/humidity.svg" />
										</span>
										<span className="flex-col">{`${currentHum}%`}</span>
									</div>
									<div className="flex-row pt-5px" title="Vlhkost půdy">
										<span className="flex-col w-12 max-h-full">
											<img src="/assets/icons/dashboard/moisture.svg" />
										</span>
										<span className="flex-col">{`${currentMoist}%`}</span>
									</div>
								</div>
								<div className="flex-col ml-5">
									<div className="flex-row pt-5px" title="Výška vody v nádrži">
										<span className="flex-col w-12 max-h-full">
											<img src="/assets/icons/dashboard/waterLevel.svg" />
										</span>
										<span className="flex-col">{/* {`${data.irrigationHistory.waterLevel}cm`} */}0cm</span>
									</div>
									<div className="flex-row pt-5px" title="Objem vody v nádrži">
										<span className="flex-col w-12 max-h-full">
											<img src="/assets/icons/dashboard/waterOverdrawn.svg" />
										</span>
										<span className="flex-col">{/* {`${data.irrigationHistory.waterAmount}l`} */}5l</span>
									</div>
									<div className="flex-row pt-5px" title="Celkový vyčerpaný objem vody">
										<span className="flex-col w-12 max-h-full">
											<img src="/assets/icons/dashboard/waterOverdrawn.svg" />
										</span>
										<span className="flex-col">{/* {`${data.irrigationHistory.waterAmount}l`} */}56l</span>
									</div>
								</div>
							</div>
						</div>
						<div className="flex-col w-8/12">
							<div className="flex-row h-44 -mt-2">
								<LiveMeasurementsChart
									chartType={settings.chartType}
									temp={temp}
									hum={hum}
									moist={moist} /* data.settings.chartType */
								/>
							</div>
						</div>
					</div>
				</CardContent>
			</Card>
			<div className="flex-row">
				<div className="flex-col w-6/12">
					<div className="flex-row">
						<div className="flex-col w-full">
							<Card className="card-left">
								<CardContent>
									<ReactScrollWheelHandler
										upHandler={(e) => console.log('scroll up')}
										downHandler={(e) => console.log('scroll down')}
										leftHandler={(e) => console.log('scroll left')}
										rightHandler={(e) => console.log('scroll right')}
									>
										<div /* ref={() => useHorizontalScroll} */ className="flex-row overflow-x-scroll">
											{weather?.map((weatherItem: any, i: number) => {
												const date = new Date()
												let time = date.getHours() + i
												time > 23 && (time = i)

												return (
													<div className="flex-col mx-2 mb-2 weatherWrapper">
														<div className="flex-row mx-2 w-14 flex-center">
															<div className="flex-col">
																<span className="text-2xl">{`${time}:00`}</span>
															</div>
														</div>
														<div className="flex-col">
															<div className="flex-row text-xs">
																<img src="/assets/icons/dashboard/temperature.svg" className="w-8 max-h-full" />
																<span> {`${weatherItem.temp}°C`}</span>
															</div>
															<div className="flex-row text-xs">
																<img src="/assets/icons/dashboard/humidity.svg" className="w-8 max-h-full" />
																<span> {`${weatherItem.humidity}%`}</span>
															</div>
														</div>
														<div className="flex-row">
															<div className="flex-col flex-center">
																<img
																	src={`/assets/icons/weatherForecast/${weatherForecastDataToIcons(
																		weatherItem.weather[0]
																	)}.svg`}
																	className="dashboard w-14 h-full"
																/>
															</div>
														</div>
													</div>
												)
											})}
										</div>
									</ReactScrollWheelHandler>
								</CardContent>
							</Card>
						</div>
					</div>
					<div className="flex-row">
						<div className="flex-col w-full">
							<Card className="card-left">
								<CardContent>
									<div className="flex-row ">
										<span>Historie zavlažování</span>
									</div>
									<div className="flex-row 2xl:h-96 lg:h-52">
										<IrrigationChart
											chartType={settings.chartType}
											moist={moist}
											hum={hum}
											temp={temp}
											irrigationCount={irrigationCount}
										/>
									</div>
								</CardContent>
							</Card>
						</div>
					</div>
				</div>
				<div className="flex-col w-6/12">
					<Card className="card-right h-full">
						<CardContent>
							<div className="flex-row">
								<span>Spotřeba vody</span>
							</div>
							<div className="flex-row 2xl:h-64 lg:h-48">
								<WaterConsumptionChart
									chartType={settings.chartType}
									waterOverdrawn={5 /* data.irrigationHistory.waterOverdrawn */}
									irrigationCount={irrigationCount}
								/>
							</div>
							<div className="flex-row mt-3">
								<span>Historie měření</span>
							</div>
							<div className="flex-row 2xl:h-80 lg:h-52">
								<MeasurementsHistoryChart chartType={settings.chartType} moist={moist} hum={hum} temp={temp} />
							</div>
						</CardContent>
					</Card>
				</div>
			</div>
		</div>
	)
}
