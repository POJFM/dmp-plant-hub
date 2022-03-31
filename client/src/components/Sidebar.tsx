import { useState, useEffect } from 'react'
import axios from 'axios'
import { Link } from 'react-router-dom'
import DashboardIcon from '@material-ui/icons/Dashboard'
import ControlIcon from '@material-ui/icons/ControlCamera'
import SettingsIcon from '@material-ui/icons/Settings'
import RefreshIcon from '@material-ui/icons/Refresh'

export default function Sidebar(props: any) {
	const [linkHover, setlinkHover] = useState('off'),
		[activeLink, setActiveLink] = useState('blank'),
		[christmas, setChristmas] = useState(false)

	useEffect(() => {
		setActiveLink(`${window.location.pathname}`)
		var today = new Date()
		var dd = String(today.getDate()).padStart(2, '0')
		var mm = String(today.getMonth() + 1).padStart(2, '0')
		mm === '12' && dd === '24' && setChristmas(true)
	}, [])

	const handleRestart = () => {
		axios
			.post(
				`${process.env.REACT_APP_GO_API_URL}/live/control`,
				{
					pumpState: false,
					restart: true,
				},
				{
					headers: {
						'Content-Type': 'application/x-www-form-urlencoded',
					},
				}
			)
			// .then((res) => {
			// 	console.log(res)
			// })
			// .catch((error) => {
			// 	console.error(error)
			// })
	}

	return (
		<div className="sidebar">
			<div className="flex-col">
				<div className="flex-row">
					<div className="sidebar-row sidebar-row-top">
						{!christmas && <img src="/assets/logo/logo-icon.png" className="w-16 max-h-full" />}
						{christmas && <img src="/assets/logo/logo-christmas.png" className="w-16 max-h-full" />}
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
							<div className="text-2xl flex items-center">
								<DashboardIcon />
							</div>
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
							<div className="text-2xl flex items-center">
								<ControlIcon />
							</div>
							<span className={`sidebar-row-tf ml-1 title-2 ${activeLink === '/control' && 'sidebar-row-tf-active'}`}>
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
							<div className="text-2xl flex items-center">
								<SettingsIcon />
							</div>
							<span className={`sidebar-row-tf ml-1 title-2 ${activeLink === '/settings' && 'sidebar-row-tf-active'}`}>
								Settings
							</span>
						</div>
					</div>
				</Link>
				<span>
					<div
						className={`flex-row sidebar-row ${linkHover === 'refresh' && 'sidebar-row-hover'}`}
						onMouseEnter={() => setlinkHover('refresh')}
						onMouseLeave={() => setlinkHover('off')}
						onClick={() => handleRestart()}
					>
						<div className={`flex-row sidebar-row-tf`}>
							<div className="text-2xl flex items-center">
								<RefreshIcon />
							</div>
							<span className={`sidebar-row-tf ml-1 title-2`}>Restart HW</span>
						</div>
					</div>
				</span>
			</div>
		</div>
	)
}
