require('dotenv').config()
import fs from 'fs'
import path from 'path'
import Sequelize from 'sequelize'

const basename = path.basename(__filename)
const db = {}

//@ts-ignore
const sequelize = new Sequelize(process.env.POSTGRES_DB, process.env.POSTGRES_USER, process.env.POSTGRES_PASSWORD, {
	host: process.env.POSTGRES_HOST,
	port: process.env.POSTGRES_PORT,
	dialect: 'postgres',
})

fs.readdirSync(path.join(__dirname, '/models'))
	.filter((file) => file.indexOf('.') !== 0 && file !== basename && file.slice(-3) === '.ts')
	.forEach((file) => {
		const model = sequelize.import(path.join(__dirname, '/models', file))
		db[model.name] = model
	})

Object.keys(db).forEach((modelName) => {
	if (db[modelName].associate) {
		db[modelName].associate(db)
	}
})

//@ts-ignore
db.sequelize = sequelize
//@ts-ignore
db.Sequelize = Sequelize

export default db
