<template>
	<div class="secret-panel">
		<div class="secret-panel__title">{{ panelTitle }}</div>
		<div class="secret-panel__secret-list">
			<div v-for="(secret, idx) in secretList" :key="secret.id">
				<hr v-if="idx !== 0" />
				<SecretUpload :upload="secret" />
			</div>
			<form @submit.prevent="" class="secret-panel__file-form">
				<input type="file" name="qdsf" multiple />
			</form>
		</div>
	</div>
</template>

<script lang="ts">
import { Upload } from '@/types';
import { defineComponent, PropType } from 'vue';
import SecretUpload from "./SecretUpload.vue";

export default defineComponent({
	name: "SecretUploadPanel",
	components: {
		SecretUpload,
	},
	props: {
		panelTitle: String,
		secretList: {
			type: Array as PropType<Array<Upload>>,
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

.secret-panel__file-form {
	text-align: center;
}
</style>
