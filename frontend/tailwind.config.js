module.exports = {
  future: {
    removeDeprecatedGapUtilities: true,
    purgeLayersByDefault: true
  },
  purge: {
    enabled: false,
    content: ["./src/**/*.vue", "./public/index.html"],
    preserveHtmlElements: true
  },
  theme: {
    extend: {}
  },
  variants: {},
  plugins: []
};
