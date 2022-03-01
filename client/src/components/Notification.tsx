import { useEffect, useState } from 'react'
import axios from 'axios'
import Card from '@material-ui/core/Card'
import CardContent from '@material-ui/core/CardContent'
import Loading from './states/Loading'
import Finished from './states/Finished'
import Warning from './states/Warning'

export default function Notification(props: any) {
	const [title, setTitle] = useState<string>(), // TEST => 'Zavlažování'
		[notificationClass, setNotificationClass] = useState('hidden'),
		[notificationStateClass, setNotificationStateClass] = useState<string>(), // TEST => 'var(--irrigationBlue)'
		[state, setState] = useState<string>(), // TEST => 'inProgress'
		[action, setAction] = useState<any>(), // TEST => 'Probíhá zavlažování'
		[notify, setNotify] = useState(false)

	let getNotificationsInterval: any

	useEffect(() => {
		!getNotificationsInterval && (getNotificationsInterval = setInterval(() => getNotifications(), 2000))
	}, [])

	const getNotifications = () => {
		axios
			.request({
				method: 'GET',
				url: `${process.env.REACT_APP_GO_REST_API_URL}/live/notify`,
				headers: {
					'Content-Type': 'application/json',
				},
			})
			.then((res) => {
				if (res.data.state !== 'inactive') {
					setTitle(res.data.title)
					setState(res.data.state)
					setAction(res.data.action)
					setNotify(true)
				}
			})
			.catch((error) => {
				console.error(error)
			})
	}

	// useEffect(() => {
	// 	setNotificationClass('notification')
	// 	if (state === 'physicalHelpRequired') {
	// 		setNotificationStateClass('var(--warningRed)')
	// 		setNotificationActionClass('var(--warningRed)')
	// 	}
	// 	state === 'inProgress' && setNotificationActionClass('var(--irrigationBlue)')
	// }, [notify])

	// TEST KOKOT
	// const setKokot = () => {
	// 	setState('finished')
	// 	setAction('Zavlažování dokončeno')
	// 	setNotificationStateClass('var(--green)')
	// }

	// const setKokot2 = () => {
	// 	setTitle('Kontrola nádrže')
	// 	setState('inProgress')
	// 	setNotificationStateClass('var(--irrigationBlue)')
	// 	setAction('Probíhá kontrola nádrže')
	// }

	// const setKokot3 = () => {
	// 	setTitle('Doplňte nádrž')
	// 	setState('physicalHelpRequired')
	// 	setNotificationStateClass('var(--warningRed)')
	// 	setAction('Nádrž je prázdná')
	// }

	// setTimeout(() => setNotificationClass('notification'), 3000)
	// setTimeout(() => setKokot(), 8000)
	// setTimeout(() => setKokot2(), 12000)
	// setTimeout(() => setKokot3(), 13000)

	// TEST KOKOT END
	// if your dick is big enough
	// you can take in bigger snuff

	return (
		<div className={notificationClass}>
			<Card className="card p-0-i drop-shadow-2xl">
				<CardContent className="p-0-i">
					<div className="flex-row p-2 pl-4" style={{ background: notificationStateClass }}>
						<span className="text-white title-2 font-semibold">{title}</span>
					</div>
					<div className="flex-row p-2">
						<div className="flex-col p-2 w-3/12">
							<div className="flex-row h-8">
								<span>
									{state === 'inProgress' && <Loading />}
									{state === 'finished' && <Finished />}
									{state === 'physicalHelpRequired' && <Warning />}
								</span>
							</div>
						</div>
						<div className="flex-col p-2 w-9/12">
							<div className="flex-row">
								<span className="title-2" style={{ color: notificationStateClass }}>
									{action}
								</span>
							</div>
						</div>
					</div>
				</CardContent>
			</Card>
		</div>
	)
}
