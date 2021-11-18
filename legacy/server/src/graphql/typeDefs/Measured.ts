import { gql } from 'apollo-server-express'

export default gql`
	# -----------------------------------------------
	# TYPES
	# -----------------------------------------------
	type Measured {
		id: ID
		moisture: Number!
		temperature: Number!
		humidity: Number!,
    measureTime: Date!,
	}

	# -----------------------------------------------
	# QUERIES
	# -----------------------------------------------
	extend type Query {
		measured: Measured
	}

	# -----------------------------------------------
	# MUTATIONS
	# -----------------------------------------------
	extend type Mutation {
		createMeasured(input: CreateMeasuredInput!): Measured!
	}

	# -----------------------------------------------
	# INPUT
	# -----------------------------------------------
	input CreateMeasuredInput {
		id: ID
		moisture: Number!
		temperature: Number!
		humidity: Number!,
    measureTime: Date!,
	}
`
