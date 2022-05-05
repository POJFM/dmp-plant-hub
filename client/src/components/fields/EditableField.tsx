interface IEditableField {
	name: string,
	defaultValue: string | number | undefined,
	active: boolean,
	width: number,
	dataType: string
}

export default function EditableField({ name, defaultValue, active, width, dataType }: IEditableField) {
	let activeClass
	active && (activeClass = 'input-field')
	!active && (activeClass = 'input-field-inactive')

	return (
		<div className={`flex-col float-left ${activeClass} w-${width}`} >
			<input
				type={dataType === 'number' ? 'number' : 'text'}
				id={name}
				className="text-center float-left input-field-input"
				name={name}
				defaultValue={defaultValue}
			/>
		</div>
	)
}
