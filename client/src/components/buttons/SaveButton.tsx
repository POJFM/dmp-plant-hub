import { useState } from 'react'
import Card from '@material-ui/core/Card'
import CardContent from '@material-ui/core/CardContent'

export default function SaveButton({ ...props }: any) {
	const [buttonHover, setButtonHover] = useState(false)

	return (
		<div className="button-wrapper">
			<Card className="button hover:button">
				<CardContent>
					<div className="button hover:button-hover">
						<span className="button-text">Uložit</span>
					</div>
				</CardContent>
			</Card>
		</div>
	)
}
