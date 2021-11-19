import { makeStyles } from '@material-ui/core'

export const useDashboardStyles = makeStyles({
	canvas: {
		height: '400px',
		width: '500px',
	},
	refreshButtonWrapper: {
		position: 'relative'
	},
	refreshButton: {
		float: 'right',
		display: 'flex',
		'justify-content': 'right',
		'align-items': 'right',
	},
	weatherForecastTimeWrapper: {
		display: 'flex',
		justifyContent: 'center',
		alignItems: 'center'
	},
	weatherForecastTime: {
		fontSize: '1.6rem',
	},
	weatherForecastValue: {
		fontSize: '0.9rem'
	},
	weatherForecastValueIcon: {
		width: '25px',
		maxHeight: '100%'
	},
	weatherForecastIcon: {
		width: '75px',
		maxHeight: '100%'
	}
})
