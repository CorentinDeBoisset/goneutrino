<template>
	<div class="secret-panel">
		<div class="secret-panel__title">{{ panelTitle }}</div>
		<div class="secret-panel__secret-list">
			<div v-for="(secret, idx) in secretList" :key="secret.id">
				<hr v-if="idx !== 0" />
				<SecretDownload :secret="secret" />
			</div>
		</div>
	</div>
</template>

<script lang="ts">
import { Download } from '@/types';
import { defineComponent, PropType } from 'vue';
import SecretDownload from "./SecretDownload.vue";

export default defineComponent({
	name: "SecretDownloadPanel",
	components: {
		SecretDownload,
	},
	props: {
		panelTitle: String,
		secretList: {
			type: Array as PropType<Array<Download>>,
			default: [],
		},
	},
});
</script>

<style scoped>
hr {
	border: 0;
	border-top: solid 1px rgba(0, 0, 0, 0.2);
}
</style>
