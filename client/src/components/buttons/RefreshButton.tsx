import { useState } from 'react'
import RefreshIcon from '@material-ui/icons/Refresh'
import { useStyles } from '../../styles/rootStyles'

export const RefreshButton = ( { ...props }: any) => {
	const classes = useStyles()
	const [buttonHover, setButtonHover] = useState(false)

	return (
		<RefreshIcon
			className={`${buttonHover && classes.refreshButtonHover}`}
			onMouseOver={() => setButtonHover(true)}
			onMouseOut={() => setButtonHover(false)}
		/>
	)
}
