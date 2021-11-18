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
	const [weatherFetched, setWeatherFetched] = useState(true)
	const [currentLocation, setCurrentLocation] = useState('')

	const fetchWeatherForecast = () => {
		if (weatherFetched) {
			axios
				.request({
					method: 'GET',
					url: `${process.env.REACT_APP_FORECAST_API_URL}${currentLocation}`,
					headers: {
						'Content-Type': 'application/json',
					},
				})
				.then((response) => {
					setWeather(response)
					setWeatherFetched(false)
				})
				.catch((error) => {
					console.error(error)
				})
		}
	}

	//setInterval(() => fetchWeatherForecast(), weatherFetched ? 1000 : 600000)
	// if (weatherFetched) {
	// 	//setTimeout(() => () => fetchWeather(), 1000)

	// }

	// 10min
	setInterval(() => fetchWeatherForecast(), 10000)

	fetch('/measure')
		.then((res) => res.json())
		.then((data) => {
			setTemperature(data.temperature)
			setHumidity(data.humidity)
		})
		.then((err) => {
			console.log(err)
		})

	fetch('/init/measured')
		.then((res) => res.json())
		.then((data) => {
			setCurrentLocation(data.location)
		})
		.then((err: any) => {
			console.log(err)
		})

	console.log('weather')
	console.log(weather)
	console.log(JSON.stringify(weather))

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
							<div className={`row ${classes.cardRow}`}>
								<span>Teplota: {temperature}°C</span>
							</div>
							<div className={`row ${classes.cardRow}`}>
								<span>Vlhkost vzduchu: {humidity}%</span>
							</div>
							<div className={`row ${classes.cardRow}`}>
								<span>Vlhkost půdy: 0%</span>
							</div>
						</div>
						<div className="col">
							<div className={`row ${classes.cardRow}`}>
								<span>Výška hladiny vody: 0cm</span>
							</div>
							<div className={`row ${classes.cardRow}`}>
								<span>Vody vyčerpáno: 0l</span>
							</div>
						</div>
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
										{weather?.data?.list?.map((weatherItem: any) => {
											console.log(weatherItem.main.temp)
											return (
												<div className="col">
													<span className="row">{weatherItem.dt_txt}</span>
													<span className="row">{weatherItem.main.temp}°C</span>
													<span className="row">{weatherItem.main.humidity}%</span>
													<span className="row">
														<img src={`/assets/icons/weatherForecast/${weatherItem.weather[0]?.icon}.png`} />
													</span>
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
