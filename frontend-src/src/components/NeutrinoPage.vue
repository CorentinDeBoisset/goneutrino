<template>
  <div class="neutrino-page">
    <localeChanger />
  </div>
</template>

<script lang="ts">
import LocaleChanger from "./LocaleChanger.vue";
import { defineComponent } from "vue";

export default defineComponent({
  name: "NeutrinoPage",
  components: {
    LocaleChanger,
  },
});
</script>
