import { useState } from 'react'

export default function EnumEditableField(key: any, values: any) {
	return (
		// {values.map(value: any) => {
		//   return(

		//   )
		// }}
		<div className="input-field">
			<input type="text" id="theme" className="input-field-input" name={key} /* value={defaultValue} */ />
		</div>
	)
}
