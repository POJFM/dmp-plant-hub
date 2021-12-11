import { useState } from 'react'

export default function EditableField({ key, defaultValue }: any) {
	const [activeLabel, setActiveLabel] = useState(false)

	return (
		<div className="float-left inline-block input-field" onClick={() => setActiveLabel(true)}>
			<input
				type="text"
				id="name"
				className="float-left inline-block input-field-input"
				name={key}
				value={defaultValue}
			/>
		</div>
	)
}
