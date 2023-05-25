<template>
  <form @submit.prevent="newchat()" class="new-chat-page">
    <div class="new-chat__submit-block">
      <div class="new-chat__input-label section-title">
        {{ t("new-chat__title") }}
      </div>
      <button
        class="btn"
        :class="{
          'loading': loading,
        }"
        type="submit"
      >
        {{ t("new-chat__submit-action") }}
      </button>
    </div>
    <div v-if="error.length > 0" class="new-chat__error-block">
      {{ error }}
    </div>
  </form>
</template>

<script setup lang="ts">
import { ref } from "vue";
import { t } from "@/i18n";
import { useNeutrinoStore } from "@/appStore";
import router from "@/router";

const store = useNeutrinoStore();

const loading = ref(false);
const error = ref("");

async function newchat() {
  loading.value = true;
  await fetch("/api/v1/new-chat", {
    method: "POST",
  })
    .then((res) => {
      if (res.status >= 500) {
        throw { msg: "api__generic_error" };
      }
      if (res.status >= 400) {
        throw res.json();
      }
      return res.json();
    })
    .then((res) => {
      router.push({ name: 'chat', params: { 'uuid': res.uuid } });
    })
    .catch((err) => {
      error.value = t(err.msg);
    }).finally(() => {
      loading.value = false;
    })
};
</script>

<style scoped>
.new-chat-page {
  margin: 1rem;
}

.new-chat__input-label {
  font-size: 1.5rem;
  line-height: 1.8rem;
  text-align: left;
}

.new-chat__input-block {
  display: flex;
  flex-direction: column;
  align-items: center;
}

.new-chat__input-block input {
  margin: 1rem 0 2rem;
}

.new-chat__submit-block {
  text-align: right;
  margin: 2rem 0 0;
}

.new-chat__error-block {
  color: #d94f04;
  margin: 1rem 0 0.5rem;
  text-align: center;
}
</style>
