import { useEffect, useState, Component } from 'react'
import axios from 'axios'
import { useTranslation } from 'react-i18next'
import { useMutation, useQuery } from '@apollo/client'
import { createSettingsMut } from '../graphql/mutations'
import Map from './Map'
import Card from '@material-ui/core/Card'
import CardContent from '@material-ui/core/CardContent'
import TextInputField from './fields/TextInputField'
import ToggleButton from './buttons/ToggleButton'
import SaveButton from './buttons/SaveButton'
import { settings } from './../graphql/queries'

export default function InitForm(props: any) {
	const { t } = useTranslation()
	const { data: isSettings } = useQuery(settings)
	const [formActiveState, setFormActiveState] = useState(props.test || false),
		[saveButtonState, setSaveButtonState] = useState(true),
		[automaticIrrigationState, setAutomaticIrrigationState] = useState(true),
		[irrigationDuration, setIrrigationDuration] = useState<number>(),
		[defaultWaterAmount, setDefaultWaterAmount] = useState<number>(),
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

	let initMeasurementsInterval: ReturnType<typeof setTimeout>,
		fetchLocationFromCoordsInterval: ReturnType<typeof setTimeout>,
		fetchLocationFromCoordsFixingInterval: ReturnType<typeof setTimeout>

	useEffect(() => {
		isSettings?.getSettings.length < 1 && setFormActiveState(true)
	}, [isSettings])

	useEffect(() => {
		!initMeasurementsInterval &&
			(initMeasurementsInterval = setInterval(() => initMeasurements(), 3000))
	}, [formActiveState])

	const [createSettings] = useMutation(createSettingsMut, {
		variables: {
			id: 0,
			limits_trigger: automaticIrrigationState,
			water_level_limit: waterLevelLimit && waterLevelLimit - waterLevelLimit * 0.1,
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
			default_water_amount: 5,
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
			navigator.geolocation?.getCurrentPosition((position) => {
				setLatitude(position.coords.latitude)
				setLongitude(position.coords.longitude)
				axios
					.post(
						`${process.env.REACT_APP_GO_API_URL}/init/measured`,
						{
							lat: latitude,
							lon: longitude
						},
						{
							headers: {
								'Content-Type': 'application/x-www-form-urlencoded',
							},
						}
					)
					.then((res) => {
						console.log(res)
					})
					.catch((error) => {
						console.error(error)
					})
				formActiveState && setTimeout(() => fetchLocationFromCoords(), coords ? 100_000_000 : 3000)

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
				url: `${process.env.REACT_APP_GO_API_URL}/api/geocodes`,
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

	const fetchCoordsFromLocation = (searchLocationValue: any) => {
		axios
			.request({
				method: 'GET',
				url: `${process.env.REACT_APP_GO_API_URL}/api/geocodes`,
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
				setMoistLimit(res.data.moistLimit)
				setWaterLevelLimit(res.data.waterLevelLimit)
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
										<span className="title-1">PlantHub - {t('initForm.title')}</span>
									</div>
									<div className="flex-row mb-2">
										<div className="flex-col p-1 pt-5px mt-2">
											<div className="flex-row">
												<span className="title-2">{t('initForm.automatic')}</span>
											</div>
											<div className="flex-row mt-2">
												<span className="title-2">{t('initForm.scheduled')}</span>
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
											name={`${t('initForm.irrigationDuration')} (s)`}
											defaultValue={irrigationDuration}
											dataType="number"
											active={true}
										/>
									</div>
									<div
										className="flex-row p-1 pt-5px mt-2"
										onBlur={(data: any) => setDefaultWaterAmount(parseInt(data.target.value))}
										onChange={(data: any) =>
											data.target.value == 0 ? setSaveButtonState(false) : setSaveButtonState(true)
										}
									>
										<TextInputField
											item="defaultWaterAmount"
											name={`${t('initForm.defaultWaterAmount')} (l)`}
											defaultValue={defaultWaterAmount}
											dataType="number"
											active={true}
										/>
									</div>
									<div
										className="flex-row p-1 pt-5px mt-2"
										onBlur={(data: any) => setMoistLimit(data.target.value)}
									>
										<TextInputField
											item="moistLimit"
											name={`${t('initForm.moistLimit')} (%)`}
											defaultValue={moistLimit ? Math.round(moistLimit * 100) / 100 : undefined}
											dataType="number"
											active={automaticIrrigationState}
										/>
									</div>
									<div
										className="flex-row p-1 pt-5px mt-2"
										onBlur={(data: any) => setWaterAmountLimit(data.target.value)}
									>
										<TextInputField
											item="waterAmountLimit"
											name={`${t('initForm.waterAmountLimit')} (l)`}
											defaultValue={waterAmountLimit}
											dataType="number"
											active={automaticIrrigationState}
										/>
									</div>
									<div
										className="flex-row p-1 pt-5px mt-2"
										onBlur={(data: any) => setWaterLevelLimit(data.target.value)}
									>
										<TextInputField
											item="waterLevelLimit"
											name={`${t('initForm.waterLevelLimit')} (cm)`}
											defaultValue={waterLevelLimit ? Math.round(waterLevelLimit * 100) / 100 : undefined}
											dataType="number"
											active={automaticIrrigationState}
										/>
									</div>
									<div className="flex-row p-1 pt-5px mt-2" onBlur={(data: any) => setHourRange(data.target.value)}>
										<TextInputField
											item="hourRange"
											name={`${t('initForm.hourRange')} (h)`}
											defaultValue={hourRange}
											dataType="number"
											active={scheduledIrrigationState}
										/>
									</div>
									<div
										className="flex-row p-1 pt-5px mt-2"
										onBlur={(data: any) => {
											setLocation(data.target.value)
											fetchCoordsFromLocation(data.target.value)
										}}
									>
										<TextInputField
											item="location"
											name={t('initForm.location')}
											defaultValue={location}
											dataType="string"
											active={true}
										/>
									</div>
									<div className="flex-row p-1 pt-5px mt-2">
										<div
											onClick={() => {
												createSettings()
												setFormActiveState(false)
											}}
										>
											<SaveButton active={saveButtonState} name="initsave" />
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
