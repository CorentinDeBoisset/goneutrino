<template>
	<div class="secret-exchange">
		<h3 class="section-title">Partage de secrets avec XXX</h3>
		<div class="secret-exchange__panels flex-row align-stretch">
			<SecretUploadPanel class="secret-exchange__panel own-secrets col" panel-title="Secrets envoyés" :secret-list="ownSecrets" />
			<div class="secret-exchange__panel-separator col-auto"></div>
			<SecretDownloadPanel class="secret-exchange__panel peer-secrets col" panel-title="Secrets reçus" :secret-list="peerSecrets" />
		</div>
	</div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import SecretDownloadPanel from './SecretDownloadPanel.vue';
import SecretUploadPanel from './SecretUploadPanel.vue';
import { Download, Upload, DownloadStatus, UploadStatus } from '../types';

export default defineComponent({
	name: "SecretExchange",
	components: {
		SecretDownloadPanel: SecretDownloadPanel,
		SecretUploadPanel: SecretUploadPanel,
	},
	data() {
		return {
			ownSecrets: [
				{ id: "1", name: "toto", progress: 0, speed: 123, status: UploadStatus.Encryption },
				{ id: "50", name: "toto", progress: 0.95, speed: 123, status: UploadStatus.Upload },
				{ id: "13", name: "otherq dsfdsqqsd qsdf fdsq qsdf fdsq aer reaz azer reza azer reaz ffsdqqsdf secret.jpg", progress: 1, speed: 123, status: UploadStatus.Done }
			] as Array<Upload>,
			peerSecrets: [
				{ id: "1212", name: "supnetniurae", progress: 0, status: DownloadStatus.Available },
				{ id: "3344", name: "supnetniurae", progress: 0.6, status: DownloadStatus.Download },
				{ id: "6323", name: "somethinqsdf sqdf qsdf qsfd fdq qsdf fdsq fsqd qsdf gelse", progress: 0.4, status: DownloadStatus.Download },
				{ id: "4732", name: "supnetniurae", progress: 1, status: DownloadStatus.Decryption },
				{ id: "4732", name: "supnetniurae", progress: 1, status: DownloadStatus.Done },
			] as Array<Download>,
		};
	}
})
</script>

<style scoped>
.secret-exchange__panels {
	margin: 1rem;
}
.secret-exchange__panel-separator {
	width: 1px;
	background-color: #007172;
	margin: 0 2rem;
}
</style>
