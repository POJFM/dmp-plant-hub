import { gql } from 'apollo-server-express'

import appSettingsSchema from './User'
import limitsSchema from './User'
import measuredSchema from './User'
import userSchema from './User'

const linkedSchema = gql`
	type Query {
		_: Boolean
	}
	type Mutation {
		_: Boolean
	}
`

export default [linkedSchema, appSettingsSchema, limitsSchema, measuredSchema, userSchema]
