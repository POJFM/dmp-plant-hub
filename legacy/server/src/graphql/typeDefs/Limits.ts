import { gql } from 'apollo-server-express'

export default gql`
	# -----------------------------------------------
	# TYPES
	# -----------------------------------------------
	type Limits {
		waterLevel: Number!
		waterOverdrawn: Number!
		moisture: Number!,
	}

	# -----------------------------------------------
	# QUERIES
	# -----------------------------------------------
	extend type Query {
		limits: Limits
	}

	# -----------------------------------------------
	# MUTATIONS
	# -----------------------------------------------
	extend type Mutation {
		createLimits(input: LimitsInput!): Limits!
    updateLimits(input: LimitsInput!): Limits!
	}

	# -----------------------------------------------
	# INPUT
	# -----------------------------------------------
	input LimitsInput {
		waterLevel: Number!
		waterOverdrawn: Number!
		moisture: Number!,
	}
`
