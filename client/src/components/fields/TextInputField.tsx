export default function TextInputField({ item, name, defaultValue, active }: any) {
	let activeClass
	active && (activeClass = 'input-field')
	!active && (activeClass = 'input-field-inactive')

	return (
		<div className={activeClass}>
			<label htmlFor="name" className="input-field-label">
				{name}
			</label>
			<input type="text" id={item} className="input-field-input mt-4" name={item} defaultValue={defaultValue} />
		</div>
	)
}
