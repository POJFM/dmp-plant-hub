import { useEffect, useState } from 'react'
import Card from '@material-ui/core/Card'
import CardContent from '@material-ui/core/CardContent'
import TextInputField from './fields/TextInputField'
import ToggleButton from './buttons/ToggleButton'
import SaveButton from './buttons/SaveButton'

export default function InitForm() {
	const [currentMoistureLevel, setCurrentMoistureLevel] = useState(0)
	const [currentWaterLevel, setCurrentWaterLevel] = useState(0)
	const [currentWaterOverdrawnLevel, setCurrentWaterOverdrawnLevel] = useState(0)
	const [currentLocation, setCurrentLocation] = useState('Frýdek-Místek')
	const [saveButtonState, setSaveButtonState] = useState('true')
	const [automaticIrrigationState, setAutomaticIrrigationState] = useState('true')
	const [scheduledIrrigationState, setScheduledIrrigationState] = useState('false')

	useEffect(() => {
		document.title = 'Plant Hub | Initialization'

		fetch('/init/measured')
			.then((res) => res.json())
			.then((data) => {
				setCurrentMoistureLevel(data.moistureLevel)
				setCurrentWaterLevel(data.waterLevel)
				setCurrentWaterOverdrawnLevel(data.waterOverdrawnLevel)
				setCurrentLocation(data.location)
			})
	}, [])

	const updateToggleState = (type: string) => {
		if (type === 'automaticIrrigation') {
			if (automaticIrrigationState === 'false') {
				setAutomaticIrrigationState('true')
				setSaveButtonState('true')
			} else {
				setAutomaticIrrigationState('false')
				if (scheduledIrrigationState === 'false') {
					setSaveButtonState('false')
				}
			}
		}

		if (type === 'scheduledIrrigation') {
			if (scheduledIrrigationState === 'false') {
				setScheduledIrrigationState('true')
				setSaveButtonState('true')
			} else {
				setScheduledIrrigationState('false')
				if (automaticIrrigationState === 'false') {
					setSaveButtonState('false')
				}
			}
		}
	}

	return (
		<div className="init-form">
			<Card className="card w-400px h-500px">
				<CardContent>
					<div className="flex-col p-3">
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
										<ToggleButton key="limitsTrigger" toggleState={automaticIrrigationState} />
									</div>
								</div>
								<div className="flex-row mt-2">
									<div onClick={() => updateToggleState('scheduledIrrigation')}>
										<ToggleButton key="scheduledTrigger" toggleState={scheduledIrrigationState} />
									</div>
								</div>
							</div>
						</div>
						<div className="flex-row p-1 pt-5px mt-2">
							<TextInputField
								key="moistureLevelLimit"
								name="Vlhkostní limit (%)"
								defaultValue={currentMoistureLevel}
								active={automaticIrrigationState}
							/>
						</div>
						<div className="flex-row p-1 pt-5px mt-2">
							<TextInputField
								key="moistureLevelLimit"
								name="Vlhkostní limit (%)"
								defaultValue={currentMoistureLevel}
								active={automaticIrrigationState}
							/>
						</div>
						<div className="flex-row p-1 pt-5px mt-2">
							<TextInputField
								key="waterLevelLimit"
								name="Limit hladiny vody (cm)"
								defaultValue={currentWaterLevel}
								active={automaticIrrigationState}
							/>
						</div>
						<div className="flex-row p-1 pt-5px mt-2">
							<TextInputField
								key="waterOverdrawnLevelLimit"
								name="Limit přečerpané vody (l)"
								defaultValue={currentWaterOverdrawnLevel}
								active={automaticIrrigationState}
							/>
						</div>
						<div className="flex-row p-1 pt-5px mt-2">
							<TextInputField key="hourRange" name="Rozsah hodin (h)" active={scheduledIrrigationState} />
						</div>
						<div className="flex-row p-1 pt-5px mt-2">
							<TextInputField key="location" name="Lokace" defaultValue={currentLocation} active="true" />
						</div>
						<div className="flex-row-reverse p-1 pt-5px mt-2">
							<SaveButton /* onClick={() => saveValues()} */ active={saveButtonState} />
						</div>
					</div>
				</CardContent>
			</Card>
		</div>
	)
}
