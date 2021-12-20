import { useState } from 'react'

export default function EditableField({ key, defaultValue, active, width }: any) {
	let activeClass
	active && (activeClass = 'input-field')
	!active && (activeClass = 'input-field-inactive')

	return (
		<div className={`float-left inline-block ${activeClass}`}>
			<input
				type="text"
				id={key}
				className={`w-${width} text-center float-left inline-block input-field-input`}
				name={key}
				defaultValue={defaultValue}
			/>
		</div>
	)
}
