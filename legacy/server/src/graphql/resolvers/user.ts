import { UserInputError } from 'apollo-server'
import { Op } from 'sequelize'
import { isAuthenticated, isSessionAuthenticated } from './authorization'
import { combineResolvers } from 'graphql-resolvers'

export default {
	Query: {
		// ! This query is for the logged in user
		isLoggedIn: async (root, args, { session }, info) => {
			return isSessionAuthenticated(session)
		},
		me: combineResolvers(isAuthenticated, async (root, args, { db, me }, info) => {
			const user = await db.user.findByPk(me.id)
			return user
		}),
		// ! This query grabs all the users
		users: combineResolvers(isAuthenticated, async (root, args, { db }, info) => {
			const users = await db.user.findAll()
			if (!users) {
				throw new Error('No users found')
			}
			return users
		}),
	},
	Mutation: {
		// ! This mutation creates new user
		createUser: async (root, { input }, { db, session }) => {
			const { username, email } = input
			const userExists = await db.user.findOne({
				where: {
					[Op.or]: [{ email }, { username }],
				},
			})
			if (userExists) {
				throw new Error('A user with this email or username already exists')
			}
			const user = await db.user.create({
				...input,
			})

			session.user = {
				id: user.dataValues.id,
				username: user.dataValues.username,
			}

			return user
		},
		login: async (root, { username, password }, { db, session }, info) => {
			const user = await db.user.findOne({
				where: { username },
			})
			if (!user) {
				throw new UserInputError(`User with ${username} does not exist`)
			}

			const isValid = await user.validatePassword(password)
			if (!isValid) {
				throw new UserInputError('Password is invalid')
			}

			session.user = {
				id: user.dataValues.id,
				username: user.dataValues.username,
			}

			return user
		},
		logout: async (root, args, { session, res }, info) => {
			let loggedOutUser = session.user
			await session.destroy()
			res.clearCookie(process.env.SESSION_NAME)
			return loggedOutUser
		},
	},
}
