import type { Config } from "tailwindcss";

export default {
  content: [
    "./src/pages/**/*.{js,ts,jsx,tsx,mdx}",
    "./src/components/**/*.{js,ts,jsx,tsx,mdx}",
    "./src/app/**/*.{js,ts,jsx,tsx,mdx}",
  ],
  theme: {
    colors: {
      "dark-primary": "#131a1c",
      "dark-secondary": "#1b2224",
      red: "#e74c4c",
      green: "#6bb05d",
      blue: "#0183ff",
      grey: "#dddfe2",
      white: "#fff",
    },
  },
  plugins: [],
} satisfies Config;
