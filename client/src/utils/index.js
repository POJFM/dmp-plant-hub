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
  console.log(months)
	console.log(currentMonth)
	console.log(sortedMonths)
  return months
}