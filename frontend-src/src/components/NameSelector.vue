<template>
	<form @submit.prevent="" class="name-selector">
		<div class="name-selector__input-block">
			<div class="name-selector__input-label section-title">{{ $t("name-selector__hello") }}</div>
			<div class="name-selector__input-label section-title">{{ $t("name-selector__pick-nickname") }}</div>
			<input v-model="nickname" :placeholder="$t('name-selector__nickname-placeholder')" />
		</div>
		<div class="name-selector__submit-block">
			<button class="btn" type="submit">{{ $t("name-selector__submit-action") }}</button>
		</div>
		<div v-show="false" class="name-selector__error-block">
			Une erreur est survenue
		</div>
	</form>
</template>

<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
	name: "NameSelector",
	data() {
		return {
			nickname: "",
		};
	},
});
</script>

<style scoped>
.name-selector {
	margin: 0.5rem 0 0;
}

.name-selector__input-label {
	font-size: 1.5rem;
	line-height: 1.8rem;
	text-align: left;
}

.name-selector__input-block input {
	margin: 1rem 0 2rem;
}

.name-selector__submit-block {
	text-align: right;
	margin: 2rem 0 0;
}

.name-selector__error-block {
	color: #d94f04;
	margin: 1rem 0 0.5rem;
	text-align: center;
}
</style>
