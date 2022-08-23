<template>
  <transition name="splash-transition">
    <SplashLogo
      class="fullheight"
      v-if="splashStatus == 'loading'"
      :glow="glowSplash"
      :message="$t('splash_message')"
    />
    <NeutrinoPage class="fullheight" v-else-if="splashStatus == 'finished'" />
  </transition>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import SplashLogo from "./components/SplashLogo.vue";
import NeutrinoPage from "./components/NeutrinoPage.vue";
import { initKeys } from "./crypto";

export default defineComponent({
  name: "App",
  data() {
    return {
      splashStatus: "init",
      glowSplash: false,
    };
  },
  mounted() {
    this.splashStatus = "loading";
    window.setTimeout(() => {
      this.glowSplash = true;
    }, 900);
    window.setTimeout(() => {
      this.splashStatus = "finished";
    }, 3200);

    initKeys();
  },
  created() {
    const locale = localStorage.getItem("locale");
    if (locale != null && this.$i18n.availableLocales.indexOf(locale) !== -1) {
      this.$i18n.locale = locale;
    } else {
      localStorage.removeItem("locale");
    }
    const browserLocale = navigator.language.split("-");
    if (
      browserLocale.length &&
      browserLocale[0].length &&
      this.$i18n.availableLocales.indexOf(browserLocale[0]) !== -1
    ) {
      this.$i18n.locale = browserLocale[0];
    }
  },
  components: {
    SplashLogo,
    NeutrinoPage,
  },
});
</script>

<style scoped>
.splash-transition-enter-from {
  opacity: 0;
}

.splash-transition-enter-active {
  transition: opacity 250ms ease-out;
}

.splash-transition-leave-active,
.splash-transition-leave-to {
  transition: opacity 150ms ease-in;
  opacity: 0;
}
</style>
