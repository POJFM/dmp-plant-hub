const colors = require('tailwindcss/colors')
module.exports = {
	purge: ['./src/**/*.{js,jsx,ts,tsx}', './public/index.html'],
	darkMode: 'class',
	theme: {
		extend: {
			colors: {
				bg: 'var(--bg)',
				green: 'var(--green)',
				lightGreen: 'var(--lightGreen)',
				lighterGreen: 'var(--lighterGreen)',
				darkGreen: 'var(--darkGreen)',
				white: 'var(--white)',
				transition: 'var(--transition)',
			},
		},
		colors: {
			transparent: 'transparent',
			current: 'currentColor',
		},
	},
	variants: {
		extend: {},
	},
}
