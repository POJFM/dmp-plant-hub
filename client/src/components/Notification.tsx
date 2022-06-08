import { useEffect, useState } from 'react'
import axios from 'axios'
import { useTranslation } from 'react-i18next'
import Card from '@material-ui/core/Card'
import CardContent from '@material-ui/core/CardContent'
import Loading from './states/Loading'
import Finished from './states/Finished'
import Warning from './states/Warning'

export default function Notification() {
	const { t } = useTranslation()
	const [title, setTitle] = useState<string>(),
		[notificationClass, setNotificationClass] = useState('hidden'),
		[notificationStateClass, setNotificationStateClass] = useState<string>(),
		[state, setState] = useState<string>(),
		[action, setAction] = useState<any>(),
		[notify, setNotify] = useState(false)

	let waterLevelLeft = ""
	let getNotificationsInterval: ReturnType<typeof setTimeout>

	useEffect(() => {
		!getNotificationsInterval && (getNotificationsInterval = setInterval(() => getNotifications(), 2000))
	}, [])

	const triggerNotification = (res: any, state: boolean) => {
		setTitle(res.data.title)
		setState(res.data.state)
		if(!res.data.action.includes("-")) {
			setAction(res.data.action)
			waterLevelLeft = ""
		} else {
			setAction(res.data.action.split('-')[0])
			waterLevelLeft = res.data.action.split(/\-(....?)/)[1].substring(3)
		}
		setNotify(state)
		if(state) {
			setNotificationClass('notification')
			switch (res.data.state) {
				case 'inProgress': setNotificationStateClass('var(--irrigationBlue)'); break
				case 'finished': setNotificationStateClass('var(--green)'); break
				case 'physicalHelpRequired': setNotificationStateClass('var(--warningRed)'); break
			}
		} else {
			setNotificationClass('hidden')
		}
	}

	const getNotifications = () => {
		axios
			.request({
				method: 'GET',
				url: `${process.env.REACT_APP_GO_API_URL}/live/notify`,
				headers: {
					'Content-Type': 'application/json',
				},
			})
			.then((res) => {
				console.log(res)
				if (res.data.state !== 'inactive') {
					triggerNotification(res, true)
				} else {
					triggerNotification(res, false)
				}
			})
			.catch((error) => {
				console.error(error)
			})
	}

	return (
		<div className={notificationClass}>
			<Card className="card p-0-i drop-shadow-2xl">
				<CardContent className="p-0-i">
					<div className="flex-row p-2 pl-4" style={{ background: notificationStateClass }}>
						<span className="text-white title-2 font-semibold">{t(`notification.${title}`)}</span>
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
									{t(`notification.${action}`, waterLevelLeft && ({ value: waterLevelLeft }))}
								</span>
							</div>
						</div>
					</div>
				</CardContent>
			</Card>
		</div>
	)
}
