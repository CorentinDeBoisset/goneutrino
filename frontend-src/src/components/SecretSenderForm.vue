<template>
  <form @submit.prevent="" class="sender-form">
    <div class="sender-form__button">
      <input
        type="file"
        name="attachments"
        @change="onAttachmentChange"
        hidden
      />
      <img src="../assets/attachment.svg" />
    </div>
    <div class="sender-form__input-zone">
      <textarea
        class="sender-form__textarea"
        type="text"
        v-model="messageContent"
        placeholder="Message"
        rows="1"
      >
      </textarea>
      <div class="sender-form__sizer" aria-hidden="true">
        {{ messageContent }}
      </div>
    </div>
    <div
      class="sender-form__button"
      :class="{ disabled: messageContent.length === 0 }"
      @click="sendMessage"
    >
      <img src="../assets/send.svg" />
    </div>
  </form>
</template>

<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
  name: "SecretSenderForm",
  data() {
    return {
      messageContent: "",
    };
  },
  methods: {
    onAttachmentChange() {
      return;
    },
    sendMessage() {
      return;
    },
  },
});
</script>

<style scoped>
.sender-form {
  display: flex;
  padding: 0.5rem 0.7rem;
  background-color: var(--secondary-ulight-color);
}

.sender-form__sizer {
  white-space: pre-wrap;
  visibility: hidden;
}

.sender-form__textarea {
  resize: none;
  overflow-y: scroll;
}

.sender-form__sizer,
.sender-form__textarea {
  /* important to size the sizer and the textarea with the exact same properties */
  grid-area: 1 / 1 / 2 / 2;
  border: none;
  padding: 0.3rem 0.5rem 0.3rem;
  font: inherit;
  max-height: 4.5rem;
  outline: none;
  color: white;
  background-color: rgba(255, 255, 255, 0.2);
  border-radius: 0.3rem;
}

.sender-form__input-zone {
  flex: 1;
  display: grid;
  margin: 0 0.5rem;
}

.sender-form__button {
  flex: 0;
  display: flex;
  align-items: center;
}
.sender-form__button img {
  width: 1.5rem;
}
.sender-form__button.disabled {
  opacity: 0.5;
}

.sender-form__button:not(.disabled) {
  cursor: pointer;
}
</style>
