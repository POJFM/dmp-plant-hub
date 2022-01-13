import gql from 'graphql-tag'

export const createSettings = gql`
	mutation createSettings(
		$limits_trigger: Boolean!
		$water_level_limit: Number
		$water_amount_limit: Number
		$moist_limit: Number
		$scheduled_trigger: Boolean!
		$hour_range: Number
		$location: String!
		$irrigation_duration: Boolean
		$chart_type: Boolean!
		$language: Boolean!
		$theme: Boolean!
		$lat: Number!
		$lon: Number!
	) {
		createSettings(
			limits_trigger: $limits_trigger
			water_level_limit: $water_level_limit
			water_amount_limit: $water_amount_limit
			moist_limit: $moist_limit
			hour_range: $hour_range
			location: $location
			irrigation_duration: $irrigation_duration
			chart_type: $chart_type
			language: $language
			theme: $theme
			lat: $lat
			lon: $lon
		) {
			limits_trigger
			water_level_limit
			water_amount_limit
			moist_limit
			scheduled_trigger
			hour_range
			location
			irrigation_duration
			chart_type
			language
			theme
			lat
			lon
		}
	}
`
