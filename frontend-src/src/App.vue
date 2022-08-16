<template>
  <div>
    <LocaleChanger />
    <transition name="splash-transition">
      <SplashLogo
        v-if="showSplash"
        :glow="glowSplash"
        :message="$t('splash_message')"
      />
    </transition>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import SplashLogo from "./components/SplashLogo.vue";
import LocaleChanger from "./components/LocaleChanger.vue";

export default defineComponent({
  name: "App",
  data() {
    return {
      showSplash: false,
      glowSplash: false,
    };
  },
  mounted() {
    this.showSplash = true;
    window.setTimeout(() => {
      this.glowSplash = true;
    }, 900);
    window.setTimeout(() => {
      this.showSplash = false;
    }, 3200);
  },
  components: {
    SplashLogo,
    LocaleChanger,
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
