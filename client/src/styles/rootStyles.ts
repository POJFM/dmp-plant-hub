import { makeStyles } from '@material-ui/core'

// Colors
const bgColor = '#F5F5D1'
const green = '#66BC3E'
const lightGreen = '#7FD059'
const lighterGreen = '#A2E782'
const darkGreen = '#54A82D'
const white = '#fff'

// Defaults
const transition = 'all .1s ease-in-out'

export const useStyles = makeStyles({
	initForm: {
		position: 'fixed',
		top: 0,
		left: 0,
		background: 'rgba(0,0,0,0.6)',
		zIndex: 2,
		width: '100%',
		height: '100%',
		display: 'flex',
		justifyContent: 'center',
		alignItems: 'center',
	},
	initFormCard: {
		width: '400px',
		height: '500px',
	},
	initFormCardRow: {
		padding: '30px',
	},
	initFormCardRowFirst: {
		marginTop: '20px',
	},
	app: {
		background: bgColor,
	},
	sidebar: {
		height: '100vh',
		background: green,
	},
	sidebarRow: {
		padding: '15px',
		transition: transition,
		display: 'flex',
		justifyContent: 'center',
		alignItems: 'center',
	},
	sidebarRowTop: {
		display: 'flex',
		justifyContent: 'center',
		alignItems: 'center',
	},
	sidebarLogo: {
		width: '50px',
		maxHeight: '100%'
	},
	sidebarTitle: {
		color: white,
		fontWeight: 600,
		fontSize: '1.6em',
	},
	sidebarTextField: {
		color: white,
		fontWeight: 600,
		transition: transition,
		display: 'flex',
		alignItems: 'center',
	},
	sidebarText: {
		marginLeft: '10px',
		fontSize: '1.1em',
	},
	sidebarRowHover: {
		background: lightGreen,
		cursor: 'pointer',
	},
	sidebarRowActive: {
		background: lighterGreen,
		cursor: 'pointer',
	},
	sidebarRowTextFieldActive: {
		color: green,
	},
	card: {
		marginTop: '10px',
		marginRight: '15px',
		padding: '5px',
	},
	cardTwoLeft: {
		marginRight: '-5px',
	},
	cardTwoRight: {
		marginLeft: '-5px',
	},
	cardRow: {
		paddingTop: '5px',
	},
	cardRowTitle: {
		fontSize: '1.3em',
	},
	inputField: {
		borderBottom: `1px solid ${darkGreen}`,
		padding: '0 10px',
		position: 'relative',
	},
	inputFieldEditable: {
		borderBottom: `1px solid ${darkGreen}`,
		padding: '0 10px',
		position: 'relative',
		width: '50px',
	},
	inputFieldLabel: {
		color: darkGreen,
		fontSize: '0.8em',
		position: 'absolute',
	},
	inputFieldInput: {
		border: 'none',
		fontSize: '1.2em',
		outline: 0,
		marginTop: '20px',
	},
	buttonWrapper: {
		padding: 0,
		position: 'relative',
	},
	button: {
		background: green,
		transition: transition,
		display: 'flex',
		justifyContent: 'center',
		alignItems: 'center',
		width: '100px',
		height: '50px',
		paddingTop: '6px',
    float: 'right'
	},
	buttonHover: {
		background: lightGreen,
		cursor: 'pointer',
	},
	buttonText: {
		color: white,
		fontSize: '1.2em',
		fontWeight: 600,
	},
	refreshButtonHover: {
		cursor: 'pointer'
	},
})
