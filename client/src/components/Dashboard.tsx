import { useEffect, useState, useRef } from 'react'
import axios from 'axios'
import Card from '@material-ui/core/Card'
import CardContent from '@material-ui/core/CardContent'
//import { RefreshButton } from './buttons/RefreshButton'
import { useQuery } from '@apollo/client'
import { plantState, plantStateHistory, irrigationHistory, posts } from '../graphql/queries'
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

export default function Dashboard() {
	const [temperature, setTemperature] = useState(0)
	const [humidity, setHumidity] = useState(0)
	const [weather, setWeather] = useState<any>()
	const [currentLatitude, setCurrentLatitude] = useState('')
	const [currentLongitude, setCurrentLongitude] = useState('')
	const { loading, error, data } = useQuery(posts)
	console.log(data)
	// TEST
	const settings = { chartType: 0 }

	useEffect(() => {
		document.title = 'Plant Hub | Dashboard'
	}, [])

	// fetch('/measure')
	// 	.then((res) => res.json())
	// 	.then((data) => {
	// 		setTemperature(data.temperature)
	// 		setHumidity(data.humidity)
	// 	})
	// 	.then((err) => {
	// 		console.log(err)
	// 	})

	// fetch('/init/measured')
	// 	.then((res) => res.json())
	// 	.then((data) => {
	// 		setCurrentLatitude(data?.latitude)
	// 		setCurrentLongitude(data?.longitude)
	// 	})
	// 	.then((err: any) => {
	// 		console.log(err)
	// 	})

	console.log(`${process.env.REACT_APP_FORECAST_API_URL}?lat=49.68333&lon=18.35`)

	//https://graphql.org/blog/rest-api-graphql-wrapper/
	const fetchWeatherForecast = () => {
		axios
			.request({
				method: 'GET',
				url: `${process.env.REACT_APP_FORECAST_API_URL}?lat=49.68333&lon=18.35`,
				//url: `${process.env.REACT_APP_FORECAST_API_URL}?lat=${currentLatitude}&lon=${currentLongitude}`,
				headers: {
					'Content-Type': 'application/json',
				},
			})
			.then((response) => {
				let currentItem = 0
				const date = new Date()
				var currentTime = `${date.getHours()}`
				currentTime.length === 1 && (currentTime = 0 + currentTime)
				response.data.properties.timeseries.map(
					(timeseries: any, i: number) => String(/..(?=:)/.exec(timeseries.time)) === currentTime && (currentItem = i)
				)
				setWeather(response.data.properties.timeseries.slice(currentItem, 15 + currentItem))
			})
			.catch((error) => {
				console.error(error)
			})
	}

	// Fetch on render then every 30mins
	setTimeout(() => fetchWeatherForecast(), weather ? 300_000 : 1)

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

	/* 	ChartJS.scaleService.updateScaleDefaults('category', {
		gridLines: {
				drawBorder: false,
				drawOnChartArea: false,
				drawTicks: false
		},
		ticks: {
				// autoSkip: false,
				padding: 20
		},
		maxBarThickness: 10
}); */

	const weatherForecastBullshitToIcons = (weatherState: string) => {
		if (weatherState?.includes('rainandthunder')) {
			return 'rainAndThunder'
		} else if (weatherState?.includes('snowandthunder')) {
			return 'snowAndThunder'
		} else if (weatherState?.includes('partlycloudy')) {
			if (weatherState?.includes('night')) {
				return 'partlyCloudyNight'
			} else {
				return 'partlyCloudy'
			}
		} else if (weatherState?.includes('clearsky')) {
			if (weatherState?.includes('night')) {
				return 'clearskyNight'
			} else {
				return 'clearsky'
			}
		} else if (weatherState?.includes('fair')) {
			return 'fair'
		} else if (weatherState?.includes('snow')) {
			return 'snow'
		} else if (weatherState?.includes('fog')) {
			return 'fog'
		} else if (weatherState?.includes('rain')) {
			return 'rain'
		} else if (weatherState?.includes('sleet')) {
			return 'sleet'
		} else if (weatherState?.includes('cloudy')) {
			return 'cloudy'
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
									<div className="flex-row pt-5px">
										<span className="flex-col w-12 max-h-full">
											<img src="/assets/icons/dashboard/temperature.svg" />
										</span>
										<span className="flex-col">{temperature}°C</span>
									</div>
									<div className="flex-row pt-5px">
										<span className="flex-col w-12 max-h-full">
											<img src="/assets/icons/dashboard/humidity.svg" />
										</span>
										<span className="flex-col">{humidity}%</span>
									</div>
									<div className="flex-row pt-5px">
										<span className="flex-col w-12 max-h-full">
											<img src="/assets/icons/dashboard/moisture.svg" />
										</span>
										<span className="flex-col">0%</span>
									</div>
								</div>
								<div className="flex-col ml-5">
									<div className="flex-row pt-5px">
										<span className="flex-col w-12 max-h-full">
											<img src="/assets/icons/dashboard/waterLevel.svg" />
										</span>
										<span className="flex-col">0cm</span>
									</div>
									<div className="flex-row pt-5px">
										<span className="flex-col w-12 max-h-full">
											<img src="/assets/icons/dashboard/waterOverdrawn.svg" />
										</span>
										<span className="flex-col">0l</span>
									</div>
								</div>
							</div>
						</div>
						<div className="flex-col w-8/12">
							<div className="flex-row h-44 -mt-2">
								<LiveMeasurementsChart chartType={settings.chartType} />
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
											{weather?.map((weatherItem: any) => {
												return (
													<div className="flex-col mx-2 mb-2 weatherWrapper">
														<div className="flex-row flex-center">
															<div className="flex-col">
																<span className="text-2xl">{/..(?=:).../.exec(weatherItem.time)}</span>
															</div>
														</div>
														<div className="flex-col">
															<div className="flex-row flex-center text-xs">
																<img src="/assets/icons/dashboard/temperature.svg" className="w-8 max-h-full" />
																<span> {weatherItem.data.instant.details.air_temperature}°C</span>
															</div>
															<div className="flex-row flex-center text-xs">
																<img src="/assets/icons/dashboard/humidity.svg" className="w-8 max-h-full" />
																<span> {weatherItem.data.instant.details.relative_humidity}%</span>
															</div>
														</div>
														<div className="flex-row">
															<div className="flex-col flex-center">
																<img
																	src={`/assets/icons/weatherForecast/${weatherForecastBullshitToIcons(
																		weatherItem?.data.next_1_hours?.summary?.symbol_code
																	)}.svg`}
																	className="w-20 max-h-full"
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
									<div className="flex-row h-80">
										<IrrigationChart chartType={settings.chartType} />
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
							<div className="flex-row h-60">
								<WaterConsumptionChart chartType={settings.chartType} />
							</div>
							<div className="flex-row mt-3">
								<span>Historie měření</span>
							</div>
							<div className="flex-row h-72">
								<MeasurementsHistoryChart chartType={settings.chartType} />
							</div>
						</CardContent>
					</Card>
				</div>
			</div>
		</div>
	)
}
