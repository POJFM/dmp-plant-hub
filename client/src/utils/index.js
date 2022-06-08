import axios from 'axios'
import i18n from 'i18next'

export const getMonths = () => {
	let months = []
	const monthsTranslate = {
			Jan: i18n.t("months.january"),
			Feb: i18n.t("months.february"),
			Mar: i18n.t("months.march"),
			Apr: i18n.t("months.april"),
			May: i18n.t("months.may"),
			Jun: i18n.t("months.june"),
			Jul: i18n.t("months.july"),
			Aug: i18n.t("months.august"),
			Sep: i18n.t("months.september"),
			Oct: i18n.t("months.october"),
			Nov: i18n.t("months.november"),
			Dec: i18n.t("months.december"),
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
		tMonth = tMonth.substring(1)
	}
	return parseInt(tMonth)
}

export const dayRegex = (string) => {
	let tDay = string.split(/\-(.....?)/)[1].substring(4)
	while (tDay.charAt(0) === '0') {
		tDay = tDay.substring(1)
	}
	return parseInt(tDay)
}

export const timeRegex = (string) => {
	let tTime = string.split(/\-(...........?)/)[1].substring(6)
	while (tTime.charAt(0) === '0') {
		tTime = tTime.substring(1)
	}
	return tTime
}
