import gql from 'graphql-tag'

// SINGLE TYPE
// language enum CZ/EN
export const updateSettings = gql`
	mutation updateSettings(
		$limits_trigger: Number!
		$moist_limit: Number
		$water_amount_limit: Number
		$water_level_limit: Number
		$scheduled_trigger: Number!
		$hours_range: Number
		$irrigation_duration: Number
		$chart_type: Number!
		$theme: Number!
		$language: Number!
		$location: String!
		$lat: Number!
		$lon: Number!
	) {
		updateSettings(
			limits_trigger: $limits_trigger
			moist_limit: $moist_limit
			water_amount_limit: $water_amount_limit
			water_level_limit: $water_level_limit
			hours_range: $hours_range
			irrigation_duration: $irrigation_duration
			chart_type: $chart_type
			theme: $theme
			language: $language
			location: $location
			lat: $lat
			lon: $lon
		) {
			limits_trigger
			water_level_limit
			water_amount_limit
			moist_limit
			scheduled_trigger
			hours_range
			irrigation_duration
			chart_type
			theme
			language
			location
			lat
			lon
		}
	}
`
