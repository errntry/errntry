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
      "2xl": "1536px",
      "3xl": "1920px"
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
        emerald: color.emerald,
        fuchsia: color.fuchsia
      },
      animation: {
        moveBackground: "moveBackground 10s ease-in alternate infinite",
        shake: "shake .5s infinite",
        wiggle: "wiggle 1s ease-in-out infinite"
      },
      keyframes: {
        moveBackground: {
          "0%": { backgroundPosition: "0 50%" },
          "50%": { backgroundPosition: "50% 100%" },
          "100%": { backgroundPosition: "50% 0" }
        },
        shake: {
          "0%": { transform: "translate(1px, 1px) rotate(0deg)" },
          "10%": { transform: "translate(-1px, -2px) rotate(-1deg)" },
          "20%": { transform: "translate(-3px, 0px) rotate(1deg)" },
          "30%": { transform: "translate(3px, 2px) rotate(0deg)" },
          "40%": { transform: "translate(1px, -1px) rotate(1deg)" },
          "50%": { transform: "translate(-1px, 2px) rotate(-1deg)" },
          "60%": { transform: "translate(-3px, 1px) rotate(0deg)" },
          "70%": { transform: "translate(3px, 1px) rotate(-1deg)" },
          "80%": { transform: "translate(-1px, -1px) rotate(1deg)" },
          "90%": { transform: "translate(1px, 2px) rotate(0deg)" },
          "100%": { transform: "translate(1px, -2px) rotate(-1deg)" }
        },
        wiggle: {
          "0%, 100%": { transform: "rotate(-3deg)" },
          "50%": { transform: "rotate(3deg)" }
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
