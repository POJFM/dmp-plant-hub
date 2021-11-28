import { useEffect } from 'react'
import Card from '@material-ui/core/Card'
import CardContent from '@material-ui/core/CardContent'
import EditableField from './fields/EditableField'
import EnumEditableField from './fields/EnumEditableField'

export default function Settings() {
	useEffect(() => {
		document.title = 'Plant Hub | Settings'
	}, [])

	return (
		<div className="settings">
			<Card className="card">
				<CardContent>
					<div className="flex-row">
						<div className="flex-col">
							<div className="flex-row pt-2 title-2">
								<span className="title-1">Nastavení zavlažování</span>
							</div>
							<span>!checkbox pro auto i scheduled zavlažování!</span>
							<div className="flex-row">
								<div className="flex-col">
									<div className="flex-row pt-2">
										<span>Limit vlhkosti půdy (%): </span>
									</div>
									<div className="flex-row pt-2">
										<span>Limit přečerpané vody (l): </span>
									</div>
									<div className="flex-row pt-2">
										<span>Limit hladiny vody (cm): </span>
									</div>
								</div>
								<div className="flex-col ml-3">
									<div className="flex-row pt-1">
										<EditableField key="moistureLimit" defaultValue="0" />
									</div>
									<div className="flex-row pt-1">
										<EditableField key="waterOverdrawnLimit" defaultValue="5" />
									</div>
									<div className="flex-row pt-1">
										<EditableField key="waterLevelLimit" defaultValue="55" />
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
										<EnumEditableField key="theme" values={[{ label: 'Česky' }, { label: 'Anglicky' }]} />
									</div>
									<div className="flex-row pt-1">
										<EnumEditableField key="theme" values={[{ label: 'Světlý' }, { label: 'Tmavý' }]} />
									</div>
									<div className="flex-row pt-1">
										<EditableField key="city" defaultValue="Frýdek-Místek" />
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
