import gql from 'graphql-tag'

// SINGLE TYPE
// language enum CZ/EN
export const createSettings = gql`
	mutation createSettings {
		settings {
			limitsTrigger
			moistLimit
			waterAmountLimit
			waterLevelLimit
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

export const updateSettings = gql`
	mutation updateSettings {
		settings {
			id
			limitsTrigger
			moistLimit
			waterAmountLimit
			waterLevelLimit
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
