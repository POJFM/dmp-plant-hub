import { useEffect, useState } from 'react'
// import { createStore } from 'redux'
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom'
import Sidebar from './components/Sidebar'
import Dashboard from './components/Dashboard'
import Control from './components/Control'
import Settings from './components/Settings'
import InitForm from './components/InitForm'
import { useStyles } from './styles/rootStyles'

import 'bootstrap/dist/css/bootstrap.css'
import './styles/globals.css'

export default function App() {
	const classes = useStyles()
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
		<div className={`row ${classes.app}`}>
			<Router>
				<div className="col-2">
					<Sidebar />
				</div>
				<div className="col-10">
					<Switch>
						<Route
							exact
							path="/"
							component={Dashboard}
							// render={() => <Dashboard city={currentLocation} />}
						/>
						<Route exact path="/control" component={Control} />
						<Route exact path="/settings" component={Settings} />
					</Switch>
				</div>
				{/* <InitForm /> */}
				{!visited && <InitForm />}
			</Router>
		</div>
	)
}
