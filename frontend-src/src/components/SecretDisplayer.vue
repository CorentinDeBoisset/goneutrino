<template>
  <div class="secret-displayer__panel">
    <div
      v-for="secret in secrets"
      :key="secret.id"
      class="secret-displayer__item"
    >
      <div v-if="secret.downloadString" class="secret-displayer__upload-string">
        <div class="secret-displayer__content">
          {{ secret.downloadString.content }}
        </div>
        <div class="secret-item__date">
          {{ formatDate(secret.downloadString.date) }}
        </div>
      </div>
      <div
        v-else-if="secret.downloadFile"
        class="secret-displayer__upload-file"
      >
        <file-message
          class="secret-displayer__content"
          :file="secret.downloadFile"
        />
        <div class="secret-item__date">
          {{ formatDate(secret.downloadFile.date) }}
        </div>
      </div>
      <div
        v-else-if="secret.uploadString"
        class="secret-displayer__download-string"
      >
        <div class="secret-displayer__content">
          {{ secret.uploadString.content }}
        </div>
        <div class="secret-item__date">
          {{ formatDate(secret.uploadString.date) }}
        </div>
      </div>
      <div
        v-else-if="secret.uploadFile"
        class="secret-displayer__download-file"
      >
        <file-message
          class="secret-displayer__content"
          :file="secret.uploadFile"
        />
        <div class="secret-item__date">
          {{ formatDate(secret.uploadFile.date) }}
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
// import { ref } from "vue";
import { SecretItem } from "@/types";
import FileMessage from "./FileMessage.vue";

export interface Props {
  secrets: Array<SecretItem>;
}

defineProps<Props>();

// TODO: this should be factorized somewhere
function formatDate(d: Date): string {
  const now = new Date();
  // If the date is today, we only show the hours and minutes
  if (
    now.getUTCFullYear() === d.getUTCFullYear() &&
    now.getUTCMonth() === d.getUTCMonth() &&
    now.getUTCDate() === d.getUTCDate()
  ) {
    // FIXME: improve using the translation library
    return `${d.getHours()}:${d.getMinutes()}`;
  }

  // Else, we also show the day
  // FIXME: same as above, not every language uses the same date format
  return `${d.getDate()}/${d.getMonth()} ${d.getHours()}:${d.getMinutes()}`;
}
</script>

<style scoped>
.secret-displayer__panel {
  border-radius: 0.3rem 0.3rem 0 0;
  /* background-color: var(--ultralight-gray-color); */
  padding: 0 1rem;
}

.secret-displayer__item {
  position: relative;
  display: flex;
  flex-direction: column;
  align-items: flex-start;
}

.secret-displayer__download-string,
.secret-displayer__download-file {
  display: flex;
  flex-direction: column;
  align-items: flex-start;
}

.secret-displayer__upload-string,
.secret-displayer__upload-file {
  display: flex;
  flex-direction: column;
  align-items: flex-end;
}

.secret-displayer__upload-string,
.secret-displayer__upload-file {
  align-self: flex-end;
}

.secret-displayer__upload-string .secret-item__date,
.secret-displayer__upload-file .secret-item__date {
  align-self: flex-start;
}

.secret-displayer__content {
  border-radius: 0.5rem;
  margin: 0.3rem 0;
  padding: 0.3rem 0.5rem;
}

.secret-displayer__upload-string .secret-displayer__content,
.secret-displayer__upload-file .secret-displayer__content {
  background-color: var(--secondary-light-color);
  color: white;
}

.secret-displayer__download-string .secret-displayer__content,
.secret-displayer__download-file .secret-displayer__content {
  background-color: var(--light-gray-color);
  color: black;
}

.secret-item__date {
  font-size: 0.75rem;
  padding: 0 0.3rem;
}
</style>
