<template>
	<div class="secret-item">
		<div class="secret-item__name">
			{{ secretName }}
		</div>
		<div class="secret-item__action">
			<tw-emoji str="⬇️" />
			<tw-emoji str="❌" />
		</div>
	</div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';

export default defineComponent({
	name: "SecretDownload",
	props: {
		secretName: String,
		secretType: String,
	},
});
</script>

<style scoped>
.secret-item {
	padding: 1rem;
	margin: 0.5rem;
	border-radius: 3px;
	display: flex;
	flex-direction: row;
}
.secret-item__name {
	flex: 1;
	margin: 0 0.4rem;
}

.secret-item__action {
	flex: 0;
	margin: 0 0.2rem;
}
</style>
