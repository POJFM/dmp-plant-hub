import { isAuthenticated } from './authorization'
import { combineResolvers } from 'graphql-resolvers'

export default {
	Query: {
		appSettings: combineResolvers(isAuthenticated, async (root, args, { db, appSettings }, info) => {
			const appSettingsData = await db.appSettings.findByPk(appSettings.id)
			return appSettingsData
		})
	},
	Mutation: {
		updateAppSettings: async (root, { input }, { db, session }) => {
			const { language, theme } = input

			const appSettings = await db.appSettings.create({
				...input,
			})

/* 			session.measured = {
				id: measured.dataValues.id,
				moisture: measured.dataValues.moisture,
				temperature: measured.dataValues.temperature,
				humidity: measured.dataValues.humidity,
				measureTime: measured.dataValues.measureTime,
			} */

			return appSettings
		},
	},
}
