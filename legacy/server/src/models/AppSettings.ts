export default (sequelize, DataTypes) => {
	const AppSettings = sequelize.define(
		'appSettings',
		{
			language: {
				type: DataTypes.ENUM('cs', 'en'),
				defaultValue: 'cs',
				allowNull: false,
			},
			theme: {
				type: DataTypes.ENUM('light', 'dark'),
				defaultValue: 'light',
				allowNull: false,
			},
		},
		{
			freezeTableName: true,
		}
	)

	return AppSettings
}
