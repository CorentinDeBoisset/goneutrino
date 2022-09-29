	<template>
	<div class="secret-exchange">
		<h3 class="section-title">Partage de secrets avec XXX</h3>
		<div class="secret-exchange__panels">
			<SecretDisplayer :secrets="secretHistory"></SecretDisplayer>
			<SecretSenderForm></SecretSenderForm>
		</div>
	</div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import SecretDisplayer from "./SecretDisplayer.vue";
import SecretSenderForm from "./SecretSenderForm.vue";
import { Download, Upload, DownloadStatus, UploadStatus } from '../types';

export default defineComponent({
	name: "SecretExchange",
	components: {
		SecretDisplayer,
		SecretSenderForm,
	},
	data() {
		return {
			secretHistory: [
				{ id: "1", date: new Date(2022, 7, 3, 12, 4, 5), name: "toto", progress: 0, speed: 123, status: UploadStatus.Encryption },
				{ id: "50", date: new Date(2022, 7, 3, 12, 5, 5), name: "toto", progress: 0.95, speed: 123, status: UploadStatus.Upload },
				{ id: "13", date: new Date(2022, 7, 3, 12, 6, 5), name: "otherqdsfdsqqsdqsdffdsqqsdffdsqaerreazazerrezaazerreaz ffsdqqsdf secret.jpg", progress: 1, speed: 123, status: UploadStatus.Done },
				{ id: "1212", date: new Date(2022, 7, 3, 12, 7, 5), name: "supnetniurae", progress: 0, status: DownloadStatus.Available },
				{ id: "3344", date: new Date(2022, 7, 3, 12, 8, 5), name: "supnetniurae", progress: 0.6, status: DownloadStatus.Download },
				{ id: "6323", date: new Date(2022, 7, 3, 12, 9, 5), name: "somethinqsdf sqdf qsdf qsfd fdq qsdf fdsq fsqd qsdf gelse", progress: 0.4, status: DownloadStatus.Download },
				{ id: "4732", date: new Date(2022, 7, 3, 12, 10, 5), name: "supnetniurae", progress: 1, status: DownloadStatus.Decryption },
				{ id: "4732", date: new Date(2022, 7, 3, 12, 11, 5), name: "supnetniurae", progress: 1, status: DownloadStatus.Done },
			] as Array<Upload | Download>,
		};
	}
})
</script>

<style scoped>
</style>
