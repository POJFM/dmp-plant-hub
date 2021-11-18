import { useEffect, useState } from 'react'
import { useStyles } from './../styles/rootStyles'
import Card from '@material-ui/core/Card'
import CardContent from '@material-ui/core/CardContent'
import TextInputField from './fields/TextInputField'
import SaveButton from './buttons/SaveButton'

export default function InitForm() {
	const classes = useStyles()
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
		<div className={classes.initForm}>
			<Card className={`${classes.card} ${classes.initFormCard}`}>
				<CardContent>
					<div className="row">
						<div className="col">
							<div className={`row ${classes.cardRow} ${classes.initFormCardRow} ${classes.initFormCardRowFirst}`}>
								<TextInputField
									key="moistureLevelLimit"
									name="Vlhkostní limit (%)"
									defaultValue={currentMoistureLevel}
								/>
							</div>
							<div className={`row ${classes.cardRow} ${classes.initFormCardRow}`}>
								<TextInputField key="waterLevelLimit" name="Limit hladiny vody (cm)" defaultValue={currentWaterLevel} />
							</div>
							<div className={`row ${classes.cardRow} ${classes.initFormCardRow}`}>
								<TextInputField
									key="waterOverdrawnLevelLimit"
									name="Limit přečerpané vody (l)"
									defaultValue={currentWaterOverdrawnLevel}
								/>
							</div>
							<div className={`row ${classes.cardRow} ${classes.initFormCardRow}`}>
								<TextInputField key="location" name="Lokace" defaultValue={currentLocation} />
							</div>
							<div className={`row ${classes.cardRow} ${classes.initFormCardRow}`}>
								<SaveButton /* onClick={() => saveValues()} */ />
							</div>
						</div>
					</div>
				</CardContent>
			</Card>
		</div>
	)
}
