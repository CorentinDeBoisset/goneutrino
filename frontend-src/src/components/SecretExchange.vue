<template>
	<div class="secret-exchange">
		<h3 class="section-title">Partage de secrets avec XXX</h3>
		<div class="secret-exchange__panels">
			<SecretUploadPanel class="secret-exchange__panel own-secrets" panel-title="Secrets envoyés" :secret-list="ownSecrets" />
			<div class="secret-exchange__panel-separator"></div>
			<SecretDownloadPanel class="secret-exchange__panel peer-secrets" panel-title="Secrets reçus" :secret-list="peerSecrets" />
		</div>
	</div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import SecretDownloadPanel from './SecretDownloadPanel.vue';
import SecretUploadPanel from './SecretUploadPanel.vue';

export default defineComponent({
	name: "SecretExchange",
	components: {
		SecretDownloadPanel: SecretDownloadPanel,
		SecretUploadPanel: SecretUploadPanel,
	},
	data() {
		return {
			ownSecrets: [{ id: "1", name: "toto" }, { id: "13", name: "othersecret.jpg" }],
			peerSecrets: [{ id: "1212", name: "supnetniurae" }, { id: "6323", name: "somethingelse" }],
		};
	}
})
</script>

<style scoped>
.secret-exchange__panels {
	display: flex;
	flex-direction: row;
	justify-content: center;
	align-items: stretch;
}

.secret-exchange__panel-separator {
	width: 1px;
	background-color: #007172;
}

.secret-exchange__panel {
	padding: 1rem;
	margin: 1rem;
	flex-grow: 1;
}
</style>
