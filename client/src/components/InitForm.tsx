import { useEffect, useState, Component } from 'react'
import axios from 'axios'
import { useMutation, useQuery } from '@apollo/client'
import { createSettingsMut, measurements } from '../graphql/mutations'
import { Mutation } from '@apollo/react-components';
import Map from './Map'
import Card from '@material-ui/core/Card'
import CardContent from '@material-ui/core/CardContent'
import TextInputField from './fields/TextInputField'
import ToggleButton from './buttons/ToggleButton'
import SaveButton from './buttons/SaveButton'
import { settings } from './../graphql/queries'

export default function InitForm(props: any) {
	const { data: isSettings } = useQuery(settings),
		[formActiveState, setFormActiveState] = useState(false),
		[saveButtonState, setSaveButtonState] = useState(true),
		[automaticIrrigationState, setAutomaticIrrigationState] = useState(true),
		[irrigationDuration, setIrrigationDuration] = useState<number>(),
		[scheduledIrrigationState, setScheduledIrrigationState] = useState(false),
		[limitValues, setLimitValues] = useState<any>(),
		[moistLimit, setMoistLimit] = useState<number>(),
		[waterAmountLimit, setWaterAmountLimit] = useState<number>(),
		[waterLevelLimit, setWaterLevelLimit] = useState<number>(),
		[hourRange, setHourRange] = useState<number>(),
		[initCoords, setInitCoords] = useState(true),
		[location, setLocation] = useState<string>(),
		[coords, setCoords] = useState(false),
		[latitude, setLatitude] = useState<number>(),
		[longitude, setLongitude] = useState<number>(),
		[mapClicked, setMapClicked] = useState(false)

	let initMeasurementsInterval: any, fetchLocationFromCoordsInterval: any, fetchLocationFromCoordsFixingInterval: any

	useEffect(() => {
		isSettings?.getSettings.length < 1 && setFormActiveState(true)

		formActiveState &&
			!initMeasurementsInterval &&
			(initMeasurementsInterval = setInterval(() => initMeasurements(), 3000))
	}, [isSettings])

	const [createSettings] = useMutation(createSettingsMut, {
		variables: {
			limits_trigger: automaticIrrigationState,
			water_level_limit: waterLevelLimit,
			water_amount_limit: waterAmountLimit,
			moist_limit: moistLimit,
			scheduled_trigger: scheduledIrrigationState,
			hour_range: hourRange,
			location: location,
			irrigation_duration: irrigationDuration,
			chart_type: true,
			language: false,
			theme: false,
			lat: latitude,
			lon: longitude,
		},
		refetchQueries: [{ query: settings }],
	})

	interface IGetCoordsProps {
		label: string
	}

	// Initial coords from user's position
	class GetCoords extends Component<IGetCoordsProps> {
		constructor(props: any) {
			super(props)
			this.state = {}
		}

		componentDidMount() {
			navigator.geolocation.getCurrentPosition((position) => {
				setLatitude(position.coords.latitude)
				setLongitude(position.coords.longitude)
				setInitCoords(false)
			})
		}
		render(): JSX.Element {
			return <></>
		}
	}

	const fetchLocationFromCoords = () => {
		axios
			.request({
				method: 'GET',
				url: `https://api.opencagedata.com/geocode/v1/json?q=${latitude}+${longitude}&key=${process.env.REACT_APP_GEOCODE_API_KEY}`,
				headers: {
					'Content-Type': 'application/json',
				},
			})
			.then((res) => {
				setLocation(
					res.data.results[0].components?.town ||
					res.data.results[0].components?.village ||
					res.data.results[0].components?.city
				)
				setCoords(true)
			})
			.catch((error) => {
				console.error(error)
			})
	}

	formActiveState && setTimeout(() => fetchLocationFromCoords(), coords ? 100_000_000 : 1000)

	const fetchCoordsFromLocation = (searchLocationValue: any) => {
		axios
			.request({
				method: 'GET',
				url: `https://api.opencagedata.com/geocode/v1/json?q=${searchLocationValue}&key=${process.env.REACT_APP_GEOCODE_API_KEY}`,
				headers: {
					'Content-Type': 'application/json',
				},
			})
			.then((res) => {
				res.data.results.map((result: any) => {
					if (result.components.country_code === 'cz') {
						setLatitude(result?.geometry.lat)
						setLongitude(result?.geometry.lng)
					}
				})
				setCoords(true)
			})
			.catch((error) => {
				console.error(error)
			})
	}

	const initMeasurements = () => {
		axios
			.request({
				method: 'GET',
				url: `${process.env.REACT_APP_GO_API_URL}/init/measured`,
				headers: {
					'Content-Type': 'application/json',
				},
			})
			.then((res) => {
				console.log(res)
				setMoistLimit(res.data.moistLimit)
				setWaterLevelLimit(res.data.waterLevelLimit)
				console.log(moistLimit)
				console.log(waterLevelLimit)
				setLimitValues(true)
			})
			.catch((error) => {
				console.error(error)
			})
	}

	const updateToggleState = (type: string) => {
		if (type === 'automaticIrrigation') {
			if (automaticIrrigationState === false) {
				setAutomaticIrrigationState(true)
				setSaveButtonState(true)
			} else {
				setAutomaticIrrigationState(false)
				if (scheduledIrrigationState === false) {
					setSaveButtonState(false)
				}
			}
		}

		if (type === 'scheduledIrrigation') {
			if (scheduledIrrigationState === false) {
				setScheduledIrrigationState(true)
				setSaveButtonState(true)
				irrigationDuration == 0 ? setSaveButtonState(false) : setSaveButtonState(true)
			} else {
				setScheduledIrrigationState(false)
				if (automaticIrrigationState === false) {
					setSaveButtonState(false)
				}
			}
		}
	}

	useEffect(() => {
		if (
			document.querySelector(
				'#root > div > div.init-form > div > div > div > div.flex-col.pl-5 > div > div > div > div:nth-child(2) > div:nth-child(2) > div > div:nth-child(4) > div'
			)
		) {
			const address = document.getElementsByClassName('address-line')
			for (var i = 0; i < address.length; i++) {
				if (address[i].getAttribute('jsinstance') === `${i}` || /(\d\d\d \d\d)/.exec(address[i].innerHTML)) {
					const selectedAddress = address[i].innerHTML.replace(`${/...\s\d\d\s/.exec(address[i].innerHTML)}`, '')
					fetchCoordsFromLocation(selectedAddress)
					setLocation(`${selectedAddress}`)
					document.querySelectorAll('#location')[0].innerHTML = selectedAddress
				}
			}
		}
		setMapClicked(false)
	}, [mapClicked])

	return (
		<>
			{formActiveState && (
				<div className="init-form">
					<Card className="card p-0-i">
						<CardContent className="p-0-i">
							<div className="flex-row">
								<div className="flex-col pl-8 pt-8 pr-3 pb-8">
									<div className="flex-row flex-center p-1 mb-2">
										<span className="title-1">PlantHub - Inicializace</span>
									</div>
									<div className="flex-row mb-2">
										<div className="flex-col p-1 pt-5px mt-2">
											<div className="flex-row">
												<span className="title-2">Automaticky</span>
											</div>
											<div className="flex-row mt-2">
												<span className="title-2">Plánovaně</span>
											</div>
										</div>
										<div className="flex-col p-1 pt-5px mt-2 ml-2">
											<div className="flex-row">
												<div onClick={() => updateToggleState('automaticIrrigation')}>
													<ToggleButton item="limitsTrigger" toggleState={automaticIrrigationState} />
												</div>
											</div>
											<div className="flex-row mt-2">
												<div onClick={() => updateToggleState('scheduledIrrigation')}>
													<ToggleButton item="scheduledTrigger" toggleState={scheduledIrrigationState} />
												</div>
											</div>
										</div>
									</div>
									<div
										className="flex-row p-1 pt-5px mt-2"
										onBlur={(data: any) => setIrrigationDuration(data.target.value)}
										onChange={(data: any) =>
											data.target.value == 0 ? setSaveButtonState(false) : setSaveButtonState(true)
										}
									>
										<TextInputField
											item="irrigationDuration"
											name="Doba zavlažování (s)"
											defaultValue={irrigationDuration}
											active={true}
										/>
									</div>
									<div className="flex-row p-1 pt-5px mt-2" onBlur={(data: any) => setMoistLimit(data.target.value)}>
										<TextInputField
											item="moistLimit"
											name="Limit vlhkosti půdy (%)"
											defaultValue={moistLimit}
											active={automaticIrrigationState}
										/>
									</div>
									<div
										className="flex-row p-1 pt-5px mt-2"
										onBlur={(data: any) => setWaterAmountLimit(data.target.value)}
									>
										<TextInputField
											item="waterAmountLimit"
											name="Limit objemu nádrže (l)"
											defaultValue={waterAmountLimit}
											active={automaticIrrigationState}
										/>
									</div>
									<div
										className="flex-row p-1 pt-5px mt-2"
										onBlur={(data: any) => setWaterLevelLimit(data.target.value)}
									>
										<TextInputField
											item="waterLevelLimit"
											name="Limit hladiny vody (cm)"
											defaultValue={waterLevelLimit}
											active={automaticIrrigationState}
										/>
									</div>
									<div className="flex-row p-1 pt-5px mt-2" onBlur={(data: any) => setHourRange(data.target.value)}>
										<TextInputField
											item="hourRange"
											name="Rozsah hodin (h)"
											defaultValue={hourRange}
											active={scheduledIrrigationState}
										/>
									</div>
									<div
										className="flex-row p-1 pt-5px mt-2"
										onBlur={(data: any) => fetchCoordsFromLocation(data.target.value)}
									>
										<TextInputField item="location" name="Lokace" defaultValue={location} active="true" />
									</div>
									<div className="flex-row p-1 pt-5px mt-2">
										<div
											onClick={() => {
												createSettings()
												setFormActiveState(false)
											}}
										>
											<SaveButton active={saveButtonState} />
										</div>
									</div>
								</div>
								<div className="flex-col pl-5" onClick={() => setMapClicked(true)}>
									<Map lat={latitude || 0} lon={longitude || 0} />
								</div>
							</div>
						</CardContent>
					</Card>
					{initCoords && <GetCoords {...props} />}
				</div>
			)}
		</>
	)
}
