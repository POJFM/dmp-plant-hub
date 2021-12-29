export default function ToggleButton({ item, toggleState, values }: any) {
	// if values are not set then toggle switch acts as an ON / OFF switch
	let toggleStateClass, toggleIcon0Class, toggleIcon1Class

	toggleState
		? (toggleStateClass = 'input-field-input-toggle-slider-active')
		: (toggleStateClass = 'input-field-input-toggle-slider-inactive')

	if (values) {
		if (toggleState) {
			toggleStateClass = 'input-field-input-toggle-slider-values-1'
			toggleIcon0Class = 'inactive'
			toggleIcon1Class = 'active'
		} else {
			toggleStateClass = 'input-field-input-toggle-slider-values-0'
			toggleIcon0Class = 'active'
			toggleIcon1Class = 'inactive'
		}
	}

	return (
		<div className="flex-row">
			{values && (
				<img
					src={`/assets/icons/toggleSwitch/${
						toggleIcon0Class === 'active' ? values[0].label : values[0].label + 'Inactive'
					}.svg`}
					className="flex-col flex-center mr-2 w-9 transition duration-500 ease-in-out"
				/>
			)}
			<div className="flex-col input-field-toggle-checkbox-wrapper">
				<input type="checkbox" id={item} name={item} className="input-field-input-toggle" />
				<span className={`input-field-input-toggle-slider ${toggleStateClass}`}></span>
			</div>
			{values && (
				<img
					src={`/assets/icons/toggleSwitch/${
						toggleIcon1Class === 'active' ? values[1].label : values[1].label + 'Inactive'
					}.svg`}
					className="flex-col flex-center ml-2 w-9 transition duration-500 ease-in-out"
				/>
			)}
		</div>
	)
}
