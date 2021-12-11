import { useEffect, useState } from 'react'
//import { settings } from '../graphql/queries'
import Card from '@material-ui/core/Card'
import CardContent from '@material-ui/core/CardContent'
import EditableField from './fields/EditableField'
import EnumEditableField from './fields/EnumEditableField'
import ToggleButton from './buttons/ToggleButton'

export default function Settings() {
	// get state based on data saved in DB
	const [chartTypeState, setChartTypeState] = useState('true')
	const [languageState, setLanguageState] = useState('true')
	const [themeState, setThemeState] = useState('true')

	useEffect(() => {
		document.title = 'Plant Hub | Settings'
	}, [])

	const updateToggleState = (type: string) => {
		if (type === 'chartType') {
			if (chartTypeState === 'false') {
				setChartTypeState('true')
			} else {
				setChartTypeState('false')
			}
		}

		if (type === 'language') {
			if (languageState === 'false') {
				setLanguageState('true')
			} else {
				setLanguageState('false')
			}
		}

		if (type === 'theme') {
			if (themeState === 'false') {
				setThemeState('true')
			} else {
				setThemeState('false')
			}
		}
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
