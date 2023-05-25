<template>
  <form @submit.prevent="register()" class="home-page">
    <div class="home-page__input-block">
      <div class="home-page__input-label section-title">
        {{ t("home-page__hello") }}
      </div>
      <div class="home-page__input-label section-title">
        {{ t("home-page__pick-nickname") }}
      </div>
      <input
        class="col-6"
        type="text"
        v-model="nickname"
        :placeholder="t('home-page__nickname-placeholder')"
      />
    </div>
    <div class="home-page__submit-block">
      <button
        class="btn"
        :class="{
          'inactive': nickname.length == 0,
          'loading': loading,
        }"
        type="submit"
      >
        {{ t("home-page__submit-action") }}
      </button>
    </div>
    <div v-if="error.length > 0" class="home-page__error-block">
      {{ error }}
    </div>
  </form>
</template>

<script setup lang="ts">
import { ref } from "vue";
import { t } from "@/i18n";
import { useRouter } from 'vue-router'
import { useNeutrinoStore } from "@/appStore";

const router = useRouter();
const store = useNeutrinoStore();

const nickname = ref("");
const loading = ref(false);
const error = ref("");

async function register() {
  if (nickname.value.length <= 0 || store.keyPair.publicKey === null) {
    // TODO: show why we cannot validate
    return;
  }

  loading.value = true;

  await fetch("/api/v1/register", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({
      Name: nickname.value,
      PublicKey: store.keyPair.publicKey.armor(),
    }),
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
    .then(() => {
      store.setNickname(nickname.value)
      router.replace({ name: 'new-chat' });
    })
    .catch((err) => {
      error.value = t(err.msg);
    }).finally(() => {
      loading.value = false;
    });
}
</script>

<style scoped>
.home-page {
  margin: 1rem;
}

.home-page__input-label {
  font-size: 1.5rem;
  line-height: 1.8rem;
  text-align: left;
}

.home-page__input-block {
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  justify-content: center;
}

.home-page__input-block input {
  margin: 1rem 0 2rem;
}

.home-page__submit-block {
  text-align: right;
  margin: 2rem 0 0;
}

.home-page__error-block {
  color: #d94f04;
  margin: 1rem 0 0.5rem;
  text-align: center;
}
</style>
