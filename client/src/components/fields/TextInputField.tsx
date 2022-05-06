interface ITextInputField {
	item: string,
	name: string,
	defaultValue: string | number |  undefined,
	active: boolean,
	dataType: string
}

export default function TextInputField({ item, name, defaultValue, active, dataType }: ITextInputField) {
	let activeClass, isDisabled

	if(active) {
		activeClass = 'input-field'
		isDisabled = false
	} else {
		activeClass = 'input-field-inactive'
		isDisabled = true
	}

	return (
		<div className={activeClass}>
			<label htmlFor="name" className="input-field-label">
				{name}
			</label>
			<input 
				type={dataType === 'number' ? 'number' : 'text'} 
				id={item} 
				data-testid={item} 
				className="input-field-input mt-4" 
				name={item} 
				defaultValue={defaultValue}
				disabled={isDisabled}
			/>
		</div>
	)
}
