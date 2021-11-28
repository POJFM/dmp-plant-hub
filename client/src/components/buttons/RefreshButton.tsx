import { useState } from 'react'
import RefreshIcon from '@material-ui/icons/Refresh'

export const RefreshButton = ({ ...props }: any) => {
	const [buttonHover, setButtonHover] = useState(false)

	return <RefreshIcon className="button hover:button" />
}
