import bcrypt from 'bcryptjs'

export default (sequelize, DataTypes) => {
	const User = sequelize.define(
		'user',
		{
			name: {
				type: DataTypes.STRING,
				allowNull: false,
			},
			username: {
				type: DataTypes.STRING,
				allowNull: false,
				unique: {
					args: true,
					message: 'Username must be unique!',
				},
			},
			email: {
				type: DataTypes.STRING,
				allowNull: false,
				unique: true,
				validate: {
					isEmail: {
						args: true,
						msg: 'Invalid email',
					},
				},
			},
			password: {
				type: DataTypes.STRING,
				allowNull: false,
			},
		},
		{
			freezeTableName: true,
		}
	)

	User.findByLogin = async (login) => {
		let user = await User.findOne({
			where: { username: login },
		})
		return user
	}

	User.beforeCreate(async (user) => {
		if (user.password) {
			user.password = await user.generatePasswordHash()
		}
	})

	User.prototype.updatePasswordHash = async function (password) {
		const saltRounds = 10
		return await bcrypt.hash(password, saltRounds)
	}

	User.prototype.generatePasswordHash = async function () {
		const saltRounds = 10
		return await bcrypt.hash(this.password, saltRounds)
	}

	User.prototype.validatePassword = async function (password) {
		return await bcrypt.compare(password, this.password)
	}

	return User
}
