export default (sequelize, DataTypes) => {
	const Measured = sequelize.define(
		'measured',
		{
      moisture: {
				type: DataTypes.NUMBER,
				allowNull: false,
			},
      temperature: {
				type: DataTypes.NUMBER,
				allowNull: false,
			},
      humidity: {
				type: DataTypes.NUMBER,
				allowNull: false,
			},
      measureTime: {
        type: DataTypes.DATE(6),
        allowNull: false,
      }
		},
		{
			freezeTableName: true,
		}
	)

	return Measured
}
