import { createI18n } from "vue-i18n";

import en from "./locales/en.json";
import fr from "./locales/fr.json";

const i18n = createI18n({
  legacy: false,
  locale: process.env.VUE_APP_I18N_LOCALE || "en",
  fallbackLocale: process.env.VUE_APP_I18N_FALLBACK_LOCALE || "en",
  messages: {
    en: en,
    fr: fr,
  },
});

const t = i18n.global.t;

export { i18n as default, t };
