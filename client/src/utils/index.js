import axios from 'axios'

export const getMonths = () => {
	let months = []
	const monthsTranslate = {
			Jan: 'Leden',
			Feb: 'Únor',
			Mar: 'Březen',
			Apr: 'Duben',
			May: 'Květen',
			Jun: 'Červen',
			Jul: 'Červenec',
			Aug: 'Srpen',
			Sep: 'Září',
			Oct: 'Říjen',
			Nov: 'Listopad',
			Dec: 'Prosinec',
		},
		monthsOrigin = ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun', 'Jul', 'Aug', 'Sep', 'Oct', 'Nov', 'Dec'],
		currentMonth = new Date().getMonth(),
		sortedMonths = [...monthsOrigin.slice(0, currentMonth), ...monthsOrigin.slice(currentMonth)]

	sortedMonths.map((sortedMonth) => months?.push(monthsTranslate[`${sortedMonth}`]))

	return months
}

export const fetchLocationFromCoords = (latitude, longitude) => {
	axios
		.request({
			method: 'GET',
			url: `https://api.opencagedata.com/geocode/v1/json?q=${latitude}+${longitude}&key=${process.env.REACT_APP_GEOCODE_API_KEY}`,
			headers: {
				'Content-Type': 'application/json',
			},
		})
		.then((res) => {
			return (
				res.data.results[0].components?.town ||
				res.data.results[0].components?.village ||
				res.data.results[0].components?.city
			)
		})
		.catch((error) => {
			console.error(error)
		})
}

export const fetchCoordsFromLocation = (searchLocationValue) => {
	console.log(searchLocationValue)
	axios
		.request({
			method: 'GET',
			url: `https://api.opencagedata.com/geocode/v1/json?q=${searchLocationValue}&key=${process.env.REACT_APP_GEOCODE_API_KEY}`,
			headers: {
				'Content-Type': 'application/json',
			},
		})
		.then((res) => {
			res.data.results.map((item) => {
				if (item.components.country_code === 'cz') {
					return { lat: item?.geometry.lat, lon: item?.geometry.lng }
				}
			})
		})
		.catch((error) => {
			console.error(error)
		})
}

export const monthRegex = (string) => {
	let tMonth = string.split(/\-(..?)/)[1].substring(1)
	while (tMonth.charAt(0) === '0') {
		tMonth = tMonth.substring(1);
	}
	return parseInt(tMonth)
}