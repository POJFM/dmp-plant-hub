import { useEffect, useState } from 'react'
import axios from 'axios'
// import { createStore } from 'redux'
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom'
import Sidebar from './components/Sidebar'
import Dashboard from './components/Dashboard'
import Control from './components/Control'
import Settings from './components/Settings'
import InitForm from './components/InitForm'
import Notification from './components/Notification'

import './styles/main.css'

export default function App() {
	const [visited, setVisited] = useState(false)

	useEffect(() => {
		if (localStorage['alreadyVisited']) setVisited(true)
		else {
			//this is the first time
			localStorage['alreadyVisited'] = true
			setVisited(false)
		}
	}, [])

	return (
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
				{/* <InitForm /> */}
				{!visited && <InitForm />}
			</Router>
		</div>
	)
}
