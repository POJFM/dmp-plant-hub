import { useEffect } from 'react'
import Card from '@material-ui/core/Card'
import { useStyles } from './../styles/rootStyles'
import { useSettingsStyles } from './../styles/settings'
import CardContent from '@material-ui/core/CardContent'
import EditableField from './fields/EditableField'
import EnumEditableField from './fields/EnumEditableField'

export default function Settings() {
	const classes = useStyles()
	const settingsClasses = useSettingsStyles()
	useEffect(() => {
		document.title = 'Plant Hub | Settings'
	}, [])

	return (
		<div className="col settings">
			<Card className={classes.card}>
				<CardContent>
					<div className="row">
						<div className="col">
							<div className={`row ${classes.cardRow} ${classes.cardRowTitle}`}>
								<span>Nastavení zavlažování</span>
							</div>
							<div className="row">
								<div className="col">
									<div className={`row ${classes.cardRow}`}>
										<span>
											Limit vlhkosti půdy: <EditableField key="moistureLimit" defaultValue="0" />
										</span>
									</div>
									<div className={`row ${classes.cardRow}`}>
										<span>
											Limit přečerpané vody: <EditableField key="waterOverdrawnLimit" defaultValue="5" />
										</span>
									</div>
									<div className={`row ${classes.cardRow}`}>
										<span>
											Limit hladiny vody: <EditableField key="waterLevelLimit" defaultValue="55" />
										</span>
									</div>
									<div className={`row ${classes.cardRow}`}>
										<span>Nastavit čas automatického zavlažování: </span>
									</div>
								</div>
							</div>
						</div>
					</div>
					<div className="row">
						<div className="col">
							<div className={`row ${classes.cardRow} ${classes.cardRowTitle}`}>
								<span>Nastavení aplikace</span>
							</div>
							<div className="row">
								<div className="col">
									<div className={`row ${classes.cardRow}`}>
										<span>Jazyk: </span>
										{/* <EnumEditableField key="theme" values={[{ label: 'Česky' }, { label: 'Anglicky' }]} /> */}
									</div>
									<div className={`row ${classes.cardRow}`}>
										<span>Motiv: </span>
										{/* <EnumEditableField key="theme" values={[{ label: 'Světlý' }, { label: 'Tmavý' }]} /> */}
									</div>
									<div className={`row ${classes.cardRow}`}>
										<span>
											Lokace: <EditableField key="city" defaultValue="Frýdek-Místek" />
										</span>
									</div>
								</div>
							</div>
						</div>
					</div>
				</CardContent>
			</Card>
		</div>
	)
}
