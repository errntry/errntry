const color = require("tailwindcss/colors");

module.exports = {
  purge: {
    mode: "all",
    enabled: false,
    content: ["./src/**/*.{js,jsx,ts,tsx}", "./public/index.html"],
    preserveHtmlElements: true
  },
  darkMode: false, // or 'media' or 'class'
  theme: {
    screens: {
      xs: "320px",
      sm: "480px",
      md: "640px",
      lg: "1024px",
      xl: "1280px",
      "2xl": "1536px"
    },
    extend: {
      fontFamily: {
        sans: ["'Mulish'", "sans-serif"],
        serif: ["'Monoton'", "display"],
        mono: ["'Langar'", "display"]
      },
      colors: {
        cyan: color.cyan,
        amber: color.amber,
        emerald: color.emerald
      },
      animation: {
        moveBackground: "moveBackground 10s ease-in alternate infinite"
      },
      keyframes: {
        moveBackground: {
          "0%": { backgroundPosition: "0 50%" },
          "50%": { backgroundPosition: "50% 100%" },
          "100%": { backgroundPosition: "50% 0" }
        }
      }
    }
  },
  variants: {
    extend: {
      display: ["focus-within", "hover", "active", "focus"],
      ringOffsetWidth: ["hover", "active"]
    }
  },
  plugins: [require("@tailwindcss/aspect-ratio")]
};
