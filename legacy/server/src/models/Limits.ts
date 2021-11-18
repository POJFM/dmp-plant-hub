export default (sequelize, DataTypes) => {
	const Limits = sequelize.define(
		'limits',
		{
			waterLevel: {
				type: DataTypes.NUMBER,
				allowNull: false,
			},
			waterOverdrawn: {
				type: DataTypes.NUMBER,
				allowNull: false,
			},
			moisture: {
				type: DataTypes.NUMBER,
				allowNull: false,
			},
		},
		{
			freezeTableName: true,
		}
	)

	return Limits
}
