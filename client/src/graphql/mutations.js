import gql from 'graphql-tag'

// SINGLE TYPE
// language enum CZ/EN
export const updateSettings = gql`
	mutation updateSettings {
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
