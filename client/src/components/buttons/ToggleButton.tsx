import { useState } from 'react'

// if values are not set then toggle switch acts as ON / OFF switch
export default function ToggleButton({ key, toggleState, values }: any) {
  let toggleStateClass
  toggleState === 'true' && (toggleStateClass = 'input-field-input-toggle-slider-active')
  toggleState === 'false' && (toggleStateClass = 'input-field-input-toggle-slider')

	return (
		// {values.map(value: any) => {
		//   return(

		//   )
		// }}
		<div>
			{values && <img src={`/assets/icons/toggleSwitch/${values[0].label}.svg`} />}
			<div className="input-field-toggle-checkbox-wrapper">
				<input type="checkbox" className="input-field-input-toggle" name={key} />
				<span className={toggleStateClass}></span>
			</div>
			{values && <img src={`/assets/icons/toggleSwitch/${values[1].label}.svg`} />}
		</div>
	)
}
