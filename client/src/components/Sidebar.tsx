import { useState, useEffect } from 'react'
import { Link } from 'react-router-dom'
//import PlantHubIcon from 'img/planthub.png'
import DashboardIcon from '@material-ui/icons/Dashboard'
import SettingsIcon from '@material-ui/icons/Settings'
import ControlIcon from '@material-ui/icons/ControlCamera'
import { useStyles } from './../styles/rootStyles'

export default function Sidebar(props: any) {
	const classes = useStyles()
	const [linkHover, setlinkHover] = useState('off')
	const [activeLink, setActiveLink] = useState('blank')
	useEffect(() => {
		setActiveLink(`${window.location.pathname}`)
	}, [])

	return (
		<div className={`row ${classes.sidebar}`}>
			<div className="col">
				<div className="row">
					{/* <div className={`col ${classes.sidebarRow}`}></div> */}
					<div className={`${classes.sidebarRow} ${classes.sidebarRowTop}`}>
						<img src="/assets/logo/logo-icon.png" className={classes.sidebarLogo} />
						<span className={`${classes.sidebarTitle} ${classes.sidebarRow}`}>PlantHub</span>
					</div>
				</div>
				<Link to="/">
					<div
						className={`row ${classes.sidebarRow} ${linkHover === 'dashboard' && classes.sidebarRowHover} ${
							activeLink === '/' && classes.sidebarRowActive
						}`}
						onMouseEnter={() => setlinkHover('dashboard')}
						onMouseLeave={() => setlinkHover('off')}
						onClick={() => setActiveLink('/')}
					>
						<div className={`${classes.sidebarTextField} ${activeLink === '/' && classes.sidebarRowTextFieldActive}`}>
							<DashboardIcon />
							<span
								className={`${classes.sidebarTextField} ${classes.sidebarText} ${
									activeLink === '/' && classes.sidebarRowTextFieldActive
								}`}
							>
								Dashboard
							</span>
						</div>
					</div>
				</Link>
				<Link to="/control">
					<div
						className={`row ${classes.sidebarRow} ${linkHover === 'control' && classes.sidebarRowHover} ${
							activeLink === '/control' && classes.sidebarRowActive
						}`}
						onMouseEnter={() => setlinkHover('control')}
						onMouseLeave={() => setlinkHover('off')}
						onClick={() => setActiveLink('/control')}
					>
						<div
							className={`col ${classes.sidebarTextField} ${
								activeLink === '/control' && classes.sidebarRowTextFieldActive
							}`}
						>
							<ControlIcon />
							<span
								className={`${classes.sidebarTextField} ${classes.sidebarText} ${
									activeLink === '/control' && classes.sidebarRowTextFieldActive
								}`}
							>
								Control
							</span>
						</div>
					</div>
				</Link>
				<Link to="/settings">
					<div
						className={`row ${classes.sidebarRow} ${linkHover === 'settings' && classes.sidebarRowHover} ${
							activeLink === '/settings' && classes.sidebarRowActive
						}`}
						onMouseEnter={() => setlinkHover('settings')}
						onMouseLeave={() => setlinkHover('off')}
						onClick={() => setActiveLink('/settings')}
					>
						<div
							className={`col ${classes.sidebarTextField} ${
								activeLink === '/settings' && classes.sidebarRowTextFieldActive
							}`}
						>
							<SettingsIcon />
							<span
								className={`${classes.sidebarTextField} ${classes.sidebarText} ${
									activeLink === '/settings' && classes.sidebarRowTextFieldActive
								}`}
							>
								Settings
							</span>
						</div>
					</div>
				</Link>
			</div>
		</div>
	)
}
