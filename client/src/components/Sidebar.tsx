import { useState, useEffect } from 'react'
import { Link } from 'react-router-dom'
//import PlantHubIcon from 'img/planthub.png'
import DashboardIcon from '@material-ui/icons/Dashboard'
import SettingsIcon from '@material-ui/icons/Settings'
import ControlIcon from '@material-ui/icons/ControlCamera'

export default function Sidebar(props: any) {
	const [linkHover, setlinkHover] = useState('off')
	const [activeLink, setActiveLink] = useState('blank')
	useEffect(() => {
		setActiveLink(`${window.location.pathname}`)
	}, [])

	return (
		<div className="sidebar">
			<div className="flex-col">
				<div className="flex-row">
					{/* <div className={`col sidebar-row`}></div> */}
					<div className="sidebar-row sidebar-row-top">
						<img src="/assets/logo/logo-icon.png" className="w-16 max-h-full" />
						<span className="sidebar-row sidebar-title">PlantHub</span>
					</div>
				</div>
				<Link to="/">
					<div
						className={`flex-row sidebar-row ${linkHover === 'dashboard' && 'sidebar-row-hover'} ${
							activeLink === '/' && 'sidebar-row-active'
						}`}
						onMouseEnter={() => setlinkHover('dashboard')}
						onMouseLeave={() => setlinkHover('off')}
						onClick={() => setActiveLink('/')}
					>
						<div className={`flex-row sidebar-row-tf ${activeLink === '/' && 'sidebar-row-tf-active'}`}>
							<DashboardIcon />
							<span className={`sidebar-row-tf ml-1 title-2 ${activeLink === '/' && 'sidebar-row-tf-active'}`}>
								Dashboard
							</span>
						</div>
					</div>
				</Link>
				<Link to="/control">
					<div
						className={`flex-row sidebar-row ${linkHover === 'control' && 'sidebar-row-hover'} ${
							activeLink === '/control' && 'sidebar-row-active'
						}`}
						onMouseEnter={() => setlinkHover('control')}
						onMouseLeave={() => setlinkHover('off')}
						onClick={() => setActiveLink('/control')}
					>
						<div className={`flex-row sidebar-row-tf ${activeLink === '/control' && 'sidebar-row-tf-active'}`}>
							<ControlIcon />
							<span
								className={`sidebar-row-tf ml-1 title-2 ${activeLink === '/control' && 'sidebar-row-tf-active'}`}
							>
								Control
							</span>
						</div>
					</div>
				</Link>
				<Link to="/settings">
					<div
						className={`flex-row sidebar-row ${linkHover === 'settings' && 'sidebar-row-hover'} ${
							activeLink === '/settings' && 'sidebar-row-active'
						}`}
						onMouseEnter={() => setlinkHover('settings')}
						onMouseLeave={() => setlinkHover('off')}
						onClick={() => setActiveLink('/settings')}
					>
						<div className={`flex-row sidebar-row-tf ${activeLink === '/settings' && 'sidebar-row-tf-active'}`}>
							<SettingsIcon />
							<span
								className={`sidebar-row-tf ml-1 title-2 ${activeLink === '/settings' && 'sidebar-row-tf-active'}`}
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
