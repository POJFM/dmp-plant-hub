import gql from 'graphql-tag'

// SINGLE TYPE
// language enum CZ/EN
export const settings = gql`
	query settings {
		settings {
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

export const settingsCheck = gql`
	query settingsCheck {
		settings {
			id
		}
	}
`

export const dashboard = gql`
	query dashboard {
		dashboard {
			id
			measurements {
				id
				timestamp
				hum
				temp
				moist
				with_irrigation
			}
			irrigation_history {
				id
				timestamp
				water_level
				water_amount
				water_overdrawn
			}
			settings {
				chart_type
				lat
				lon
			}
		}
	}
`