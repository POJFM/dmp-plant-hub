import gql from 'graphql-tag'

export const createSettingsMut = gql`
	mutation createSettings(
		$limits_trigger: Boolean
		$water_level_limit: Float
		$water_amount_limit: Float
		$moist_limit: Float
		$scheduled_trigger: Boolean
		$hour_range: Int
		$location: String
		$irrigation_duration: Int
		$chart_type: Boolean
		$language: Boolean
		$theme: Boolean
		$lat: Float
		$lon: Float
		$default_water_amount: Float
	) {
		createSettings(
			input: {
				limits_trigger: $limits_trigger
				water_level_limit: $water_level_limit
				water_amount_limit: $water_amount_limit
				moist_limit: $moist_limit
				scheduled_trigger: $scheduled_trigger
				hour_range: $hour_range
				location: $location
				irrigation_duration: $irrigation_duration
				chart_type: $chart_type
				language: $language
				theme: $theme
				lat: $lat
				lon: $lon
				default_water_amount: $default_water_amount
			}
		) {
			id
		}
	}
`

export const updateSettingsMut = gql`
	mutation updateSettings(
		$limits_trigger: Boolean
		$water_level_limit: Float
		$water_amount_limit: Float
		$moist_limit: Float
		$scheduled_trigger: Boolean
		$hour_range: Int
		$location: String
		$irrigation_duration: Int
		$chart_type: Boolean
		$language: Boolean
		$theme: Boolean
		$lat: Float
		$lon: Float
		$default_water_amount: Float
	) {
		updateSettings(
			input: {
				limits_trigger: $limits_trigger
				water_level_limit: $water_level_limit
				water_amount_limit: $water_amount_limit
				moist_limit: $moist_limit
				scheduled_trigger: $scheduled_trigger
				hour_range: $hour_range
				location: $location
				irrigation_duration: $irrigation_duration
				chart_type: $chart_type
				language: $language
				theme: $theme
				lat: $lat
				lon: $lon
				default_water_amount: $default_water_amount
			}
		) {
			id
		}
	}
`

const ADD_TODO = gql`
	mutation AddTodo($type: String!) {
		addTodo(type: $type) {
			id
			type
		}
	}
`

export const measurements = gql`
	mutation createMeasurement($hum: Number!, $temp: Number!, $moist: Number!, $with_irrigation: Boolean!) {
		createMeasurement(hum: $hum, temp: $temp, moist: $moist, with_irrigation: $with_irrigation) {
			id
			timestamp
			hum
			temp
			moist
			with_irrigation
		}
	}
`