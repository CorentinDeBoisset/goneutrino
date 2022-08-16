const { defineConfig } = require("@vue/cli-service");
module.exports = defineConfig({
  transpileDependencies: true,
  lintOnSave: false,
  outputDir: "../web/frontend-app",
  assetsDir: "static",

  pages: {
    index: {
      entry: "src/main.ts",
      template: "public/index.html",
      title: "Secret-send - Send your secrets securely to third-parties",
    },
  },

  pluginOptions: {
    i18n: {
      locale: "en",
      fallbackLocale: "en",
      localeDir: "locales",
      enableLegacy: false,
      runtimeOnly: false,
      compositionOnly: false,
      fullInstall: true,
    },
  },
});
