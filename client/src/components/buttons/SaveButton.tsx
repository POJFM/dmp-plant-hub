import Card from '@material-ui/core/Card'
import CardContent from '@material-ui/core/CardContent'

interface ISaveButton {
	active: boolean
	name?: string
}

export default function SaveButton({ active, name }: ISaveButton) {
	let activeClass
	active && (activeClass = 'button hover:button-hover')
	!active && (activeClass = 'button-inactive')

	return (
		<div className="button-wrapper">
			<Card className="button-card">
				<CardContent>
					<div className={activeClass}>
						<button name={name} className="button-text">Uložit</button>
					</div>
				</CardContent>
			</Card>
		</div>
	)
}
