import gql from 'graphql-tag'

// SINGLE TYPE
// language enum CZ/EN
export const settings = gql`
	query settings {
		settings {
			limitsTrigger
			waterLevelLimit
			waterAmountLimit
			moistLimit
			scheduledTrigger
			hoursRange
			chartType
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
		dashboard {
			id
			irrigationHistory {
				id
				timestamp
				hum
				temp
				moist
				waterLevel
				waterAmount
				waterOverdrawn
			}
			plantState {
				id
				timestamp
				hum
				temp
				moist
			}
			settings {
				chartType
				lat
				lon
			}
		}
	}
`