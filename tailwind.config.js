/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    './pages/**/*.{js,ts,jsx,tsx}',
    './components/**/*.{js,ts,jsx,tsx}',
  ],
  theme: {
    extend: {
      colors: {
        mystic: {
          primary: '#2E3440',
          secondary: '#3B4252',
          accent: '#88C0D0',
          highlight: '#EBCB8B',
          text: '#ECEFF4',
        },
      },
    },
  },
  plugins: [],
  darkMode: 'class', // Enable dark mode
}
