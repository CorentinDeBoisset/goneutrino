<template>
  <div class="fullheight neutrino-page flex-column align-center justify-between">
    <div class="neutrino-page__upper-body flex-row">
      <div class="col-6 mx-auto">
        <header class="neutrino-header flex-column align-center">
          <img
            class="neutrino-logo"
            src="./assets/logo_large.svg"
            alt="Neutrino logo"
          />
          <h1>
            <span aria-hidden="true">&#8211;</span>
            {{ t("header-title") }}
            <span aria-hidden="true">&#8211;</span>
          </h1>
          <h2>{{ t("header-subtitle") }}</h2>
        </header>
        <section class="neutrino-content">
          <router-view />
        </section>
      </div>
    </div>
    <footer class="neutrino-footer">
      <LocaleChanger />
    </footer>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from "vue";
import i18n from "@/i18n";
import { useNeutrinoStore } from '@/appStore'
import { t } from "@/i18n";
import LocaleChanger from "./components/LocaleChanger.vue";

const store = useNeutrinoStore();

const storedLocale = localStorage.getItem("locale");
if (
  storedLocale != null &&
  i18n.global.availableLocales.indexOf(storedLocale) !== -1
) {
  console.log(`Loaded the locale "${storedLocale}" from localStorage`);
  i18n.global.locale.value = storedLocale;
} else {
  localStorage.removeItem("locale");
  const browserLocale = navigator.language.split("-");
  if (
    browserLocale.length &&
    browserLocale[0].length &&
    i18n.global.availableLocales.includes(browserLocale[0])
  ) {
    i18n.global.locale.value = browserLocale[0];
  }
}

let initalization: Promise<void>;

onMounted(async () => {
  initalization = store.initializeData();
});
</script>

<style scoped>
.neutrino-header {
  margin: 1rem 0 1rem;
  text-align: center;
}

.neutrino-header h1,
.neutrino-header h2 {
  color: var(--primary-light-color);
  padding: 0;
}

.neutrino-header h1 {
  font-size: 1.3rem;
  letter-spacing: 0.5rem;
  margin: 0.7rem 0 0 0;
  text-transform: uppercase;
  font-style: italic;
  font-weight: 400;
}

.neutrino-header h2 {
  font-size: 1rem;
  margin: 0.5rem 0 0 0;
  font-weight: 200;
  font-style: italic;
}

.neutrino-logo {
  width: 6.5rem;
}

.neutrino-content {
  border: solid 1px var(--secondary-light-color);
  margin: 1rem 0 0;
  border-radius: 0.3rem;
  background-color: #fff;
  box-shadow: 0.2rem 0.2rem 0.5rem rgba(0, 0, 0, 0.1);
}

.neutrino-footer {
  padding: 2rem 3rem 2rem 0;
}
</style>
