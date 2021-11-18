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
	}
})
