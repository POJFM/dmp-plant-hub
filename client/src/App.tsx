import { useEffect, useState } from 'react'
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom'
import Sidebar from './components/Sidebar'
import Dashboard from './components/Dashboard'
import Control from './components/Control'
import Settings from './components/Settings'
import InitForm from './components/InitForm'
import Notification from './components/Notification'
import Christmas from './components/events/Christmas'

import './styles/main.css'

export default function App() {
	const [christmas, setChristmas] = useState(false)

	useEffect(() => {
		var today = new Date()
		var dd = String(today.getDate()).padStart(2, '0')
		var mm = String(today.getMonth() + 1).padStart(2, '0')
		mm === '12' && dd === '24' && setChristmas(true)

		let theme
		localStorage && (theme = localStorage.getItem("theme"))
		if (theme === 'dark') {
			document.body.classList.add('dark')
			document.documentElement.setAttribute("theme", "dark")
			localStorage.setItem("theme", "dark")
		} else if (theme === 'light') {
			document.body.classList.remove('dark')
			document.documentElement.setAttribute("theme", "light")
			localStorage.setItem("theme", "light")
		}
	}, [])

	return (
		<>
			{christmas && (
				<div className="flex-row">
					<Christmas />
				</div>
			)}
			<div className="flex-row app">
				<Router>
					<div className="flex-col w-2/12">
						<Sidebar />
					</div>
					<div className="flex-col w-10/12">
						<Switch>
							<Route exact path="/" component={Dashboard} />
							<Route exact path="/control" component={Control} />
							<Route exact path="/settings" component={Settings} />
						</Switch>
						<div className="flex-center">
							<Notification />
						</div>
					</div>
					<InitForm />
				</Router>
			</div>
		</>
	)
}
