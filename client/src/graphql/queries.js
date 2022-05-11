import gql from 'graphql-tag'

export const settings = gql`
	query settings {
		getSettings {
			limits_trigger
			water_level_limit
			water_amount_limit
			moist_limit
			scheduled_trigger
			hour_range
			irrigation_duration
			default_water_amount
			chart_type
			theme
			language
			location
			lat
			lon
		}
	}
`

export const dashboard = gql`
	query dashboard {
		getMeasurements {
			id
			timestamp
			hum
			temp
			moist
			with_irrigation
		}
		getIrrigation {
			id
			timestamp
			water_level
			water_amount
			water_overdrawn
		}
		getSettings {
			chart_type
			lat
			lon
		}
	}
`
