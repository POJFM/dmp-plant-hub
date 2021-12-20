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
import { updateSettings } from '../graphql/mutations'
//import { fetchCoordsFromLocation } from 'src/utils'

export default function Settings() {
	const { loading: settingsLoading, error: settingsError, data: settingsData } = useQuery(settings)
	const [updateSettingsData, { data, loading, error }] = useMutation(updateSettings)

	const [buttonsState, setButtonsState] = useState(false), // false
		[automaticIrrigationState, setAutomaticIrrigationState] = useState(true), // settingsData.limitsTrigger
		[automaticIrrigationStateClass, setAutomaticIrrigationStateClass] = useState<string>(), // settingsData.limitsTrigger ? '#000000' : 'var(--inactiveGrey)'
		[scheduledIrrigationState, setScheduledIrrigationState] = useState(false), // settingsData.scheduledTrigger
		[scheduledIrrigationStateClass, setScheduledIrrigationStateClass] = useState<string>(), // settingsData.scheduledTrigger ? '#000000' : 'var(--inactiveGrey)'
		[moistureLimit, setMoistureLimit] = useState<number>(50), // settingsData.moistureLimit
		[waterAmountLimit, setWaterAmountLimit] = useState<number>(3), // settingsData.waterAmountLimit
		[waterLevelLimit, setWaterLevelLimit] = useState<number>(6), // settingsData.waterLevelLimit
		[hoursRange, setHoursRange] = useState<number>(10), // settingsData.hoursRange
		[chartTypeState, setChartTypeState] = useState(1), // settingsData.chartType
		[languageState, setLanguageState] = useState(1), // settingsData.language
		[themeState, setThemeState] = useState(1), // settingsData.theme
		[getCoordsState, setGetCoordsState] = useState(false),
		[getCoords, setGetCoords] = useState<string>(),
		[location, setLocation] = useState<string>('Frýdek-Místek'), // settingsData.location
		[latitude, setLatitude] = useState<number>(), // settingsData.lat
		[longitude, setLongitude] = useState<number>() // settingsData.lon

	useEffect(() => {
		document.title = 'Plant Hub | Settings'
	}, [])

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
		type === 'moistureLimit' && setMoistureLimit(data?.target?.value)
		type === 'waterAmountLimit' && setWaterAmountLimit(data?.target?.value)
		type === 'waterLevelLimit' && setWaterLevelLimit(data?.target?.value)
		type === 'hourRange' && setHoursRange(data?.target?.value)
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
		setButtonsState(true)

		if (type === 'automaticIrrigation') {
			if (automaticIrrigationState === false) {
				setAutomaticIrrigationState(true)
				setButtonsState(true)
				setAutomaticIrrigationStateClass('#000000')
			} else {
				setAutomaticIrrigationState(false)
				setAutomaticIrrigationStateClass('var(--inactiveGrey)')
				if (scheduledIrrigationState === false) {
					setButtonsState(false)
				}
			}
		}

		if (type === 'scheduledIrrigation') {
			if (scheduledIrrigationState === false) {
				setScheduledIrrigationState(true)
				setButtonsState(true)
				setScheduledIrrigationStateClass('#000000')
			} else {
				setScheduledIrrigationState(false)
				setScheduledIrrigationStateClass('var(--inactiveGrey)')
				if (automaticIrrigationState === false) {
					setButtonsState(false)
				}
			}
		}

		if (type === 'chartType') {
			if (chartTypeState === 0) {
				setChartTypeState(1)
			} else {
				setChartTypeState(0)
			}
		}

		if (type === 'language') {
			if (languageState === 0) {
				setLanguageState(1)
			} else {
				setLanguageState(0)
			}
		}

		if (type === 'theme') {
			if (themeState === 0) {
				setThemeState(1)
			} else {
				setThemeState(0)
			}
		}
	}

	const handleCancelButton = () => {
		setButtonsState(false)
		// Will throw error because API is not accessible
		setAutomaticIrrigationState(settingsData.limitsTrigger)
		setAutomaticIrrigationStateClass(settingsData.limitsTrigger ? '#000000' : 'var(--inactiveGrey)')
		setScheduledIrrigationState(settingsData.scheduledTrigger)
		setScheduledIrrigationStateClass(settingsData.scheduledTrigger ? '#000000' : 'var(--inactiveGrey)')
		setMoistureLimit(settingsData.moistureLimit)
		setWaterAmountLimit(settingsData.waterAmountLimit)
		setWaterLevelLimit(settingsData.waterLevelLimit)
		setHoursRange(settingsData.hoursRange)
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
											defaultValue={hoursRange}
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
									<div className="flex-row pt-2">
										<span>Jazyk: </span>
									</div>
									<div className="flex-row pt-2">
										<span>Motiv: </span>
									</div>
									<div className="flex-row pt-2">
										<span>Lokace: </span>
									</div>
								</div>
								<div className="flex-col ml-3">
									<div className="flex-row pt-1">
										<div onClick={() => updateToggleState('chartType')}>
											<ToggleButton
												key="chartType"
												toggleState={chartTypeState}
												values={[{ label: 'Spojnicový' }, { label: 'Sloupcový' }]}
											/>
										</div>
									</div>
									<div className="flex-row pt-1">
										<div onClick={() => updateToggleState('language')}>
											<ToggleButton
												key="language"
												toggleState={languageState}
												values={[{ label: 'Česky' }, { label: 'Anglicky' }]}
											/>
										</div>
									</div>
									<div className="flex-row pt-1">
										<div onClick={() => updateToggleState('theme')}>
											<ToggleButton
												key="theme"
												toggleState={themeState}
												values={[{ label: 'Světlý' }, { label: 'Tmavý' }]}
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
											hoursRange: hoursRange,
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
