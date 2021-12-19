import { useState, useEffect } from 'react'

// if values are not set then toggle switch acts as an ON / OFF switch
export default function ToggleButton({ item, toggleState, values }: any) {
	const [green, setGreen] = useState<string>()
	const [grey, setGrey] = useState<string>()
	let toggleStateClass, toggleIcon0Class, toggleIcon1Class

	// get colors from css variables
	useEffect(() => {
		setGreen(getComputedStyle(document.body).getPropertyValue('--lightGreen'))
		setGrey(getComputedStyle(document.body).getPropertyValue('--inactiveGrey'))
	}, [])

	if (!values) {
		toggleState && (toggleStateClass = 'input-field-input-toggle-slider-active')
		!toggleState && (toggleStateClass = 'input-field-input-toggle-slider-inactive')
	} else {
		if (toggleState) {
			toggleStateClass = 'input-field-input-toggle-slider-values-1'
			toggleIcon0Class = 'bg-inactiveGrey'
			toggleIcon1Class = 'bg-lightGreen'
		} else {
			toggleStateClass = 'input-field-input-toggle-slider-values-0'
			toggleIcon1Class = 'bg-inactiveGrey'
			toggleIcon0Class = 'bg-lightGreen'
		}
	}

	return (
		<div className="flex-row">
			{/* až bude změnit na svg a dát tam color variable */}
			{values && <img src={`/assets/icons/toggleSwitch/${values[0].label}.svg`} className="flex-col mr-2" />}
			<div className="flex-col input-field-toggle-checkbox-wrapper">
				<input type="checkbox" id={item} name={item} className="input-field-input-toggle" />
				<span className={`input-field-input-toggle-slider ${toggleStateClass}`}></span>
			</div>
			{values && <img src={`/assets/icons/toggleSwitch/${values[1].label}.svg`} className="flex-col ml-2" />}
		</div>
	)
}
