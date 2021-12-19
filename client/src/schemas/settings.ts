export enum ChartType { LINE, BAR }

export enum Theme { LIGHT, DARK }

export enum Language { CZ, EN }

export type Settings = {
	limitsTrigger: boolean
	waterLevelLimit: number
	waterAmountLimit: number
	moistureLimit: number
	scheduledTrigger: boolean
	hoursRange: number
	chartType: ChartType
	theme: Theme
	language: Language
	location: string
}