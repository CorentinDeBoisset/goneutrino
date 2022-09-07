<template>
	<div class="secret-panel">
		<div class="secret-panel__title">{{ panelTitle }}</div>
		<div class="secret-panel__secret-list">
			<div v-for="(secret, idx) in secretList">
				<SecretUpload :secret-name="secret.name" />
				<hr v-if="idx !== secretList.length - 1" />
			</div>
		</div>
	</div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import SecretUpload from "./SecretUpload.vue";

export default defineComponent({
	name: "SecretUploadPanel",
	components: {
		SecretUpload,
	},
	props: {
		panelTitle: String,
		secretList: {
			type: Array as () => Array<{ id: String, name: String }>,
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
