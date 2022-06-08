import { useEffect, useState, useMemo } from 'react'
import axios from 'axios'
import { useTranslation } from 'react-i18next'
import { MuiCard as Card } from './Card'
import EditableField from './fields/EditableField'
import ToggleButton from './buttons/ToggleButton'
import SaveButton from './buttons/SaveButton'
import CancelButton from './buttons/CancelButton'
import { useQuery, useMutation } from '@apollo/client'
import { settings } from '../graphql/queries'
import { updateSettingsMut } from '../graphql/mutations'

export default function Settings() {
	const { t, i18n } = useTranslation()
	const { data: settingsData } = useQuery(settings),
		[buttonsState, setButtonsState] = useState(false),
		[automaticIrrigationState, setAutomaticIrrigationState] = useState(
			settingsData?.getSettings[0]?.limits_trigger || true
		),
		[automaticIrrigationStateClass, setAutomaticIrrigationStateClass] = useState<string>(
			settingsData?.getSettings[0]?.limits_trigger ? '#000000' : 'var(--inactiveGrey)'
		),
		[scheduledIrrigationState, setScheduledIrrigationState] = useState(settingsData?.getSettings[0]?.scheduled_trigger),
		[irrigationDuration, setIrrigationDuration] = useState<number>(settingsData?.getSettings[0]?.irrigation_duration),
		[irrigationDurationStateClass, setIrrigationDurationStateClass] = useState('#000000'),
		[defaultWaterAmount, setDefaultWaterAmount] = useState<number>(settingsData?.getSettings[0]?.default_water_amount),
		[defaultWaterAmountStateClass, setDefaultWaterAmountStateClass] = useState('#000000'),
		[scheduledIrrigationStateClass, setScheduledIrrigationStateClass] = useState<string>(
			settingsData?.getSettings[0]?.scheduled_trigger ? '#000000' : 'var(--inactiveGrey)'
		),
		[moistLimit, setMoistLimit] = useState(settingsData?.getSettings[0]?.moist_limit),
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

	useEffect(() => {
		document.title = 'Plant Hub | Settings'
	}, [])

	const updateDataStates = () => {
		setAutomaticIrrigationState(settingsData?.getSettings[0]?.limits_trigger)
		setAutomaticIrrigationStateClass(settingsData?.getSettings[0]?.limits_trigger ? '#000000' : 'var(--inactiveGrey)')
		setScheduledIrrigationState(settingsData?.getSettings[0]?.scheduled_trigger)
		setScheduledIrrigationStateClass(settingsData?.getSettings[0]?.scheduled_trigger ? '#000000' : 'var(--inactiveGrey)')

		setIrrigationDuration(settingsData?.getSettings[0]?.irrigation_duration)
		setDefaultWaterAmount(settingsData?.getSettings[0]?.default_water_amount)
		setMoistLimit(settingsData?.getSettings[0]?.moist_limit)
		setWaterAmountLimit(settingsData?.getSettings[0]?.water_amount_limit)
		setWaterLevelLimit(settingsData?.getSettings[0]?.water_level_limit)
		setHourRange(settingsData?.getSettings[0]?.hour_range)

		setChartTypeState(settingsData?.getSettings[0]?.chart_type)
		setLanguageState(settingsData?.getSettings[0]?.language)
		setThemeState(settingsData?.getSettings[0]?.theme)

		setLocation(settingsData?.getSettings[0]?.location)
		setLatitude(settingsData?.getSettings[0]?.lat)
		setLongitude(settingsData?.getSettings[0]?.lon)
	}

	useEffect(() => {
		updateDataStates()
	}, [settingsData])

	const [createSettings] = useMutation(updateSettingsMut, {
		variables: {
			limits_trigger: automaticIrrigationState,
			water_level_limit: waterLevelLimit,
			water_amount_limit: waterAmountLimit,
			moist_limit: moistLimit,
			scheduled_trigger: scheduledIrrigationState,
			hour_range: hourRange,
			location: location,
			irrigation_duration: irrigationDuration,
			default_water_amount: defaultWaterAmount,
			chart_type: chartTypeState,
			language: languageState,
			theme: themeState,
			lat: latitude,
			lon: longitude,
		},
		refetchQueries: [{ query: settings }],
	})

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
		switch (type) {
			case 'irrigationDuration': setIrrigationDuration(data); break
			case 'defaultWaterAmount': setDefaultWaterAmount(data); break
			case 'moistLimit': setMoistLimit(data); break
			case 'waterAmountLimit': setWaterAmountLimit(data); break
			case 'waterLevelLimit': setWaterLevelLimit(data); break
			case 'hourRange': setHourRange(data); break
			case 'location': {
				setGetCoords(data)
				setGetCoordsState(true)
				break
			}
		}
	}

	useEffect(() => {
		if (getCoordsState) {
			fetchCoordsFromLocation(getCoords)
			setGetCoordsState(false)
		}
	}, [getCoordsState])

	useEffect(() => {
		if (themeState === true) {
			document.body.classList.add('dark')
			document.documentElement.setAttribute("theme", "dark")
			localStorage.setItem("theme", "dark")
		} else {
			document.body.classList.remove('dark')
			document.documentElement.setAttribute("theme", "light")
			localStorage.setItem("theme", "light")
		}
	}, [themeState])

	const updateToggleState = (type: string) => {
		switch (type) {
			case 'automaticIrrigation': {
				if (automaticIrrigationState === false) {
					setAutomaticIrrigationState(true)
					setButtonsState(true)
					setIrrigationDurationStateClass('#000000')
					setDefaultWaterAmountStateClass('#000000')
					setAutomaticIrrigationStateClass('#000000')
				} else {
					setAutomaticIrrigationState(false)
					setAutomaticIrrigationStateClass('var(--inactiveGrey)')
					if (scheduledIrrigationState === false) {
						setButtonsState(false)
						setIrrigationDurationStateClass('var(--inactiveGrey)')
						setDefaultWaterAmountStateClass('var(--inactiveGrey)')
					}
				}
				break
			}
			case 'scheduledIrrigation': {
				if (scheduledIrrigationState === false) {
					setScheduledIrrigationState(true)
					setButtonsState(true)
					setIrrigationDurationStateClass('#000000')
					setDefaultWaterAmountStateClass('#000000')
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
						setDefaultWaterAmountStateClass('var(--inactiveGrey)')
					}
				}
				break
			}
			case 'chartType': {
				setButtonsState(true)
				if (!chartTypeState) {
					setChartTypeState(true)
				} else {
					setChartTypeState(false)
				}
				break
			}
			case 'language': {
				setButtonsState(true)
				if (!languageState) {
					setLanguageState(true)
					i18n.changeLanguage("en")
				} else {
					setLanguageState(false)
					i18n.changeLanguage("cs")
				}
				break
			}
			case 'theme': {
				setButtonsState(true)
				if (!themeState) {
					setThemeState(true)
				} else {
					setThemeState(false)
				}
				break
			}
		}
	}

	const handleCancelButton = () => {
		updateDataStates()
		setButtonsState(false)
		setIrrigationDurationStateClass(buttonsState ? '#000000' : 'var(--inactiveGrey)')
		setDefaultWaterAmountStateClass(buttonsState ? '#000000' : 'var(--inactiveGrey)')
		setGetCoordsState(false)
		setGetCoords('')
	}

	return (
		<div className="settings">
			<Card>
				<div className="flex-row">
					<div className="flex-col">
						<div className="flex-row pt-2 title-2">
							<span className="title-1">{t('settings.irrigationSettings')}</span>
						</div>
						<div className="flex-row mb-2">
							<div className="flex-col p-1 pt-5px mt-2">
								<div className="flex-row">
									<span className="title-2">{t('settings.automatic')}</span>
								</div>
								<div className="flex-row mt-2">
									<span className="title-2">{t('settings.scheduled')}</span>
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
									<span style={{ color: irrigationDurationStateClass }}>{t('settings.irrigationDuration')} (s): </span>
								</div>
								<div className="flex-row pt-2">
									<span style={{ color: irrigationDurationStateClass }}>{t('settings.defaultWaterAmount')} (l): </span>
								</div>
								<div className="flex-row pt-2">
									<span style={{ color: automaticIrrigationStateClass }}>{t('settings.moistLimit')} (%): </span>
								</div>
								<div className="flex-row pt-2">
									<span style={{ color: automaticIrrigationStateClass }}>{t('settings.waterAmountLimit')} (l): </span>
								</div>
								<div className="flex-row pt-2">
									<span style={{ color: automaticIrrigationStateClass }}>{t('settings.waterLevelLimit')} (cm): </span>
								</div>
								<div className="flex-row pt-2">
									<span style={{ color: scheduledIrrigationStateClass }}>{t('settings.hourRange')} (h): </span>
								</div>
							</div>
							<div className="flex-col ml-3">
								<div
									className="flex-row pt-1"
									onChange={(data: any) => {
										updateInputData('irrigationDuration', parseInt(data.target.value))
										data.target.value == 0 && setButtonsState(false)
									}}
								>
									<EditableField
										name="irrigationDuration"
										defaultValue={irrigationDuration}
										active={irrigationDurationStateClass === '#000000'}
										width={10}
										dataType="number"
									/>
								</div>
								<div
									className="flex-row pt-1"
									onChange={(data: any) => {
										updateInputData('defaultWaterAmount', data.target.value)
										data.target.value == 0 && setButtonsState(false)
									}}
								>
									<EditableField
										name="defaultWaterAmount"
										defaultValue={defaultWaterAmount}
										active={defaultWaterAmountStateClass === '#000000'}
										width={10}
										dataType="number"
									/>
								</div>
								<div
									className="flex-row pt-1"
									onChange={(data: any) => updateInputData('moistLimit', data.target.value)}
								>
									<EditableField
										name="moistLimit"
										defaultValue={moistLimit}
										active={automaticIrrigationState}
										width={10}
										dataType="number"
									/>
								</div>
								<div
									className="flex-row pt-1"
									onChange={(data: any) => updateInputData('waterAmountLimit', data.target.value)}
								>
									<EditableField
										name="waterAmountLimit"
										defaultValue={waterAmountLimit}
										active={automaticIrrigationState}
										width={10}
										dataType="number"
									/>
								</div>
								<div
									className="flex-row pt-1"
									onChange={(data: any) => updateInputData('waterLevelLimit', data.target.value)}
								>
									<EditableField
										name="waterLevelLimit"
										defaultValue={waterLevelLimit}
										active={automaticIrrigationState}
										width={10}
										dataType="number"
									/>
								</div>
								<div
									className="flex-row pt-1"
									onChange={(data: any) => updateInputData('hourRange', data.target.value)}
								>
									<EditableField
										name="hourRange"
										defaultValue={hourRange}
										active={scheduledIrrigationState}
										width={10}
										dataType="number"
									/>
								</div>
							</div>
						</div>
					</div>
				</div>
				<div className="flex-row">
					<div className="flex-col">
						<div className="flex-row pt-2 title-2">
							<span className="title-1">{t('settings.applicationSettings')}</span>
						</div>
						<div className="flex-row">
							<div className="flex-col">
								<div className="flex-row pt-2">
									<span>{t('settings.chartType')}: </span>
								</div>
								<div className="flex-row pt-3">
									<span>{t('settings.language')}: </span>
								</div>
								<div className="flex-row pt-3">
									<span>{t('settings.theme')}: </span>
								</div>
								<div className="flex-row pt-3">
									<span>{t('settings.location')}: </span>
								</div>
							</div>
							<div className="flex-col ml-3">
								<div className="flex-row pt-1">
									<div onClick={() => updateToggleState('chartType')}>
										<ToggleButton
											item="chartType"
											toggleState={chartTypeState}
											values={[{ label: 'lineGraph' }, { label: 'barGraph' }]}
										/>
									</div>
								</div>
								<div className="flex-row pt-1">
									<div onClick={() => updateToggleState('language')}>
										<ToggleButton
											item="language"
											toggleState={languageState}
											values={[{ label: 'flagCZ' }, { label: 'flagGB' }]}
										/>
									</div>
								</div>
								<div className="flex-row pt-1">
									<div onClick={() => updateToggleState('theme')}>
										<ToggleButton
											item="theme"
											toggleState={themeState}
											values={[{ label: 'lightTheme' }, { label: 'darkTheme' }]}
										/>
									</div>
								</div>
								<div
									className="flex-row pt-1"
									onChange={(data: any) => updateInputData('location', data)}
								>
									<EditableField
										name="city"
										defaultValue={location}
										active={true}
										width={40}
										dataType="string"
									/>
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
						<div onClick={() => buttonsState && createSettings()}>
							<SaveButton active={buttonsState} />
						</div>
					</div>
				</div>
			</Card>
		</div >
	)
}
