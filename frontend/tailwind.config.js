module.exports = {
  future: {
    removeDeprecatedGapUtilities: true,
    purgeLayersByDefault: true
  },
  purge: {
    mode: "all",
    enabled: false,
    content: ["./src/**/*.vue", "./public/index.html"],
    preserveHtmlElements: true
  },
  theme: {
    extend: {
      screens: {
        xs: "320px",
        sm: "480px",
        md: "640px",
        lg: "1024px",
        xl: "1280px",
        xxl: "1440px",
        xxxl: "1920px"
      },
      fontSize: {
        "7xl": "5.333rem",
        "8xl": "7.109rem",
        "9xl": "9.476rem",
        "10xl": "12.632rem",
        "11xl": "16.8rem",
        "12xl": "22.345rem"
      },
      spacing: {
        "72": "18rem",
        "80": "20rem",
        "96": "24rem",
        "102": "28rem",
        "128": "32rem",
        "144": "36rem",
        "160": "40rem",
        "192": "48rem",
        "204": "56rem",
        "1/2": "50%",
        "1/3": "33.33%",
        "1/4": "25%",
        "1/5": "20%",
        "1/7": "14.2857%",
        "2/7": "28.5714%",
        "3/7": "42.8571%",
        "1/10": "10%",
        "2/10": "20%",
        "3/10": "30%",
        "4/10": "40%",
        "5/10": "50%",
        "6/10": "60%",
        "7/10": "70%",
        "8/10": "80%",
        "9/10": "90%"
      }
    },
    fontFamily: {
      mulish: ["'Mulish'", "sans-serif"],
      oleo: ["'Oleo Script'", "sans-serif"]
    }
  },
  variants: {
    accessibility: [
      "responsive",
      "first",
      "last",
      "focus",
      "focus-visible",
      "hover",
      "active"
    ],
    borderStyle: ["responsive", "first", "last", "hover", "active"],
    divideStyle: [
      "responsive",
      "first",
      "last",
      "hover",
      "focus",
      "focus-visible",
      "active"
    ],
    gradientColorStops: [
      "responsive",
      "first",
      "last",
      "hover",
      "focus",
      "focus-visible",
      "active"
    ],
    height: [
      "responsive",
      "first",
      "last",
      "hover",
      "focus",
      "focus-visible",
      "active"
    ],
    inset: [
      "responsive",
      "first",
      "last",
      "hover",
      "focus",
      "focus-visible",
      "active"
    ],
    margin: [
      "responsive",
      "first",
      "last",
      "hover",
      "focus",
      "focus-visible",
      "active"
    ],
    justifyContent: ["responsive", "hover", "focus"],
    padding: [
      "responsive",
      "first",
      "last",
      "hover",
      "focus",
      "focus-visible",
      "active"
    ],
    outLine: [
      "responsive",
      "first",
      "last",
      "hover",
      "focus",
      "focus-visible",
      "active"
    ],
    width: [
      "responsive",
      "first",
      "last",
      "hover",
      "focus",
      "focus-visible",
      "active"
    ],
    zIndex: ["responsive", "first", "last", "hover", "focus", "focus-visible"],
    transform: [
      "responsive",
      "first",
      "last",
      "hover",
      "focus",
      "focus-visible"
    ],
    transformOrigin: [
      "responsive",
      "first",
      "last",
      "hover",
      "focus",
      "focus-visible"
    ],
    scale: ["responsive", "first", "last", "hover", "focus", "focus-visible"],
    rotate: ["responsive", "first", "last", "hover", "focus", "focus-visible"],
    translate: [
      "responsive",
      "first",
      "last",
      "hover",
      "focus",
      "focus-visible"
    ],
    skew: ["responsive", "first", "last", "hover", "focus", "focus-visible"],
    transitionProperty: [
      "responsive",
      "first",
      "last",
      "hover",
      "focus",
      "focus-visible"
    ],
    transitionTimingFunction: [
      "responsive",
      "first",
      "last",
      "hover",
      "focus",
      "focus-visible"
    ],
    transitionDuration: [
      "responsive",
      "first",
      "last",
      "hover",
      "focus",
      "focus-visible"
    ],
    transitionDelay: [
      "responsive",
      "first",
      "last",
      "hover",
      "focus",
      "focus-visible"
    ],
    animation: [
      "responsive",
      "first",
      "last",
      "hover",
      "focus",
      "focus-visible"
    ]
  },
  plugins: []
};
