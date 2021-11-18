import { skip } from 'graphql-resolvers'
import { ForbiddenError } from 'apollo-server-express'

export const isAuthenticated = (parent, args, { me }) => (me ? skip : new ForbiddenError('Not authenticated as user'))

export const isSessionAuthenticated = (session) => {
	return session.user != undefined
}
