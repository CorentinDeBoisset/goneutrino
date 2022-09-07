<template>
	<form @submit.prevent="" class="peer-selector">
		<div class="peer-selector__input-block">
			<div class="peer-selector__input-label section-title">{{ $t("peer-selector__pick-peer") }}</div>
			<input v-model="peerId" placeholder="1234" />
		</div>
		<div class="peer-selector__submit-block">
			<button class="btn" type="submit">{{ $t("peer-selector__submit-action") }}</button>
		</div>
		<div v-show="false" class="peer-selector__error-block">
			Une erreur est surveune
		</div>
	</form>
</template>

<script lang="ts">
import { defineComponent } from 'vue';

export default defineComponent({
	name: "PeerSelector",
	data() {
		return {
			peerId: "",
		};
	},
})
</script>

<style scoped>
.peer-selector {
	margin: 0.5rem 0 0;
}

.peer-selector__input-label {
	font-size: 1.5rem;
	line-height: 1.8rem;
	text-align: left;
}

.peer-selector__input-block {
	display: flex;
	flex-direction: column;
	align-items: center;
}

.peer-selector__input-block input {
	margin: 1rem 0 2rem;
}

.peer-selector__submit-block {
	text-align: right;
	margin: 2rem 0 0;
}

.peer-selector__error-block {
	color: #d94f04;
	margin: 1rem 0 0.5rem;
	text-align: center;
}
</style>
