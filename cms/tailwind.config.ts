/** @type {import('tailwindcss').Config} */
import { type Config } from "tailwindcss";

export default {
  content: [
    "./src/**/*.{ts,tsx}",
  'node_modules/flowbite-react/**/*.{js,jsx,ts,tsx}'
],
  theme: {
    extend: {
      colors: {
        'momentum-orange': '#ec750c',
      },
    },
  },
  plugins: [require('flowbite/plugin')],
} satisfies Config

