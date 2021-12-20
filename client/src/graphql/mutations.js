import gql from 'graphql-tag'

// SINGLE TYPE
// language enum CZ/EN
export const updateSettings = gql`
	mutation updateSettings(
		$limitsTrigger: Number!
		$moistLimit: Number
		$waterAmountLimit: Number
		$waterLevelLimit: Number
		$scheduledTrigger: Number!
		$hoursRange: Number
		$chartType: Number!
		$theme: Number!
		$language: Number!
		$location: String!
		$lat: Number!
		$lon: Number!
	) {
		updateSettings(
			limitsTrigger: $limitsTrigger
			moistLimit: $moistLimit
			waterAmountLimit: $waterAmountLimit
			waterLevelLimit: $waterLevelLimit
			hoursRange: $hoursRange
			chartType: $chartType
			theme: $theme 
			language: $language
			location: $location
			lat: $lat
			lon: $lon
		) {
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