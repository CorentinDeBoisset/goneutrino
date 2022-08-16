<template>
  <form class="locale-changer" @submit.prevent="">
    <button type="button" @click="toggleSelector">
      {{ localNames[$i18n.locale] }} ▾
    </button>
    <ul v-show="openSelector">
      <li
        v-for="locale in $i18n.availableLocales"
        :key="`locale-${locale}`"
        @click="pickLocale(locale)"
      >
        {{ localNames[locale] }}
      </li>
    </ul>
  </form>
</template>

<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
  name: "LocaleChanger",
  data() {
    return {
      localNames: {
        en: "🇬🇧 - English",
        fr: "🇫🇷 - Français",
      },
      openSelector: false,
    };
  },
  methods: {
    toggleSelector() {
      console.log(this.openSelector);
      this.openSelector = !this.openSelector;
      console.log(this.openSelector);
    },
    pickLocale(locale) {
      this.$i18n.locale = locale;
      this.openSelector = false;
    },
  },
});
</script>

<style scoped>
.locale-changer select {
  padding: 0.3rem 1rem;
}
</style>
