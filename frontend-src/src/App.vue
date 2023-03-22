<template>
  <transition name="splash-transition">
    <SplashLogo
      class="fullheight"
      v-if="splashStatus == 'loading'"
      :glow="glowSplash"
      :message="$t('splash_message')"
    />
    <NeutrinoPage
      class="fullheight"
      :key-pair="keyPair"
      v-else-if="splashStatus === 'finished'"
    />
  </transition>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from "vue";
import { useI18n } from 'vue-i18n'
import SplashLogo from "./components/SplashLogo.vue";
import NeutrinoPage from "./components/NeutrinoPage.vue";
import { initKeys } from "./crypto";
import { KeyPairType } from "@/types";

const { t, locale, availableLocales } = useI18n()

const splashStatus = ref<string>("init");
const glowSplash= ref<boolean>(false);
const keyPair: KeyPairType = reactive({ publicKey: null, privateKey: null });

onMounted(async () => {
  splashStatus.value = "loading";
    window.setTimeout(() => {
      glowSplash.value = true;
    }, 900);
    window.setTimeout(() => {
      splashStatus.value = "finished";
    }, 3200);

    const newKeyPair = await initKeys()
    keyPair.privateKey = newKeyPair.privateKey
    keyPair.publicKey = newKeyPair.publicKey
})

const storedLocale = localStorage.getItem("locale");
if (storedLocale != null && availableLocales.indexOf(storedLocale) !== -1) {
  console.log(`Loaded the locale "${locale}" from localStorage`);
  locale.value = storedLocale;
} else {
  localStorage.removeItem("locale");
  const browserLocale = navigator.language.split("-");
  if (
    browserLocale.length &&
    browserLocale[0].length &&
    availableLocales.indexOf(browserLocale[0]) !== -1
  ) {
    locale.value = browserLocale[0];
  }
}
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
