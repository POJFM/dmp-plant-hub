import { isAuthenticated } from './authorization'
import { combineResolvers } from 'graphql-resolvers'

export default {
	Query: {
		measured: combineResolvers(isAuthenticated, async (root, args, { db, measured }, info) => {
			const measuredData = await db.measured.findByPk(measured.id)
			return measuredData
		})
	},
	Mutation: {
		createMeasured: async (root, { input }, { db, session }) => {
			const { moisture, temperature, humidity, measureTime } = input

			const measured = await db.measured.create({
				...input,
			})

/* 			session.measured = {
				id: measured.dataValues.id,
				moisture: measured.dataValues.moisture,
				temperature: measured.dataValues.temperature,
				humidity: measured.dataValues.humidity,
				measureTime: measured.dataValues.measureTime,
			} */

			return measured
		},
	},
}
