import { useState, useEffect, Component, createRef, RefObject } from "react";
import i18n from 'i18next'
import { Line, Bar } from 'react-chartjs-2'
import { getMonths } from 'src/utils'
import {
	Chart as ChartJS,
	CategoryScale,
	LinearScale,
	PointElement,
	LineElement,
	BarElement,
	Title,
	Tooltip,
	Legend,
} from 'chart.js'

interface IMeasurementsChart {
	chartType: number
	moist: number[],
	hum: number[],
	temp: number[]
}

interface IWaterConsumptionChart {
	chartType: number
	waterOverdrawn: number[],
	irrigationCount: number[]
}

interface IIrrigationChart {
	chartType: number,
	moist: number[],
	hum: number[],
	temp: number[],
	waterOverdrawn: number[],
	dataframe: number[]
}

ChartJS.register(CategoryScale, LinearScale, PointElement, LineElement, BarElement, Title, Tooltip, Legend)

const months = getMonths()

const GlobalChartOptions = {
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

export function LiveMeasurementsChart({ chartType, temp, hum, moist }: IMeasurementsChart) {
	const liveMeasurementsChartData = {
		data: {
			// 25 cols
			labels: [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
			datasets: [
				{
					label: `${i18n.t("dashboard.moisture")} (%)`,
					backgroundColor: 'rgb(172, 130, 49)',
					borderColor: 'rgb(137, 98, 21)',
					data: moist,
					stack: 'Stack 0',
					yAxisID: 'vp',
				},
				{
					label: `${i18n.t("dashboard.humidity")} (%)`,
					backgroundColor: 'rgb(120, 206, 255)',
					borderColor: 'rgb(30, 141, 203)',
					data: hum,
					stack: 'Stack 1',
					yAxisID: 'vv',
				},
				{
					label: `${i18n.t("dashboard.temperature")} (°C)`,
					backgroundColor: 'rgb(255, 99, 132)',
					borderColor: 'rgb(255, 0, 0)',
					data: temp,
					stack: 'Stack 2',
					yAxisID: 'tv',
				},
			],
		},
		options: {
			...GlobalChartOptions,
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

	const [userData, setUserData] = useState(liveMeasurementsChartData);

	useEffect(() => {
		setUserData(liveMeasurementsChartData)
	}, [moist.length])

	if (!chartType)
		return <Line options={userData.options} data={userData.data} />
	else return <Bar options={userData.options} data={userData.data} />
}

export const WaterConsumptionChart = ({ chartType, waterOverdrawn, irrigationCount }: IWaterConsumptionChart) => {
	const waterConsumptionChartData = {
		data: {
			labels: months,
			datasets: [
				{
					label: i18n.t("dashboard.waterConsumption"),
					backgroundColor: 'rgb(120, 206, 255)',
					borderColor: 'rgb(30, 141, 203)',
					data: waterOverdrawn,
					stack: 'Stack 4',
					yAxisID: 'yAxis1',
				},
				{
					label: i18n.t("dashboard.irrigationCount"),
					backgroundColor: 'rgb(162, 231, 130)',
					borderColor: 'rgb(102, 188, 62)',
					data: irrigationCount,
					stack: 'Stack 5',
					yAxisID: 'yAxis2',
				},
			],
		},
		options: {
			...GlobalChartOptions,
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

	const [userData, setUserData] = useState(waterConsumptionChartData);

	useEffect(() => {
		setUserData(waterConsumptionChartData)
	}, [irrigationCount])

	if (!chartType)
		return <Line options={userData.options} data={userData.data} redraw />
	else return <Bar options={userData.options} data={userData.data} redraw />
}

export function IrrigationChart({ chartType, moist, hum, temp, waterOverdrawn, dataframe }: IIrrigationChart) {
	const irrigationChartData = {
		data: {
			labels: dataframe,
			datasets: [
				{
					label: i18n.t("dashboard.moisture"),
					backgroundColor: 'rgb(172, 130, 49)',
					borderColor: 'rgb(137, 98, 21)',
					data: moist,
					stack: 'Stack 0',
					yAxisID: 'yAxis1',
				},
				{
					label: i18n.t("dashboard.humidity"),
					backgroundColor: 'rgb(120, 206, 255)',
					borderColor: 'rgb(30, 141, 203)',
					data: hum,
					stack: 'Stack 1',
					yAxisID: 'yAxis2',
				},
				{
					label: i18n.t("dashboard.temperature"),
					backgroundColor: 'rgb(255, 99, 132)',
					borderColor: 'rgb(255, 0, 0)',
					data: temp,
					stack: 'Stack 2',
					yAxisID: 'yAxis3',
				},
				{
					label: i18n.t("dashboard.waterConsumption"),
					backgroundColor: 'rgb(162, 231, 130)',
					borderColor: 'rgb(102, 188, 62)',
					data: waterOverdrawn,
					stack: 'Stack 3',
					yAxisID: 'yAxis4',
				},
			],
		},
		options: {
			...GlobalChartOptions,
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
				yAxis4: {
					ticks: {
						callback: (value: any) => {
							return `${value}l`
						},
					},
				},
			},
		},
	}

	const [userData, setUserData] = useState(irrigationChartData);

	useEffect(() => {
		setUserData(irrigationChartData)
	}, [temp])

	if (!chartType)
		return <Line options={userData.options} data={userData.data} redraw />
	else return <Bar options={userData.options} data={userData.data} redraw />
}

export function MeasurementsHistoryChart({ chartType, moist, hum, temp }: IMeasurementsChart) {
	const measurementsHistoryChartData = {
		data: {
			labels: months,
			datasets: [
				{
					label: i18n.t("dashboard.moisture"),
					backgroundColor: 'rgb(172, 130, 49)',
					borderColor: 'rgb(137, 98, 21)',
					data: moist,
					stack: 'Stack 0',
					yAxisID: 'yAxis1',
				},
				{
					label: i18n.t("dashboard.humidity"),
					backgroundColor: 'rgb(120, 206, 255)',
					borderColor: 'rgb(30, 141, 203)',
					data: hum,
					stack: 'Stack 1',
					yAxisID: 'yAxis2',
				},
				{
					label: i18n.t("dashboard.temperature"),
					backgroundColor: 'rgb(255, 99, 132)',
					borderColor: 'rgb(255, 0, 0)',
					circular: true,
					data: temp,
					stack: 'Stack 2',
					yAxisID: 'yAxis3',
					display: false,
				},
			],
		},
		options: {
			...GlobalChartOptions,
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

	const [userData, setUserData] = useState(measurementsHistoryChartData);

	useEffect(() => {
		setUserData(measurementsHistoryChartData)
	}, [temp])

	if (!chartType)
		return <Line options={userData.options} data={userData.data} redraw />
	else return <Bar options={userData.options} data={userData.data} redraw />
}
