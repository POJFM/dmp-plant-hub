import { gql } from 'apollo-server-express'

export default gql`
	# -----------------------------------------------
	# TYPES
	# -----------------------------------------------
	type AppSettings {
		language: String!
		theme: String!
	}

	# -----------------------------------------------
	# QUERIES
	# -----------------------------------------------
	extend type Query {
		appSettings: AppSettings
	}

	# -----------------------------------------------
	# MUTATIONS
	# -----------------------------------------------
	extend type Mutation {
		updateAppSettings(input: UpdateAppSettingsInput!): AppSettings!
	}

	# -----------------------------------------------
	# INPUT
	# -----------------------------------------------
	input UpdateAppSettingsInput {
		language: String!
		theme: String!
	}
`
