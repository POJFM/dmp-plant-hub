import { useState, useEffect } from 'react'
import Card from '@material-ui/core/Card'
import CardContent from '@material-ui/core/CardContent'

interface ICancelButton {
	active: boolean
}

export default function CancelButton({ active }: ICancelButton) {
	const [textColor, setTextColor] = useState<string>(),
	[activeClass, setActiveClass] = useState<string>()

	useEffect(() => {
		if(!active)  {
			setActiveClass('button-inactive')
			setTextColor('white')
		} else {
			setActiveClass('button cancel-button hover:cancel-button-hover')
			setTextColor('var(--darkGreen)')
		}
	}, [active])

	return (
		<div className="button-wrapper">
			<Card className="button-card">
				<CardContent>
					<div className={activeClass}>
						<span className="cancel-button-text" style={{color: textColor}}>Zrušit</span>
					</div>
				</CardContent>
			</Card>
		</div>
	)
}
