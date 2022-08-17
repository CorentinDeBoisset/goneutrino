import { createApp } from "vue";
import App from "./App.vue";
import TwEmoji from "./components/TwEmoji.vue";

import "normalize.css";
import "./global.css";

import i18n from "./i18n";

const app = createApp(App).use(i18n);

app.component("TwEmoji", TwEmoji);

app.mount("#app");
