import { useState } from 'react'
import Card from '@material-ui/core/Card'
import CardContent from '@material-ui/core/CardContent'

export default function SaveButton({ props, active }: any) {
	let activeClass
  active === 'true' && (activeClass = 'button hover:button-hover')
  active === 'false' && (activeClass = 'button-inactive')

	return (
		<div className="button-wrapper">
			<Card className="button-card">
				<CardContent>
					<div className={activeClass}>
						<span className="button-text">Uložit</span>
					</div>
				</CardContent>
			</Card>
		</div>
	)
}
