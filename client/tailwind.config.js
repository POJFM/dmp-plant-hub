module.exports = {
	purge: ['./src/**/*.{js,jsx,ts,tsx}', './public/index.html'],
	mode: 'jit',
	theme: {
		fontFamily: {
			default: ['Montserrat Alternates', 'sans-serif'],
		},
    textColor: {
      primary: 'white',
      secondary: 'black',
      white: "#FFFFFF",
      black: "#000000",
    },
		extend: {
      colors: {
        'bg-beige': '#F5F5D1',
        'green': '#66BC3E',
        'light-green': '#7FD059',
        'lighter-green': '#A2E782',
        'dark-green': '#54A82D',
        'inactive-grey': '#9FA69B',
        'inactive-light-grey': '#CCCCCC',
        'warning-red': '#DB4C4C',
        'irrigation-blue': '#78CEFF',
        'dt-green': '#144000',
        'dt-bg-green': '#0F2A03',
        'dt-card-green': '#314428',
				'dt-hover-green': '#416A2D',
				'dt-text-green': '#66BC3E',
      },
    },
	},
	darkMode: 'class',
	variants: {
		extend: {},
	},
}