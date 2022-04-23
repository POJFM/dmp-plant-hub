export default function TextInputField({ item, name, defaultValue, active, dataType }: any) {
	let activeClass
	active && (activeClass = 'input-field')
	!active && (activeClass = 'input-field-inactive')

	return (
		<div className={activeClass}>
			<label htmlFor="name" className="input-field-label">
				{name}
			</label>
			<input type={dataType === 'number' ? 'number' : 'text'} id={item} className="input-field-input mt-4" name={item} defaultValue={defaultValue} />
		</div>
	)
}
