import { useEffect, useState, useMemo } from 'react'
import axios from 'axios'
import { useTranslation } from 'react-i18next'
import { MuiCard as Card } from './Card'
import { useQuery } from '@apollo/client'
import { dashboard } from '../graphql/queries'
import { LiveMeasurementsChart, WaterConsumptionChart, IrrigationChart, MeasurementsHistoryChart } from './Chart'
import ReactScrollWheelHandler from 'react-scroll-wheel-handler'
import { getMonths, monthRegex, dayRegex, timeRegex } from 'src/utils'

export default function Dashboard() {
	const { t } = useTranslation()
	const [liveMeasure, setLiveMeasure] = useState(false),
		[currentTemp, setCurrentTemp] = useState(0),
		[currentMoist, setCurrentMoist] = useState(0),
		[currentHum, setCurrentHum] = useState(0),
		[overallWaterConsumption, setOverallWaterConsumption] = useState(0),
		[temp, setTemp] = useState<any>([]),
		[moist, setMoist] = useState<any>([]),
		[hum, setHum] = useState<any>([]),
		[irrigationCount, setIrrigationCount] = useState<any>(),
		[measurements, setMeasurements] = useState<any>(),
		[irrigationHistory, setIrrigationHistory] = useState<any>(),
		[waterOverdrawn, setWaterOverdrawn] = useState<any>(),
		[months, setMonths] = useState<any>(),
		[weather, setWeather] = useState<any>()

	const { loading, error, data } = useQuery(dashboard),
		chartType = useMemo(() => {
			return data?.getSettings[0]?.chart_type === false ? 0 : 1
		}, [data])

	let liveMeasurementsInterval: ReturnType<typeof setTimeout>,
		weatherForecastInterval: ReturnType<typeof setTimeout>,
		arrayPassTemp: number[] = [],
		arrayPassHum: number[] = [],
		arrayPassMoist: number[] = []

	useEffect(() => {
		let measurementsDataNotMonth: number[] = [],
			measurementsDataNotAvg: number[] = [],
			irrigationCountObj: number[] = [],
			waterOverdrawnObj: number[] = [],
			irrigationHistoryData: {
				moist: number[],
				temp: number[],
				hum: number[],
				water_overdrawn: number[],
				dataframe: string[]
			} = {
				moist: [],
				temp: [],
				hum: [],
				water_overdrawn: [],
				dataframe: []
			},
			measurementsData: {
				moist: number[],
				temp: number[],
				hum: number[]
			} = {
				moist: [],
				temp: [],
				hum: [],
			}

		data?.getMeasurements.filter(
			(filteredData: any) => {
				if (filteredData?.with_irrigation === true) {
					irrigationHistoryData?.moist?.push(filteredData?.moist)
					irrigationHistoryData?.hum?.push(filteredData?.hum)
					irrigationHistoryData?.temp?.push(filteredData?.temp)
				} else {
					measurementsDataNotMonth.push(filteredData)
				}
			}
		)

		data?.getIrrigation.map((item: any) => {
			let tMonth = monthRegex(item?.timestamp),
				tDay = dayRegex(item?.timestamp),
				tTime = timeRegex(item?.timestamp)

			irrigationHistoryData?.water_overdrawn?.push(item?.water_overdrawn)
			irrigationHistoryData?.dataframe?.push(`${tDay}.${tMonth}. ${tTime}`)

			setOverallWaterConsumption((value: number) => value + item?.water_overdrawn)
		})

		setMonths(getMonths())

		const currentMonth = new Date().getMonth() + 1

		// Extract irrigation count for each month
		for (let i = 1; i < currentMonth + 1; i++) {
			let month = 0,
				measurementsInMonth: any = [],
				waterOverdrawnInMonth = 0

			data?.getIrrigation.map((item: any) => {
				let tMonth = monthRegex(item?.timestamp)
				if (i === tMonth) {
					month++
					waterOverdrawnInMonth += item?.water_overdrawn
				}
			})

			irrigationCountObj?.push(month)
			waterOverdrawnObj?.push(waterOverdrawnInMonth)

			measurementsDataNotMonth.map((item: any) => {
				let tMonth = monthRegex(item?.timestamp)

				i === tMonth && measurementsInMonth?.push(item)
			})

			measurementsDataNotAvg?.push(measurementsInMonth)

			// i > 11 && (i = 0)
			// i === currentMonth - 1 && (i = 13)
		}

		measurementsDataNotAvg.map((month: any) => {
			let moistAvg = 0, tempAvg = 0, humAvg = 0

			month.map((item: any) => {
				moistAvg += item?.moist
				tempAvg += item?.temp
				humAvg += item?.hum
			})

			measurementsData?.moist?.push(moistAvg / month.length)
			measurementsData?.hum?.push(tempAvg / month.length)
			measurementsData?.temp?.push(humAvg / month.length)
		})

		setIrrigationCount(irrigationCountObj)
		setWaterOverdrawn(waterOverdrawnObj)
		setIrrigationHistory(irrigationHistoryData)
		setMeasurements(measurementsData)
	}, [data])

	useEffect(() => {
		document.title = 'Plant Hub | Dashboard'

		fetchWeatherForecast()

		!weatherForecastInterval && (weatherForecastInterval = setInterval(() => fetchWeatherForecast(), 300_000))
		!liveMeasurementsInterval && (liveMeasurementsInterval = setInterval(() => liveMeasurements(), 1000))
	}, [])

	const fetchWeatherForecast = () => {
		axios
			.request({
				method: 'GET',
				url: `${process.env.REACT_APP_GO_API_URL}/api/weather`,
				headers: {
					'Content-Type': 'application/json',
				},
			})
			.then((res) => {
				setWeather(res.data.hourly.slice(0, 15))
				console.log(res)
			})
		// .catch((error) => {
		// 	console.error(error)
		// })
	}

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

				arrayPassTemp?.push(res.data.temp)
				arrayPassHum?.push(res.data.hum)
				arrayPassMoist?.push(res.data.moist)

				arrayPassTemp.length > 25 && (arrayPassTemp = arrayPassTemp?.slice(1, 25))
				arrayPassHum.length > 25 && (arrayPassHum = arrayPassHum?.slice(1, 25))
				arrayPassMoist.length > 25 && (arrayPassMoist = arrayPassMoist?.slice(1, 25))

				setTemp(arrayPassTemp)
				setHum(arrayPassHum)
				setMoist(arrayPassMoist)

				setLiveMeasure(true)
			})
			.catch((error) => {
				console.error(error)
			})
	}

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

	console.log(Math.round(data?.getIrrigation[data?.getIrrigation.length - 1]?.water_level * 100) / 100)

	return (
		<div className="dashboard">
			<Card className="card h-52">
				<div className="flex-row">
					<div className="flex-col w-4/12">
						<div className="flex-row">
							<div className="flex-col ml-5 w-32">
								<div className="flex-row pt-5px" title={t('dashboard.temperature')}>
									<span className="flex-col w-12 max-h-full">
										<img src="/assets/icons/dashboard/temperature.svg" />
									</span>
									<span className="flex-col w-18 flex-center ml-2">{`${Math.round(currentTemp * 100) / 100} °C`}</span>
								</div>
								<div className="flex-row pt-5px" title={t('dashboard.humidity')}>
									<span className="flex-col w-12 max-h-full">
										<img src="/assets/icons/dashboard/humidity.svg" />
									</span>
									<span className="flex-col w-18 flex-center ml-2">{`${Math.round(currentHum * 100) / 100} %`}</span>
								</div>
								<div className="flex-row pt-5px" title={t('dashboard.moisture')}>
									<span className="flex-col w-12 max-h-full">
										<img src="/assets/icons/dashboard/moisture.svg" />
									</span>
									<span className="flex-col w-18 flex-center ml-2">{`${Math.round(currentMoist * 100) / 100} %`}</span>
								</div>
							</div>
							<div className="flex-col ml-5 w-32">
								<div className="flex-row pt-5px" title={t('dashboard.waterLevel')}>
									<span className="flex-col w-12 max-h-full">
										<img src="/assets/icons/dashboard/waterLevel.svg" />
									</span>
									<span className="flex-col flex-center ml-2">{`${data?.getIrrigation[0] ? Math.round(data?.getIrrigation[data?.getIrrigation.length - 1]?.water_level * 100) / 100 : 0
										} cm`}</span>
								</div>
								<div className="flex-row pt-5px" title={t('dashboard.waterAmount')}>
									<span className="flex-col w-12 max-h-full">
										<img src="/assets/icons/dashboard/waterAmount.svg" />
									</span>
									<span className="flex-col flex-center ml-2">{`${data?.getIrrigation[0] ? Math.round(data?.getIrrigation[data?.getIrrigation.length - 1]?.water_amount * 100) / 100 : 0
										} l`}</span>
								</div>
								<div className="flex-row pt-5px" title={t('dashboard.waterOverdrawn')}>
									<span className="flex-col w-12 max-h-full">
										<img src="/assets/icons/dashboard/waterOverdrawn.svg" />
									</span>
									<span className="flex-col flex-center ml-2">{`${data?.getIrrigation[0] ? Math.round(overallWaterConsumption * 100) / 100 : 0
										} l`}</span>
								</div>
							</div>
						</div>
					</div>
					<div className="flex-col w-8/12">
						<div className="flex-row h-44 -mt-2">
							<LiveMeasurementsChart
								chartType={chartType}
								temp={temp}
								hum={hum}
								moist={moist}
							/>
						</div>
					</div>
				</div>
			</Card>
			<div className="flex-row">
				<div className="flex-col w-6/12">
					<div className="flex-row">
						<div className="flex-col w-full">
							<Card className="card-left">
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
							</Card>
						</div>
					</div>
					<div className="flex-row">
						<div className="flex-col w-full">
							<Card className="card-left">
								<div className="flex-row ">
									<span>{t('dashboard.irrigationHistory')}</span>
								</div>
								<div className="flex-row 2xl:h-96 lg:h-52">
									<IrrigationChart
										chartType={chartType}
										moist={irrigationHistory?.moist}
										hum={irrigationHistory?.hum}
										temp={irrigationHistory?.temp}
										waterOverdrawn={irrigationHistory?.water_overdrawn}
										dataframe={irrigationHistory?.dataframe}
									/>
								</div>
							</Card>
						</div>
					</div>
				</div>
				<div className="flex-col w-6/12">
					<Card className="card-right h-full">
						<div className="flex-row">
							<span>{t('dashboard.waterConsumption')}</span>
						</div>
						<div className="flex-row 2xl:h-64 lg:h-48">
							<WaterConsumptionChart
								chartType={chartType}
								waterOverdrawn={waterOverdrawn}
								irrigationCount={irrigationCount}
							/>
						</div>
						<div className="flex-row mt-3">
							<span>{t('dashboard.measurementsHistory')}</span>
						</div>
						<div className="flex-row 2xl:h-80 lg:h-52">
							<MeasurementsHistoryChart
								chartType={chartType}
								moist={measurements?.moist}
								hum={measurements?.hum}
								temp={measurements?.temp}
							/>
						</div>
					</Card>
				</div>
			</div>
		</div>
	)
}
