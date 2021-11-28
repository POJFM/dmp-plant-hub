import { useState } from 'react'

export default function TextInputField({ key, name, defaultValue }: any) {
	const [activeLabel, setActiveLabel] = useState(false)

	return (
		<div className="input-field" onClick={() => setActiveLabel(true)}>
			<label htmlFor="name" className="input-field-label">
				{name}
			</label>
			<input type="text" id="name" className="input-field-input mt-4" name={key} value={defaultValue} />
		</div>
	)
}
