import gql from 'graphql-tag'

// SINGLE TYPE
// language enum CZ/EN
export const settings = gql`
	query settings {
		settings {
			limitsTrigger
			waterLevelLimit
			waterAmountLimit
			moistureLimit
			scheduledTrigger
			hoursRange
			chartType
			theme
			language
			location
		}
	}
`
// LIVE DATA
export const plantState = gql`
	query plantState($limit: Int = 10) {
		plantState(limit: $limit) {
			id
			humidity
			temperature
			moisture
		}
	}
`

export const plantStateHistory = gql`
	query plantState($limit: Int = 10) {
		plantState(limit: $limit) {
			id
			timestamp
			humidity
			temperature
			moisture
		}
	}
`

export const irrigationHistory = gql`
	query irrigationHistory($limit: Int = 10) {
		irrigationHistory(limit: $limit) {
			id
			timestamp
			humidity
			temperature
			moisture
			waterLevel
			waterAmount
			waterOverdrawn
		}
	}
`

// TEST
export const posts = gql`
	query posts {
		posts {
			id
			title
		}
	}
`
