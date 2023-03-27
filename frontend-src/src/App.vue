<template>
  <transition name="splash-transition">
    <SplashLogo
      class="fullheight"
      v-if="splashStatus == 'loading'"
      :glow="glowSplash"
      :message="t('splash_message')"
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
import SplashLogo from "./components/SplashLogo.vue";
import NeutrinoPage from "./components/NeutrinoPage.vue";
import { initKeys } from "./crypto";
import { KeyPairType } from "@/types";
import i18n from "@/i18n";

const t = i18n.global.t;

const splashStatus = ref<string>("init");
const glowSplash = ref<boolean>(false);
const keyPair: KeyPairType = reactive({ publicKey: null, privateKey: null });

onMounted(async () => {
  splashStatus.value = "loading";
  window.setTimeout(() => {
    glowSplash.value = true;
  }, 900);
  window.setTimeout(() => {
    splashStatus.value = "finished";
  }, 3200);

  const initKeyPair = await initKeys();

  keyPair.privateKey = initKeyPair.privateKey;
  keyPair.publicKey = initKeyPair.publicKey;
});

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
