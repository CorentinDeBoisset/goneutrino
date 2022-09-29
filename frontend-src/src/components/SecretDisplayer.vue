<template>
    <div>
        <div
            v-for="secret in secrets"
            :key="secret.id"
            class="secret-displayer__secret"
        >
            <div class="secret-displayer__secret-text">
                {{ secret.name }}
            </div>
            <div class="secret-item__secret-date">{{ formatString(secret.date) }}</div>
        </div>
    </div>
</template>

<script lang="ts">
import { defineComponent, PropType } from "vue";
import { UploadFile, UploadString, DownloadFile, DownloadString } from "@/types";

export default defineComponent({
    name: "SecretDisplayer",
    props: {
        secrets: {
            type: Object as PropType<Array<UploadFile|UploadString|DownloadFile|UploadString>>,
            required: true,
        },
    },
    methods: {
        formatString(d: Date): string {
            const now = new Date();
            // If the date is today, we only show the hours and minutes
            if (now.getUTCFullYear() === d.getUTCFullYear() && now.getUTCMonth() === d.getUTCMonth() && now.getUTCDate() === d.getUTCDate()) {
                // FIXME: improve using the translation library
                return `${d.getHours()}:${d.getMinutes()}`;
            }

            // Else, we also show the day
            // FIXME: same as above, not every language uses the same date format
            return `${d.getDate()}/${d.getMonth()} ${d.getHours()}:${d.getMinutes()}`;
        }
    }
});

</script>
