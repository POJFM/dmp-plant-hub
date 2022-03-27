import { Line, Bar } from 'react-chartjs-2'
import { getMonths } from 'src/utils'

const months = getMonths()

const chartOptions = {
	responsive: true,
	maintainAspectRatio: false,
	interaction: {
		mode: 'index' as const,
		intersect: false,
	},
	plugins: {
		legend: {
			display: false,
		},
	},
}

export function LiveMeasurementsChart({ chartType, temp, hum, moist }: any) {
	const liveMeasurementsChartData = {
		data: {
			// 25 cols
			labels: [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
			datasets: [
				{
					label: 'Vlhkost půdy (%)',
					backgroundColor: 'rgb(172, 130, 49)',
					borderColor: 'rgb(137, 98, 21)',
					data: moist,
					stack: 'Stack 0',
					yAxisID: 'vp',
				},
				{
					label: 'Vlhkost vzduchu (%)',
					backgroundColor: 'rgb(120, 206, 255)',
					borderColor: 'rgb(30, 141, 203)',
					data: hum,
					stack: 'Stack 1',
					yAxisID: 'vv',
				},
				{
					label: 'Teplota vzduchu (°C)',
					backgroundColor: 'rgb(255, 99, 132)',
					borderColor: 'rgb(255, 0, 0)',
					data: temp,
					stack: 'Stack 2',
					yAxisID: 'tv',
				},
			],
		},
		options: {
			...chartOptions,
			plugins: {
				legend: {
					display: false,
				},
			},
			scales: {
				xAxes: {
					ticks: {
						display: false,
					},
				},
			},
		},
	}

	if (!chartType)
		return <Line options={liveMeasurementsChartData.options} data={liveMeasurementsChartData.data} />
	else return <Bar options={liveMeasurementsChartData.options} data={liveMeasurementsChartData.data} />
}

export const WaterConsumptionChart = ({ chartType, waterOverdrawn, irrigationCount }: any) => {
	const waterConsumptionChartData = {
		data: {
			labels: months,
			datasets: [
				{
					label: 'Spotřebováno vody',
					backgroundColor: 'rgb(120, 206, 255)',
					borderColor: 'rgb(30, 141, 203)',
					data: /* waterOverdrawn */ [10, 7, 10, 20, 5, 35, 20],
					stack: 'Stack 4',
					yAxisID: 'yAxis1',
				},
				{
					label: 'Počet zavlažení',
					backgroundColor: 'rgb(162, 231, 130)',
					borderColor: 'rgb(102, 188, 62)',
					data: irrigationCount,
					stack: 'Stack 5',
					yAxisID: 'yAxis2',
				},
			],
		},
		options: {
			...chartOptions,
			plugins: {
				legend: {
					display: false,
				},
			},
			scales: {
				yAxis1: {
					ticks: {
						callback: (value: any) => {
							return `${value}l`
						},
					},
				},
			},
		},
	}

	if (!chartType)
		return <Line options={waterConsumptionChartData.options} data={waterConsumptionChartData.data} />
	else return <Bar options={waterConsumptionChartData.options} data={waterConsumptionChartData.data} />
}

export function IrrigationChart({ chartType, moist, hum, temp, irrigationCount }: any) {
	const irrigationChartData = {
		data: {
			labels: months,
			datasets: [
				{
					label: 'Vlhkost půdy',
					backgroundColor: 'rgb(172, 130, 49)',
					borderColor: 'rgb(137, 98, 21)',
					data: /* moist */[1, 5, 1, 25, 1, 8, 21],
					stack: 'Stack 0',
					yAxisID: 'yAxis1',
				},
				{
					label: 'Vlhkost vzduchu',
					backgroundColor: 'rgb(120, 206, 255)',
					borderColor: 'rgb(30, 141, 203)',
					data: /* hum */[3, 5, 10, 5, 8, 6, 19],
					stack: 'Stack 1',
					yAxisID: 'yAxis2',
				},
				{
					label: 'Teplota vzduchu',
					backgroundColor: 'rgb(255, 99, 132)',
					borderColor: 'rgb(255, 0, 0)',
					data: /* temp */[5, 8, 10, 20, 5, 10, 20],
					stack: 'Stack 2',
					yAxisID: 'yAxis3',
				},
				{
					label: 'Počet zavlažení',
					backgroundColor: 'rgb(162, 231, 130)',
					borderColor: 'rgb(102, 188, 62)',
					data: irrigationCount,
					stack: 'Stack 3',
					yAxisID: 'yAxis4',
				},
			],
		},
		options: {
			...chartOptions,
			plugins: {
				legend: {
					display: false,
				},
			},
			scales: {
				yAxis1: {
					ticks: {
						callback: (value: any) => {
							return `${value}%`
						},
					},
				},
				yAxis2: {
					ticks: {
						callback: (value: any) => {
							return `${value}%`
						},
					},
				},
				yAxis3: {
					ticks: {
						callback: (value: any) => {
							return `${value}°C`
						},
					},
				},
			},
		},
	}

	if (!chartType)
		return <Line options={irrigationChartData.options} data={irrigationChartData.data} />
	else return <Bar options={irrigationChartData.options} data={irrigationChartData.data} />
}

export function MeasurementsHistoryChart({ chartType, moist, hum, temp }: any) {
	const measurementsHistoryChartData = {
		data: {
			labels: months,
			datasets: [
				{
					label: 'Vlhkost půdy',
					backgroundColor: 'rgb(172, 130, 49)',
					borderColor: 'rgb(137, 98, 21)',
					data: /* moist */[1, 5, 1, 25, 1, 8, 21],
					stack: 'Stack 0',
					yAxisID: 'yAxis1',
				},
				{
					label: 'Vlhkost vzduchu',
					backgroundColor: 'rgb(120, 206, 255)',
					borderColor: 'rgb(30, 141, 203)',
					data: /* hum */[3, 5, 10, 5, 8, 6, 19],
					stack: 'Stack 1',
					yAxisID: 'yAxis2',
				},
				{
					label: 'Teplota vzduchu',
					backgroundColor: 'rgb(255, 99, 132)',
					borderColor: 'rgb(255, 0, 0)',
					circular: true,
					data: /* temp */[5, -8, -10, 20, 5, 35, 20],
					stack: 'Stack 2',
					yAxisID: 'yAxis3',
					display: false,
				},
			],
		},
		options: {
			...chartOptions,
			scales: {
				yAxis1: {
					ticks: {
						callback: (value: any) => {
							return `${value}%`
						},
					},
				},
				yAxis2: {
					ticks: {
						callback: (value: any) => {
							return `${value}%`
						},
					},
				},
				yAxis3: {
					ticks: {
						callback: (value: any) => {
							return `${value}°C`
						},
					},
				},
			},
		},
	}

	if (!chartType)
		return <Line options={measurementsHistoryChartData.options} data={measurementsHistoryChartData.data} />
	else return <Bar options={measurementsHistoryChartData.options} data={measurementsHistoryChartData.data} />
}
