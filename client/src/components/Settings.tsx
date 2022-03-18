import { useEffect, useState } from 'react'
import axios from 'axios'
import Card from '@material-ui/core/Card'
import CardContent from '@material-ui/core/CardContent'
import EditableField from './fields/EditableField'
import EnumEditableField from './fields/EnumEditableField'
import ToggleButton from './buttons/ToggleButton'
import SaveButton from './buttons/SaveButton'
import CancelButton from './buttons/CancelButton'
import { useQuery, useMutation } from '@apollo/client'
import { settings } from '../graphql/queries'
import { createSettings } from '../graphql/mutations'
//import { fetchCoordsFromLocation } from 'src/utils'

export default function Settings() {
	const { loading: settingsLoading, error: settingsError, data: settingsData } = useQuery(settings)
	const [updateSettingsData, { data, loading, error }] = useMutation(createSettings)

	console.log("GQL:")
	console.log(data)

	const [buttonsState, setButtonsState] = useState(false), // false
		[automaticIrrigationState, setAutomaticIrrigationState] = useState(
			settingsData?.getSettings[0]?.limits_trigger || true
		),
		[automaticIrrigationStateClass, setAutomaticIrrigationStateClass] = useState<string>(
			settingsData?.getSettings[0]?.limits_trigger ? '#000000' : 'var(--inactiveGrey)'
		),
		[scheduledIrrigationState, setScheduledIrrigationState] = useState(settingsData?.getSettings[0]?.scheduled_trigger),
		[irrigationDuration, setIrrigationDuration] = useState(settingsData?.getSettings[0]?.irrigation_duration),
		[irrigationDurationStateClass, setIrrigationDurationStateClass] = useState('#000000'),
		[scheduledIrrigationStateClass, setScheduledIrrigationStateClass] = useState<string>(
			settingsData?.getSettings[0]?.scheduled_trigger ? '#000000' : 'var(--inactiveGrey)'
		),
		[moistureLimit, setMoistureLimit] = useState(settingsData?.getSettings[0]?.moisture_limit),
		[waterAmountLimit, setWaterAmountLimit] = useState(settingsData?.getSettings[0]?.water_amount_limit),
		[waterLevelLimit, setWaterLevelLimit] = useState(settingsData?.getSettings[0]?.water_level_limit),
		[hourRange, setHourRange] = useState(settingsData?.getSettings[0]?.hour_range),
		[chartTypeState, setChartTypeState] = useState(settingsData?.getSettings[0]?.chart_type),
		[languageState, setLanguageState] = useState(settingsData?.getSettings[0]?.language),
		[themeState, setThemeState] = useState(settingsData?.getSettings[0]?.theme),
		[getCoordsState, setGetCoordsState] = useState(false),
		[getCoords, setGetCoords] = useState<string>(),
		[location, setLocation] = useState(settingsData?.getSettings[0]?.location),
		[latitude, setLatitude] = useState(settingsData?.getSettings[0]?.lat),
		[longitude, setLongitude] = useState(settingsData?.getSettings[0]?.lon)

	console.log(settingsData)
	console.log(settingsError)

	useEffect(() => {
		document.title = 'Plant Hub | Settings'
	}, [])

	useEffect(() => {
		setAutomaticIrrigationState(settingsData?.getSettings[0]?.limits_trigger || true)
		setAutomaticIrrigationStateClass(settingsData?.getSettings[0]?.limits_trigger ? '#000000' : 'var(--inactiveGrey)')
		setScheduledIrrigationState(settingsData?.getSettings[0]?.scheduled_trigger)
		setIrrigationDuration(settingsData?.getSettings[0]?.irrigation_duration)
		setScheduledIrrigationStateClass(settingsData?.getSettings[0]?.scheduled_trigger ? '#000000' : 'var(--inactiveGrey)')
		setMoistureLimit(settingsData?.getSettings[0]?.moisture_limit)
		setWaterAmountLimit(settingsData?.getSettings[0]?.water_amount_limit)
		setWaterLevelLimit(settingsData?.getSettings[0]?.water_level_limit)
		setHourRange(settingsData?.getSettings[0]?.hour_range)
		setChartTypeState(settingsData?.getSettings[0]?.chart_type)
		setLanguageState(settingsData?.getSettings[0]?.language)
		setThemeState(settingsData?.getSettings[0]?.theme)
		setLocation(settingsData?.getSettings[0]?.location)
		setLatitude(settingsData?.getSettings[0]?.lat)
		setLongitude(settingsData?.getSettings[0]?.lon)
	}, [settingsData])

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
				res.data.results.map((item: any) => {
					if (item.components.country_code === 'cz') {
						setLatitude(item?.geometry.lat)
						setLongitude(item?.geometry.lng)
					}
				})
			})
			.catch((error) => {
				console.error(error)
			})
	}

	const updateInputData = (type: string, data: any) => {
		setButtonsState(true)
		type === 'irrigationDuration' && setIrrigationDuration(data?.target?.value)
		type === 'moistureLimit' && setMoistureLimit(data?.target?.value)
		type === 'waterAmountLimit' && setWaterAmountLimit(data?.target?.value)
		type === 'waterLevelLimit' && setWaterLevelLimit(data?.target?.value)
		type === 'hourRange' && setHourRange(data?.target?.value)
		if (type === 'location') {
			setGetCoords(data?.target?.value)
			setGetCoordsState(true)
		}
	}

	useEffect(() => {
		fetchCoordsFromLocation(getCoords)
		setGetCoordsState(false)
	}, [getCoordsState])

	const updateToggleState = (type: string) => {
		if (type === 'automaticIrrigation') {
			if (automaticIrrigationState === false) {
				setAutomaticIrrigationState(true)
				setButtonsState(true)
				setIrrigationDurationStateClass('#000000')
				setAutomaticIrrigationStateClass('#000000')
			} else {
				setAutomaticIrrigationState(false)
				setAutomaticIrrigationStateClass('var(--inactiveGrey)')
				if (scheduledIrrigationState === false) {
					setButtonsState(false)
					setIrrigationDurationStateClass('var(--inactiveGrey)')
				}
			}
		}

		if (type === 'scheduledIrrigation') {
			if (scheduledIrrigationState === false) {
				setScheduledIrrigationState(true)
				setButtonsState(true)
				setIrrigationDurationStateClass('#000000')
				setScheduledIrrigationStateClass('#000000')
				if (irrigationDuration == 0) {
					setButtonsState(false)
				}
			} else {
				setScheduledIrrigationState(false)
				setScheduledIrrigationStateClass('var(--inactiveGrey)')
				if (automaticIrrigationState === false) {
					setButtonsState(false)
					setIrrigationDurationStateClass('var(--inactiveGrey)')
				}
			}
		}

		if (type === 'chartType') {
			setButtonsState(true)
			if (chartTypeState === 0) {
				setChartTypeState(1)
			} else {
				setChartTypeState(0)
			}
		}

		if (type === 'language') {
			setButtonsState(true)
			if (languageState === 0) {
				setLanguageState(1)
			} else {
				setLanguageState(0)
			}
		}

		if (type === 'theme') {
			setButtonsState(true)
			if (themeState === 0) {
				setThemeState(1)
			} else {
				setThemeState(0)
			}
		}

		//buttonsState ? setIrrigationDurationStateClass('var(--inactiveGrey)') : setIrrigationDurationStateClass('#000000')
	}

	const handleCancelButton = () => {
		setButtonsState(false)
		// Will throw error because API is not accessible
		setAutomaticIrrigationState(settingsData.limitsTrigger)
		setAutomaticIrrigationStateClass(settingsData.limitsTrigger ? '#000000' : 'var(--inactiveGrey)')
		setScheduledIrrigationState(settingsData.scheduledTrigger)
		setScheduledIrrigationStateClass(settingsData.scheduledTrigger ? '#000000' : 'var(--inactiveGrey)')
		setIrrigationDuration(settingsData.irrigationDuration)
		setIrrigationDurationStateClass(buttonsState ? '#000000' : 'var(--inactiveGrey)')
		setMoistureLimit(settingsData.moistureLimit)
		setWaterAmountLimit(settingsData.waterAmountLimit)
		setWaterLevelLimit(settingsData.waterLevelLimit)
		setHourRange(settingsData.hourRange)
		setChartTypeState(settingsData.chartType)
		setLanguageState(settingsData.language)
		setThemeState(settingsData.theme)
		setGetCoordsState(false)
		setGetCoords('')
		setLocation(settingsData.location)
		setLatitude(settingsData.lat)
		setLongitude(settingsData.lon)
	}

	return (
		<div className="settings">
			<Card className="card">
				<CardContent>
					<div className="flex-row">
						<div className="flex-col">
							<div className="flex-row pt-2 title-2">
								<span className="title-1">Nastavení zavlažování</span>
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
							<div className="flex-row">
								<div className="flex-col">
									<div className="flex-row pt-2">
										<span style={{ color: irrigationDurationStateClass }}>Doba zavlažování (s): </span>
									</div>
									<div className="flex-row pt-2">
										<span style={{ color: automaticIrrigationStateClass }}>Limit vlhkosti půdy (%): </span>
									</div>
									<div className="flex-row pt-2">
										<span style={{ color: automaticIrrigationStateClass }}>Limit přečerpané vody (l): </span>
									</div>
									<div className="flex-row pt-2">
										<span style={{ color: automaticIrrigationStateClass }}>Limit hladiny vody (cm): </span>
									</div>
									<div className="flex-row pt-2">
										<span style={{ color: scheduledIrrigationStateClass }}>Rozsah hodin (h): </span>
									</div>
								</div>
								<div className="flex-col ml-3">
									<div
										className="flex-row pt-1"
										onBlur={(data: any) => updateInputData('irrigationDuration', data.target.value)}
										onChange={(data: any) => data.target.value == 0 && setButtonsState(false)}
									>
										<EditableField
											key="irrigationDuration"
											defaultValue={irrigationDuration}
											active={irrigationDurationStateClass === '#000000'}
											width="10"
										/>
									</div>
									<div
										className="flex-row pt-1"
										onBlur={(data: any) => updateInputData('moistureLimit', data.target.value)}
									>
										<EditableField
											key="moistureLimit"
											defaultValue={moistureLimit}
											active={automaticIrrigationState}
											width="10"
										/>
									</div>
									<div
										className="flex-row pt-1"
										onBlur={(data: any) => updateInputData('waterAmountLimit', data.target.value)}
									>
										<EditableField
											key="waterAmountLimit"
											defaultValue={waterAmountLimit}
											active={automaticIrrigationState}
											width="10"
										/>
									</div>
									<div
										className="flex-row pt-1"
										onBlur={(data: any) => updateInputData('waterLevelLimit', data.target.value)}
									>
										<EditableField
											key="waterLevelLimit"
											defaultValue={waterLevelLimit}
											active={automaticIrrigationState}
											width="10"
										/>
									</div>
									<div
										className="flex-row pt-1"
										onBlur={(data: any) => updateInputData('hourRange', data.target.value)}
									>
										<EditableField
											key="hourRange"
											defaultValue={hourRange}
											active={scheduledIrrigationState}
											width="10"
										/>
									</div>
								</div>
							</div>
						</div>
					</div>
					<div className="flex-row">
						<div className="flex-col">
							<div className="flex-row pt-2 title-2">
								<span className="title-1">Nastavení aplikace</span>
							</div>
							<div className="flex-row">
								<div className="flex-col">
									<div className="flex-row pt-2">
										<span>Typ grafu: </span>
									</div>
									<div className="flex-row pt-3">
										<span>Jazyk: </span>
									</div>
									<div className="flex-row pt-3">
										<span>Motiv: </span>
									</div>
									<div className="flex-row pt-3">
										<span>Lokace: </span>
									</div>
								</div>
								<div className="flex-col ml-3">
									<div className="flex-row pt-1">
										<div onClick={() => updateToggleState('chartType')}>
											<ToggleButton
												key="chartType"
												toggleState={chartTypeState}
												values={[{ label: 'lineGraph' }, { label: 'barGraph' }]}
											/>
										</div>
									</div>
									<div className="flex-row pt-1">
										<div onClick={() => updateToggleState('language')}>
											<ToggleButton
												key="language"
												toggleState={languageState}
												values={[{ label: 'flagCZ' }, { label: 'flagGB' }]}
											/>
										</div>
									</div>
									<div className="flex-row pt-1">
										<div onClick={() => updateToggleState('theme')}>
											<ToggleButton
												key="theme"
												toggleState={themeState}
												values={[{ label: 'lightTheme' }, { label: 'darkTheme' }]}
											/>
										</div>
									</div>
									<div
										className="flex-row pt-1"
										// tohle je na prasáka ale funguje to
										onChange={(data: any) => updateInputData('location', data)}
										onBlur={(data: any) => updateInputData('location', data)}
									>
										<EditableField key="city" defaultValue={location} active="true" width="40" />
									</div>
								</div>
							</div>
						</div>
					</div>
					<div className="flex-row mt-5">
						<div className="flex-col">
							<div onClick={() => buttonsState && handleCancelButton()}>
								<CancelButton active={buttonsState} />
							</div>
						</div>
						<div className="flex-col ml-3">
							<div
								onClick={() =>
									buttonsState &&
									updateSettingsData({
										variables: {
											limitsTrigger: automaticIrrigationState,
											waterLevelLimit: waterLevelLimit,
											waterAmountLimit: waterAmountLimit,
											moistureLimit: moistureLimit,
											scheduledTrigger: scheduledIrrigationState,
											hourRange: hourRange,
											chartType: chartTypeState,
											theme: themeState,
											language: languageState,
											location: location,
											lat: latitude,
											lon: longitude,
										},
									})
								}
							>
								<SaveButton active={buttonsState} />
							</div>
						</div>
					</div>
				</CardContent>
			</Card>
		</div>
	)
}
