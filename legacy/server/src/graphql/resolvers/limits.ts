import { isAuthenticated } from './authorization'
import { combineResolvers } from 'graphql-resolvers'

export default {
	Query: {
		limits: combineResolvers(isAuthenticated, async (root, args, { db, limits }, info) => {
			// find first one
			const limitsData = await db.limits.findByPk(limits.id)
			return limitsData
		})
	},
	Mutation: {
		createLimits: async (root, { input }, { db, session }) => {
			const { waterLevel, waterOverdrawn, moisture } = input

			const limits = await db.limits.create({
				...input,
			})

/* 			session.measured = {
				id: measured.dataValues.id,
				moisture: measured.dataValues.moisture,
				temperature: measured.dataValues.temperature,
				humidity: measured.dataValues.humidity,
				measureTime: measured.dataValues.measureTime,
			} */

			return limits
		},
		updateLimits: async (root, { input }, { db, session }) => {
			const { waterLevel, waterOverdrawn, moisture } = input

			// dodělat
			const limits = await db.limits.update({
				...input,
			})

/* 			session.measured = {
				id: measured.dataValues.id,
				moisture: measured.dataValues.moisture,
				temperature: measured.dataValues.temperature,
				humidity: measured.dataValues.humidity,
				measureTime: measured.dataValues.measureTime,
			} */

			return limits
		},
	},
}
