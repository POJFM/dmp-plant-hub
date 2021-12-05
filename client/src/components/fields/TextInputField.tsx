import { useState } from 'react'

export default function TextInputField({ key, name, defaultValue, active }: any) {
	let activeClass
  active === 'true' && (activeClass = 'input-field')
  active === 'false' && (activeClass = 'input-field-inactive')

	return (
		<div className={activeClass}>
			<label htmlFor="name" className="input-field-label">
				{name}
			</label>
			<input type="text" id="name" className="input-field-input mt-4" name={key} value={defaultValue} />
		</div>
	)
}
