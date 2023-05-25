<template>
  <form @submit.prevent="register()" class="name-selector">
    <div class="name-selector__input-block">
      <div class="name-selector__input-label section-title">
        {{ t("name-selector__hello") }}
      </div>
      <div class="name-selector__input-label section-title">
        {{ t("name-selector__pick-nickname") }}
      </div>
      <input
        class="col-6"
        type="text"
        v-model="nickname"
        :placeholder="t('name-selector__nickname-placeholder')"
      />
    </div>
    <div class="name-selector__submit-block">
      <button
        class="btn"
        :class="{
          'inactive': !nickname || nickname.length == 0,
          'loading': loading,
        }"
        type="submit"
      >
        {{ t("name-selector__submit-action") }}
      </button>
    </div>
    <div v-if="error.length > 0" class="name-selector__error-block">
      {{ error }}
    </div>
  </form>
</template>

<script setup lang="ts">
import { KeyPairType } from "@/types";
import { ref } from "vue";
import { t } from "@/i18n";

export interface Props {
  keyPair: KeyPairType;
}

const props = defineProps<Props>();

const nickname = ref("");
const error = ref("");
const loading = ref(false);

const emit = defineEmits(["next"]);

async function register() {
  if (nickname.value.length <= 0 || props.keyPair.publicKey == null) {
    // TODO: show why we cannot validate?
    return;
  }

  loading.value = true;

  await fetch("/api/v1/register", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({
      Name: nickname.value,
      PublicKey: props.keyPair.publicKey.armor(),
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
    .then(resJson => {
      localStorage.setItem("neutrino-nickname", nickname.value);
      emit("next", { nickname: nickname.value, id: resJson.payload.id });
    })
    .catch((err) => {
      error.value = t(err.msg);
    }).finally(() => {
      loading.value = false;
    });
}
</script>

<style scoped>
.name-selector {
  margin: 1rem;
}

.name-selector__input-label {
  font-size: 1.5rem;
  line-height: 1.8rem;
  text-align: left;
}

.name-selector__input-block {
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  justify-content: center;
}

.name-selector__input-block input {
  margin: 1rem 0 2rem;
}

.name-selector__submit-block {
  text-align: right;
  margin: 2rem 0 0;
}

.name-selector__error-block {
  color: #d94f04;
  margin: 1rem 0 0.5rem;
  text-align: center;
}
</style>
