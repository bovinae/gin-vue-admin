/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./index.html",
    "./src/**/*.{vue,js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      backgroundColor: {
        "main": "#F5F5F5",
      }
    },
  },
  plugins: [],
  corePlugins: {
    preflight: false
  }
}

