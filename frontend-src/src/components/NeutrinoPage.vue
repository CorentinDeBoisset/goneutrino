<template>
	<div class="neutrino-page flex-column align-center justify-between">
		<div class="neutrino-page__upper-body flex-row">
			<div class="col-6 mx-auto">
				<header class="neutrino-header flex-column align-center">
					<img class="neutrino-logo" src="../assets/logo_large.svg" alt="Neutrino logo" />
					<h1>
						<span aria-hidden="true">&#8211;</span>
						{{ $t("header-title") }}
						<span aria-hidden="true">&#8211;</span>
					</h1>
					<h2>{{ $t("header-subtitle") }}</h2>
				</header>
				<section class="neutrino-content">
					<component :is="currentComponent"></component>
				</section>
			</div>
		</div>
		<footer class="neutrino-footer">
			<localeChanger />
		</footer>
	</div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import LocaleChanger from "./LocaleChanger.vue";
import NameSelector from "./NameSelector.vue";
import PeerSelector from "./PeerSelector.vue";
import PeerValidator from "./PeerValidator.vue";
import SecretExchange from "./SecretExchange.vue";

export default defineComponent({
	name: "NeutrinoPage",
	data() {
		return {
			currentComponent: "SecretExchange",
		};
	},
	components: {
		LocaleChanger,
		NameSelector,
		PeerSelector,
		PeerValidator,
		SecretExchange,
	},
});
</script>

<style scoped>
.neutrino-header {
	margin: 1rem 0 1rem;
	text-align: center;
}

.neutrino-header h1,
.neutrino-header h2 {
	color: var(--primary-light-color);
	padding: 0;
}

.neutrino-header h1 {
	font-size: 1.3rem;
	letter-spacing: 0.5rem;
	margin: 0.7rem 0 0 0;
	text-transform: uppercase;
	font-style: italic;
	font-weight: 400;
}

.neutrino-header h2 {
	font-size: 1rem;
	margin: 0.5rem 0 0 0;
	font-weight: 200;
	font-style: italic;
}

.neutrino-logo {
	width: 6.5rem;
}

.neutrino-content {
	border: solid 1px var(--secondary-light-color);
	padding: 1rem 2rem;
	margin: 1rem 0 0;
	border-radius: 0.3rem;
	background-color: #fff;
	box-shadow: 0.2rem 0.2rem 0.5rem rgba(0, 0, 0, 0.1);
}

.neutrino-footer {
	padding: 2rem 3rem 2rem 0;
}
</style>
