require('dotenv').config()
import express from 'express'
import { ApolloServer } from 'apollo-server-express'
import db from './db'
import resolvers from './graphql/resolvers'
import typeDefs from './graphql/typeDefs'
import http from 'http'
import session from 'express-session'
import cookieParser from 'cookie-parser'
import cors from 'cors'

// ! initialize sequelize with session store
const SequelizeStore = require('connect-session-sequelize')(session.Store)

// ! Declare express application
const app = express()

// ! Middleware express
app.use(cookieParser())
app.use(
	session({
		name: process.env.SESSION_NAME,
		secret: process.env.APP_SECRET,
		store: new SequelizeStore({
			//@ts-ignore
			db: db.sequelize,
		}),
		resave: false, // we support the touch method so per that express-session docs this should be set to false
		saveUninitialized: false,
		cookie: {
			secure: process.env.NODE_ENV == 'production',
			maxAge: 24 * 60 * 60 * 1000,
			httpOnly: false,
		},
	})
)

app.use(
	cors({
		origin: ['http://localhost:3000'],
		credentials: true,
	})
)

// ! Apollo server creation
const server = new ApolloServer({
	typeDefs,
	resolvers,
	introspection: true,
	//@ts-ignore
	cors: true,
	context: async ({ req, res, connection }) => {
		if (req) {
			return {
				db,
				res,
				//@ts-ignore
				session: req.session,
				//@ts-ignore
				me: req.session.user,
				secret: process.env.APP_SECRET,
			}
		}
	},
})

server.applyMiddleware({ app, path: '/graphql', cors: false })

const httpServer = http.createServer(app)
//@ts-ignore
db.sequelize.sync({ force: true }).then(async () => {
	console.log(`database synced!`)
})

httpServer.listen({ port: process.env.PORT }, () => {
	console.log(`Apollo server ready at http://localhost:${process.env.PORT}/graphql`)
})
