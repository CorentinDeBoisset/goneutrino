<template>
	<div class="secret-item flex-row">
		<div class="col-auto">
			<img class="secret-item__icon" src="../assets/doc.svg" alt="File logo" />
		</div>
		<div class="col">
			<div class="flex-row">
				<div class="secret-item__name col">{{ upload.name }}</div>
				<div class="col-1">
					<img v-if="upload.progress < 1" class="secret-item__cancel" src="../assets/x.svg" alt="X" />
				</div>
			</div>
			<progress max="1" :value="upload.progress" :alt="`Upload progress at ${upload.progress * 100}%`">
				{{ upload.progress * 100 }}%
			</progress>
			<div class="flex-row">
				<div class="secret-item__upload-info col">
					{{ (upload.progress * 100).toLocaleString(undefined, { maximumFractionDigits:1 }) }}%
				</div>
				<div class="secret-item__upload-info col-auto">
					<div v-if="upload.progress < 1">
						{{ formatSize(upload.speed) }}/sec
					</div>
				</div>
			</div>
		</div>
	</div>
</template>

<script lang="ts">
import { defineComponent, PropType } from 'vue';
import { formatSize } from '../util';
import { Upload, UploadStatus } from '../types';

export default defineComponent({
	name: "SecretUpload",
	props: {
		upload: { type: Object as PropType<Upload>, required: true },
	},
	methods: {
		formatSize,
	}
});
</script>

<style scoped>
.secret-item {
	margin: 1rem 0;
	border-radius: 3px;
}

.secret-item__icon {
	width: 2rem;
	margin-right: 0.8rem;
	opacity: 0.6;
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
.secret-item__cancel.hidden {
	visibility: hidden;
}
</style>
