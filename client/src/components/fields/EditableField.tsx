interface IEditableField {
	name: string,
	defaultValue: string | number | undefined,
	active: boolean,
	width: number,
	dataType: string
}

export default function EditableField({ name, defaultValue, active, width, dataType }: IEditableField) {
	let activeClass, isDisabled

	if(active) {
		activeClass = 'input-field'
		isDisabled = false
	} else {
		activeClass = 'input-field-inactive'
		isDisabled = true
	}

	return (
		<div className={`flex-col float-left ${activeClass} w-${width}`} >
			<input
				type={dataType === 'number' ? 'number' : 'text'}
				id={name}
				className="text-center float-left input-field-input"
				name={name}
				defaultValue={defaultValue}
				disabled={isDisabled}
			/>
		</div>
	)
}
