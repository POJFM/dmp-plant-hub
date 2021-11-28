import { useEffect, useState } from 'react'
import Card from '@material-ui/core/Card'
import CardContent from '@material-ui/core/CardContent'
import TextInputField from './fields/TextInputField'
import SaveButton from './buttons/SaveButton'

export default function InitForm() {
	const [currentMoistureLevel, setCurrentMoistureLevel] = useState(0)
	const [currentWaterLevel, setCurrentWaterLevel] = useState(0)
	const [currentWaterOverdrawnLevel, setCurrentWaterOverdrawnLevel] = useState(0)
	const [currentLocation, setCurrentLocation] = useState(0)
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

	return (
		<div className="init-form">
			<Card className="card w-400px h-500px">
				<CardContent>
					<div className="flex-col p-3">
						<div className="flex-row flex-center p-1 pt-5px mt-2">
							<span className="title-1">Inicializační sekvence</span>
						</div>
						<div className="flex-row p-1 pt-5px mt-2">
							<TextInputField key="moistureLevelLimit" name="Vlhkostní limit (%)" defaultValue={currentMoistureLevel} />
						</div>
						<div className="flex-row p-1 pt-5px mt-2">
							<TextInputField key="waterLevelLimit" name="Limit hladiny vody (cm)" defaultValue={currentWaterLevel} />
						</div>
						<div className="flex-row p-1 pt-5px mt-2">
							<TextInputField
								key="waterOverdrawnLevelLimit"
								name="Limit přečerpané vody (l)"
								defaultValue={currentWaterOverdrawnLevel}
							/>
						</div>
						<div className="flex-row p-1 pt-5px mt-2">
							<TextInputField key="location" name="Lokace" defaultValue={currentLocation} />
						</div>
						<div className="flex-row-reverse p-1 pt-5px mt-2">
							<SaveButton /* onClick={() => saveValues()} */ />
						</div>
					</div>
				</CardContent>
			</Card>
		</div>
	)
}
