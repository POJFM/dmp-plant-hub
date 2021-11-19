import { useEffect, useState } from 'react'
import axios from 'axios'
import Card from '@material-ui/core/Card'
import CardContent from '@material-ui/core/CardContent'
import { RefreshButton } from './buttons/RefreshButton'
import {
	Chart,
	CategoryScale,
	LineController,
	LineElement,
	PointElement,
	LinearScale,
	Title,
	ChartType,
} from 'chart.js'
import { useStyles } from './../styles/rootStyles'
import { useDashboardStyles } from './../styles/dashboard'

export default function Dashboard() {
	const dashboardClasses = useDashboardStyles()
	const classes = useStyles()
	const [temperature, setTemperature] = useState(0)
	const [humidity, setHumidity] = useState(0)
	const [weather, setWeather] = useState<any>()
	const [currentLatitude, setCurrentLatitude] = useState('')
	const [currentLongitude, setCurrentLongitude] = useState('')

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
			setWeather(response.data.properties.timeseries.slice(0, 5))
		})
		.catch((error) => {
			console.error(error)
		})
	}

	// Fetch on render then every 30mins
	setTimeout(() => fetchWeatherForecast(), weather ? 300_000 : 1)

	useEffect(() => {
		document.title = 'Plant Hub | Dashboard'

		Chart.register(CategoryScale, LineController, LineElement, PointElement, LinearScale, Title)

		const irrigationChart = new Chart('irrigationChart', {
			type: 'line',
			data: {
				labels: ['January', 'February', 'March', 'April', 'May', 'June'],
				datasets: [
					{
						label: 'Moisture',
						backgroundColor: 'rgb(255, 99, 132)',
						borderColor: 'rgb(0, 0, 255)',
						data: [5, 8, 10, 20, 5, 35, 20],
						yAxisID: 'm',
					},
					{
						label: 'Temperature',
						backgroundColor: 'rgb(255, 99, 132)',
						borderColor: 'rgb(0, 255, 0)',
						data: [8, 1, 10, 20, 5, 35, 20],
						yAxisID: 't',
					},
					{
						label: 'Humidity',
						backgroundColor: 'rgb(255, 99, 132)',
						borderColor: 'rgb(255, 0, 0)',
						data: [10, 7, 10, 20, 5, 35, 20],
						yAxisID: 'h',
					},
				],
			},
			options: {
				interaction: {
					mode: 'index',
					intersect: false,
				},
				scales: {
					m: {
						type: 'linear',
						display: true,
						position: 'left',
					},
					t: {
						type: 'linear',
						display: true,
						position: 'left',
					},
					h: {
						type: 'linear',
						display: true,
						position: 'left',
					},
				},
			},
		})
	}, [])

	const weatherForecastBullshitToIcons = (weatherState: string) => {
		if (weatherState.includes('rainandthunder')) {
			return 'rainAndThunder'
		} else if (weatherState.includes('snowandthunder')) {
			return 'snowAndThunder'
		} else if (weatherState.includes('partlycloudy')) {
			if (weatherState.includes('night')) {
				return 'partlyCloudyNight'
			} else {
				return 'partlyCloudy'
			}
		} else if (weatherState.includes('clearsky')) {
			if (weatherState.includes('night')) {
				return 'clearskyNight'
			} else {
				return 'clearsky'
			}
		} else if (weatherState.includes('fair')) {
			return 'fair'
		} else if (weatherState.includes('snow')) {
			return 'snow'
		} else if (weatherState.includes('fog')) {
			return 'fog'
		} else if (weatherState.includes('rain')) {
			return 'rain'
		} else if (weatherState.includes('sleet')) {
			return 'sleet'
		} else if (weatherState.includes('cloudy')) {
			return 'cloudy'
		}
	}

	// const waterConsumptionChart = new Chart('waterConsumptionChart', {
	// 	type: 'line',
	// 	data: {
	// 		labels: ['January', 'February', 'March', 'April', 'May', 'June'],
	// 		datasets: [
	// 			{
	// 				label: 'Moisture',
	// 				backgroundColor: 'rgb(255, 99, 132)',
	// 				borderColor: 'rgb(0, 0, 255)',
	// 				data: [5, 8, 10, 20, 5, 35, 20],
	// 				yAxisID: 'm',
	// 			},
	// 			{
	// 				label: 'Temperature',
	// 				backgroundColor: 'rgb(255, 99, 132)',
	// 				borderColor: 'rgb(0, 255, 0)',
	// 				data: [8, 1, 10, 20, 5, 35, 20],
	// 				yAxisID: 't',
	// 			},
	// 			{
	// 				label: 'Humidity',
	// 				backgroundColor: 'rgb(255, 99, 132)',
	// 				borderColor: 'rgb(255, 0, 0)',
	// 				data: [10, 7, 10, 20, 5, 35, 20],
	// 				yAxisID: 'h',
	// 			},
	// 		],
	// 	},
	// 	options: {
	// 		responsive: false,
	// 		maintainAspectRatio: false,
	// 		interaction: {
	// 			mode: 'index',
	// 			intersect: false,
	// 		},
	// 		scales: {
	// 			m: {
	// 				type: 'linear',
	// 				display: true,
	// 				position: 'left',
	// 			},
	// 			t: {
	// 				type: 'linear',
	// 				display: true,
	// 				position: 'left',
	// 			},
	// 			h: {
	// 				type: 'linear',
	// 				display: true,
	// 				position: 'left',
	// 			},
	// 		},
	// 	},
	// })
	return (
		<div className="col dashboard">
			<Card className={classes.card}>
				<CardContent>
					<div className="row">
						<div className="col">
							<div className="row">
								<div className="col">
									<div className={`row ${classes.cardRow}`}>
										<span className="col">
											<img src="/assets/icons/dashboard/temperature.svg" className={classes.icon} />
										</span>
										<span className="col">{temperature}°C</span>
									</div>
									<div className={`row ${classes.cardRow}`}>
										<span className="col">
											<img src="/assets/icons/dashboard/humidity.svg" className={classes.icon} />
										</span>
										<span className="col">{humidity}%</span>
									</div>
									<div className={`row ${classes.cardRow}`}>
										<span className="col">
											<img src="/assets/icons/dashboard/moisture.svg" className={classes.icon} />
										</span>
										<span className="col">0%</span>
									</div>
								</div>
								<div className="col">
									<div className={`row ${classes.cardRow}`}>
										<span className="col">
											<img src="/assets/icons/dashboard/waterLevel.svg" className={classes.icon} />
										</span>
										<span className="col">0cm</span>
									</div>
									<div className={`row ${classes.cardRow}`}>
										<span className="col">
											<img src="/assets/icons/dashboard/waterOverdrawn.svg" className={classes.icon} />
										</span>
										<span className="col">0l</span>
									</div>
								</div>
							</div>
						</div>
						<div className="col"></div>
					</div>
				</CardContent>
			</Card>
			<div className="row">
				<div className="col">
					<div className="row">
						<div className="col">
							<Card className={`${classes.card} ${classes.cardTwoLeft}`}>
								<CardContent>
									<div className="row">
										{weather &&
											weather?.map((weatherItem: any) => {
												return (
													<div className="col">
														<div className="row">
															<div className="col flex-center">
																<span className={dashboardClasses.weatherForecastTime}>
																	{/..(?=:).../.exec(weatherItem.time)}
																</span>
															</div>
														</div>
														<div className="row">
															<div className={`col flex-center ${dashboardClasses.weatherForecastValue}`}>
																<span>
																	<img
																		src="/assets/icons/dashboard/temperature.svg"
																		className={dashboardClasses.weatherForecastValueIcon}
																	/>
																</span>
																<span> {weatherItem.data.instant.details.air_temperature}°C</span>
															</div>
															<div className={`col flex-center ${dashboardClasses.weatherForecastValue}`}>
																<span>
																	<img
																		src="/assets/icons/dashboard/humidity.svg"
																		className={dashboardClasses.weatherForecastValueIcon}
																	/>
																</span>
																<span> {weatherItem.data.instant.details.relative_humidity}%</span>
															</div>
														</div>
														<div className="row">
															<div className="col flex-center">
																<img
																	src={`/assets/icons/weatherForecast/${weatherForecastBullshitToIcons(
																		weatherItem.data.next_1_hours.summary.symbol_code
																	)}.svg`}
																	className={dashboardClasses.weatherForecastIcon}
																/>
															</div>
														</div>
													</div>
												)
											})}
									</div>
								</CardContent>
							</Card>
						</div>
					</div>
					<div className="row">
						<div className="col">
							<Card className={`${classes.card} ${classes.cardTwoLeft}`}>
								<CardContent>
									<div className="row">
										<span>Historie zavlažování</span>
									</div>
									<div /* className={`row ${dashboardClasses.canvas}`} */>
										<canvas id="irrigationChart"></canvas>
									</div>
								</CardContent>
							</Card>
						</div>
					</div>
				</div>
				<div className="col">
					<Card className={`${classes.card} ${classes.cardTwoRight}`}>
						<CardContent>
							<div className="row">
								<p>Spotřeba vody</p>
							</div>
							<div className={`row ${dashboardClasses.canvas}`}>
								{/* <canvas id="waterConsumtionChart" width="240px" height="240px"></canvas> */}
							</div>
						</CardContent>
					</Card>
				</div>
			</div>
		</div>
	)
}
