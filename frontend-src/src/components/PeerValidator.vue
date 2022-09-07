<template>
	<form @submit.prevent="" class="peer-validator">
		<h3 class="peer-validator__title section-title">{{ $t("peer-validate__title") }}</h3>
		<div class="peer-validator__fingerprint-block-list">
			<div class="peer-validator__fingerpnts-block">
				<div class="peer-validator__fingerprint-title">{{ $t("peer-validate__own-key") }}</div>
				<div class="peer-validator__fingerprint-emoji-set">
					<tw-emoji :str="ownKey.emojiSet.join(' ')" />
				</div>
				<div class="peer-validator__fingerprint-hexa">{{ ownKey.hexa }}</div>
			</div>
			<div class="peer-validator__fingerpnts-block">
				<div class="peer-validator__fingerprint-title">{{ $t("peer-validate__peer-key") }}</div>
				<div class="peer-validator__fingerprint-emoji-set">
					<tw-emoji :str="peerKey.emojiSet.join(' ')" />
				</div>
				<div class="peer-validator__fingerprint-hexa">{{ peerKey.hexa }}</div>
			</div>
		</div>

		<p class="peer-validator__label">{{ $t("peer-validate__label") }}</p>
		<p class="peer-validator__info">{{ $t("peer-validate__more-info") }}</p>

		<div class="peer-validator__submit-block">
			<div class="peer-validator__submit-button">
				<button class="btn btn-danger" type="button">{{ $t("peer-validate__refuse-action") }}</button>
			</div>
			<div class="peer-validator__submit-button">
				<button class="btn" type="button">{{ $t("peer-validate__ok-action") }}</button>
			</div>
		</div>
		<div v-show="false" class="peer-validator__error-block">
			Une erreur est surveune
		</div>
	</form>
</template>

<script lang="ts">
import { defineComponent } from 'vue';

export default defineComponent({
	name: "PeerValidator",
	data() {
		return {
			ownKey: {
				emojiSet: ["💕", "😒", "👌"],
				hexa: "XXXX-YYYY-ZZZZ-TTTT",
			},
			peerKey: {
				emojiSet: ["🤪", "🤮", "😊"],
				hexa: "AAAA-BBBB-CCCC-DDDD",
			},
		};
	},
})
</script>

<style scoped>
.peer-validator__label {
	font-size: 1.3rem;
}

.peer-validator__info {
	font-size: 0.8rem;
	font-style: italic;
	text-align: center;
}

.peer-validator__fingerpnts-block {
	border-left: solid 0.3rem rgba(0, 0, 0, 0.2);
	padding: 0 1rem;
	margin: 1.5rem 0;
}

.peer-validator__fingerprint-hexa,
.peer-validator__fingerprint-emoji-set {
	text-align: center;
}

.peer-validator__fingerprint-emoji-set {
	font-size: 3rem;
	margin-bottom: 1rem;
}

.peer-validator__fingerprint-hexa {
	color: rgba(51, 52, 71, 0.6);
	font-size: 0.8rem;
}

.peer-validator__submit-block {
	display: flex;
	justify-content: space-evenly;
}

.peer-validator__submit-button {
	margin: 0 1rem;
}
</style>
