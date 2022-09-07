<template>
	<div class="secret-item flex-row">
		<div class="col-auto">
			<img class="secret-item__icon" src="../assets/file.svg" alt="File logo" />
		</div>
		<div class="col">
			<div class="flex-row">
				<div class="secret-item__name col-auto">{{ secretName }}</div>
				<div class="secret-item__size col">2mb</div>
				<img class="secret-item__cancel col-auto" src="../assets/x.svg" alt="X">
			</div>
			<progress max="100" :value="progress" :alt="`Upload progress at ${progress}%`">{{ progress }}%</progress>
			<div class="flex-row">
				<div class="secret-item__upload-info col">37%</div>
				<div class="secret-item__upload-info col-auto">12 KB/sec</div>
			</div>
		</div>
	</div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';

export default defineComponent({
	name: "SecretUpload",
	props: {
		secretName: String,
		secretType: String,
	},
	data() {
		return {
			progress: 70,
		};
	},
});
</script>

<style scoped>
progress {
	width: 100%;
	border: none;
	background-color: lightgray;
	color: #007172;
	border-radius: 0.3rem;
	height: 0.3rem;
}

progress::-moz-progress-bar,
progress::-webkit-progress-value,
progress::-webkit-progress-bar {
	background-color: #025259;
	border-radius: 0.3rem;
	height: 0.3rem;
}

.secret-item {
	margin: 1rem 0;
	border-radius: 3px;
}

.secret-item__icon {
	width: 2.2rem;
	margin-right: 0.5rem;
}

.secret-item__name {
	margin-right: 0.4rem;
	font-weight: 600;
}

.secret-item__size {
	margin: 0 0.7rem 0 0.5rem;
}

.secret-item__size,
.secret-item__upload-info {
	font-weight: 300;
	color: rgba(0, 0, 0, 0.6);
}

.secret-item__cancel {
	width: 1.1rem;
}
</style>
