import { createApp } from "vue";
import { createPinia } from 'pinia';
import App from "./App.vue";
import TwEmoji from "./components/TwEmoji.vue";

import "normalize.css";
import "./global.css";

import i18n from "./i18n";
import router from './router';

const pinia = createPinia();

const app = createApp(App)
    .use(i18n)
    .use(router)
    .use(pinia);

app.component("TwEmoji", TwEmoji);

app.mount("#app");
