<template>
  <form class="locale-changer" @submit.prevent="">
    <div @mouseover="hoverList = true" @mouseleave="hoverList = false">
      <button
        class="locale-changer__title"
        type="button"
        aria-controls="locale-dropdown"
        aria-expanded="false"
        @focus="focusList = true"
        @blur="focusList = false"
      >
        <tw-emoji :str="getLocalProperties($i18n.locale).emoji" />{{
          getLocalProperties($i18n.locale).name
        }}
        ▾
      </button>
      <transition name="dropdown">
        <ul
          v-show="focusList || hoverList"
          class="locale-changer__list"
          id="locale-dropdown"
        >
          <li
            v-for="locale in $i18n.availableLocales"
            class="locale-changer__item"
            :key="`locale-${locale}`"
            @click="pickLocale(locale)"
          >
            <tw-emoji :str="getLocalProperties(locale).emoji" />{{
              getLocalProperties(locale).name
            }}
          </li>
        </ul>
      </transition>
    </div>
  </form>
</template>

<script setup lang="ts">
import { ref } from "vue";
import i18n from "@/i18n";

function getLocalProperties(loc: string): { emoji: string; name: string } {
  switch (loc) {
    case "en":
      return { emoji: "🇬🇧", name: "En" };
    case "fr":
      return { emoji: "🇫🇷", name: "Fr" };
    default:
      return { emoji: "", name: "" };
  }
}

const focusList = ref(false);
const hoverList = ref(false);

function pickLocale(newLocale: string) {
  i18n.global.locale.value = newLocale;
  hoverList.value = false;
  focusList.value = false;
  localStorage.setItem("locale", newLocale);
}
</script>

<style scoped>
.locale-changer {
  position: relative;
  display: flex;
  justify-content: right;
}

.locale-changer__list {
  position: absolute;
  bottom: 100%;
  list-style: none;
  border: solid 1px #c1c1c1;
  border-radius: 0 0 3px 3px;
  padding: 0;
  margin: 0.1rem;
}

.locale-changer__item {
  padding: 0.4rem 0.5rem;
  background: rgb(237, 237, 244);
  cursor: pointer;
}

.locale-changer__item:hover {
  background: rgb(219, 219, 226);
}

.dropdown-enter-from,
.dropdown-leave-to {
  opacity: 0;
}

.dropdown-enter-active,
.dropdown-leave-active {
  transition: opacity 100ms ease-in;
}
</style>
