<template>
    <div class="file-message" @click="downloadAction()">
        <div class="file-message__icon">
            <img v-if="fileType === 'avi'" src="@/assets/file_icons/avi.svg" :alt="altText" />
            <img v-else-if="fileType === 'conf'" src="@/assets/file_icons/conf.svg" :alt="altText" />
            <img v-else-if="fileType === 'doc'" src="@/assets/file_icons/doc.svg" :alt="altText" />
            <img v-else-if="fileType === 'docx'" src="@/assets/file_icons/docx.svg" :alt="altText" />
            <img v-else-if="fileType === 'gif'" src="@/assets/file_icons/gif.svg" :alt="altText" />
            <img v-else-if="fileType === 'gz'" src="@/assets/file_icons/gz.svg" :alt="altText" />
            <img v-else-if="fileType === 'htm'" src="@/assets/file_icons/htm.svg" :alt="altText" />
            <img v-else-if="fileType === 'html'" src="@/assets/file_icons/html.svg" :alt="altText" />
            <img v-else-if="fileType === 'ini'" src="@/assets/file_icons/ini.svg" :alt="altText" />
            <img v-else-if="fileType === 'jpeg'" src="@/assets/file_icons/jpeg.svg" :alt="altText" />
            <img v-else-if="fileType === 'jpg'" src="@/assets/file_icons/jpg.svg" :alt="altText" />
            <img v-else-if="fileType === 'key'" src="@/assets/file_icons/key.svg" :alt="altText" />
            <img v-else-if="fileType === 'log'" src="@/assets/file_icons/log.svg" :alt="altText" />
            <img v-else-if="fileType === 'm4a'" src="@/assets/file_icons/m4a.svg" :alt="altText" />
            <img v-else-if="fileType === 'md'" src="@/assets/file_icons/md.svg" :alt="altText" />
            <img v-else-if="fileType === 'mkv'" src="@/assets/file_icons/mkv.svg" :alt="altText" />
            <img v-else-if="fileType === 'mov'" src="@/assets/file_icons/mov.svg" :alt="altText" />
            <img v-else-if="fileType === 'mp3'" src="@/assets/file_icons/mp3.svg" :alt="altText" />
            <img v-else-if="fileType === 'mp4'" src="@/assets/file_icons/mp4.svg" :alt="altText" />
            <img v-else-if="fileType === 'odp'" src="@/assets/file_icons/odp.svg" :alt="altText" />
            <img v-else-if="fileType === 'pdf'" src="@/assets/file_icons/pdf.svg" :alt="altText" />
            <img v-else-if="fileType === 'png'" src="@/assets/file_icons/png.svg" :alt="altText" />
            <img v-else-if="fileType === 'ppt'" src="@/assets/file_icons/ppt.svg" :alt="altText" />
            <img v-else-if="fileType === 'svg'" src="@/assets/file_icons/svg.svg" :alt="altText" />
            <img v-else-if="fileType === 'tar'" src="@/assets/file_icons/tar.svg" :alt="altText" />
            <img v-else-if="fileType === 'tgz'" src="@/assets/file_icons/tgz.svg" :alt="altText" />
            <img v-else-if="fileType === 'tif'" src="@/assets/file_icons/tif.svg" :alt="altText" />
            <img v-else-if="fileType === 'tiff'" src="@/assets/file_icons/tiff.svg" :alt="altText" />
            <img v-else-if="fileType === 'txt'" src="@/assets/file_icons/txt.svg" :alt="altText" />
            <img v-else-if="fileType === 'wav'" src="@/assets/file_icons/wav.svg" :alt="altText" />
            <img v-else-if="fileType === 'xls'" src="@/assets/file_icons/xls.svg" :alt="altText" />
            <img v-else-if="fileType === 'xlsx'" src="@/assets/file_icons/xlsx.svg" :alt="altText" />
            <img v-else-if="fileType === 'zip'" src="@/assets/file_icons/zip.svg" :alt="altText" />
            <img v-else src="@/assets/file_icons/file.svg" :alt="altText" />
        </div>
        <div class="file-message__info">
            <div class="file-message__title">{{ file.name }}</div>
            <div class="file-message__size">{{ formatSize(file.size) }}</div>
        </div>
    </div>
</template>

<script lang="ts">
import { DownloadFile, UploadFile } from "@/types";
import { defineComponent, PropType } from "vue";

const extensionRegexp = new RegExp("^.*\\.([a-zA-Z0-9]+)$")

export default defineComponent({
    name: "SecretDisplayer",
    props: {
        file: {
            type: Object as PropType<DownloadFile | UploadFile>,
            required: true,
        },
    },
    computed: {
        fileType() {
            const matches = extensionRegexp.exec(this.file.name)
            if (matches === null || matches.length < 1) {
                // The filename has no extension, we return the default
                return "file"
            }
            return matches[1];
        },
        altText() {
            return `Fichier transferé : ${this.file.name}`;
        },
    },
    methods: {
        formatSize(bytes: number) {
            if (Math.abs(bytes) < 1024) {
                return bytes + ' B';
            }
            const units = ['KB', 'MB', 'GB'];
            let u = -1;
            const r = 10 ** 2;
            do {
                bytes /= 1024;
                ++u;
            } while (Math.round(Math.abs(bytes) * r) / r >= 1024 && u < units.length - 1);

            return bytes.toFixed(2) + ' ' + units[u];
        },
        downloadAction() {
            // Start the download
        },
    }
})
</script>

<style scoped>
img {
    width: 2.3rem;
}

.file-message {
    display: flex;
    align-items: stretch;
    font-size: 0.85rem;
    cursor: pointer;
}

.file-message__info {
    display: flex;
    flex-direction: column;
    justify-content: space-evenly;
}

.file-message__icon {
    padding: 0 0.5rem 0 0;
}

.file-message__title {
    font-weight: bold;
}
</style>
