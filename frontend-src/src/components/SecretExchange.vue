<template>
  <div class="secret-exchange">
    <h3 class="section-title">Partage de secrets avec XXX</h3>
    <div class="secret-exchange__content">
      <SecretDisplayer
        class="secret-exchange__displayer"
        :secrets="secretHistory"
      ></SecretDisplayer>
      <SecretSenderForm class="secret-exchange__form"></SecretSenderForm>
    </div>
  </div>
</template>

<script setup lang="ts">
import { reactive } from "vue";
import SecretDisplayer from "./SecretDisplayer.vue";
import SecretSenderForm from "./SecretSenderForm.vue";
import { SecretItem, DownloadStatus, UploadStatus } from "../types";

const secretHistory: Array<SecretItem> = reactive([
  {
    id: "1",
    uploadFile: {
      id: "1",
      date: new Date(2022, 7, 3, 12, 4, 5, 300),
      size: 112233421,
      name: "toto.png",
      progress: 0,
      speed: 123,
      status: UploadStatus.Encryption,
    },
  },
  {
    id: "50",
    uploadFile: {
      id: "50",
      date: new Date(2022, 7, 3, 12, 5, 5, 400),
      size: 112233421,
      name: "toto.jpg",
      progress: 0.95,
      speed: 123,
      status: UploadStatus.Upload,
    },
  },
  {
    id: "1212",
    downloadFile: {
      id: "1212",
      date: new Date(2022, 7, 3, 12, 7, 5, 300),
      size: 112233421,
      name: "supnetniurae",
      progress: 0,
      speed: 0,
      status: DownloadStatus.Available,
    },
  },
  {
    id: "3344",
    downloadFile: {
      id: "3344",
      date: new Date(2022, 7, 3, 12, 8, 5, 200),
      size: 112233421,
      name: "supnetniurae",
      progress: 0.6,
      speed: 12,
      status: DownloadStatus.Download,
    },
  },
  {
    id: "125532",
    uploadString: {
      id: "125532",
      date: new Date(2022, 7, 3, 12, 5, 5, 500),
      content: "super secret string",
      status: UploadStatus.Upload,
    },
  },
  {
    id: "6323",
    downloadFile: {
      id: "6323",
      date: new Date(2022, 7, 3, 12, 9, 5, 670),
      size: 112233421,
      name: "somethinqsdf sqdf qsdf qsfd fdq qsdf fdsq fsqd qsdf gelse",
      progress: 0.4,
      speed: 12.5,
      status: DownloadStatus.Download,
    },
  },
  {
    id: "237542",
    downloadFile: {
      id: "237542",
      date: new Date(2022, 7, 3, 12, 10, 5, 750),
      size: 112233421,
      name: "supnetniurae",
      progress: 1,
      speed: -1,
      status: DownloadStatus.Decryption,
    },
  },
  {
    id: "556753",
    downloadString: {
      id: "556753",
      date: new Date(2022, 7, 3, 12, 10, 5, 750),
      content: "super contenu d’une string recue",
      status: DownloadStatus.Done,
    },
  },
  {
    id: "4732",
    downloadFile: {
      id: "4732",
      date: new Date(2022, 7, 3, 12, 11, 5, 250),
      size: 112233421,
      name: "supnetniurae.pdf",
      progress: 1,
      speed: 0,
      status: DownloadStatus.Done,
    },
  },
]);
</script>

<style scoped>
.secret-exchange__content {
  display: flex;
  flex-direction: column;
  align-items: stretch;
}

.section-title {
  margin: 2rem 2rem 1rem;
}
</style>
