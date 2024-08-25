/** @type {import('tailwindcss').Config} */
module.exports = {
	content: ['**/*.templ'],
	theme: {
		extend: {
			colors: {
				// Custom dark mode colors
				text: '#eaeaea',
				primary: '#22242C',
				secondary: '#32343D',
				ok: '#258E00',
				gray: '#8C8C8C',
				cancel: '#FF7A00'
			},
		},
	},
	plugins: [
	],
}

