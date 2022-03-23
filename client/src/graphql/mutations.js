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
		$irrigation_duration: Boolean
		$chart_type: Boolean
		$language: Boolean
		$theme: Boolean
		$lat: Float
		$lon: Float
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

// mutation createSettings {
//   createSettings(
//     input: {
//       limits_trigger: true
//       water_level_limit: 30.5
//       water_amount_limit: 41.5
//       moist_limit: 35.7
//       scheduled_trigger: true
//       hour_range: 5
//       location: "Frýdek-Místek"
//       irrigation_duration: true
//       chart_type: true
//       language: true
//       theme: true
//       lat: 18.9
//       lon: 49.5
//     }
//   ) {
//     id
//     limits_trigger
//     water_level_limit
//     water_amount_limit
//     moist_limit
//     scheduled_trigger
//     hour_range
//     location
//     irrigation_duration
//     chart_type
//     language
//     theme
//     lat
//     lon
//   }
// }

// mutation createMeasurement {
//   createMeasurement(input: {
//         hum: 20.3,
//         temp: 23.2,
//         moist: 30.5,
//         with_irrigation: true
//       }
//   ) {
//       id
//       timestamp
//       hum
//       temp
//       moist
//       with_irrigation
//   }
// }
