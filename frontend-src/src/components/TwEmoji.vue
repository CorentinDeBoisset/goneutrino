<template>
	<span>
		<span v-for="block in parsedBlocks" :key="blockIdx(block)">
			<span v-if="block.type == 'string'">{{ block.str }}</span>
			<img v-else-if="block.type == 'emoji'" class="tw-emoji" :src="block.url" :alt="block.alt"
				draggable="false" />
		</span>
	</span>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { parse } from "twemoji-parser";

export default defineComponent({
	props: {
		str: {
			type: String,
			default: "",
		},
	},
	computed: {
		parsedBlocks(): Array<
			| { type: "string"; idx: number; str: string }
			| { type: "emoji"; idx: number; alt: string; url: string }
		> {
			const parsedEmojis = parse(this.str);
			if (parsedEmojis.length > 0) {
				const ret: Array<
					| { type: "string"; idx: number; str: string }
					| { type: "emoji"; idx: number; alt: string; url: string }
				> = [];

				if (parsedEmojis[0].indices[0] > 0) {
					ret.push({
						type: "string",
						idx: 0,
						str: this.str.substring(0, parsedEmojis[0].indices[0]),
					});
				}

				for (let i = 0; i < parsedEmojis.length; i++) {
					ret.push({
						type: "emoji",
						idx: parsedEmojis[i].indices[0],
						url: parsedEmojis[i].url,
						alt: parsedEmojis[i].text,
					});
					let nextIdx: number;
					if (i == parsedEmojis.length - 1) {
						nextIdx = this.str.length;
					} else {
						nextIdx = parsedEmojis[i + 1].indices[0];
					}
					if (parsedEmojis[i].indices[1] < nextIdx) {
						ret.push({
							type: "string",
							idx: parsedEmojis[i].indices[1],
							str: this.str.substring(parsedEmojis[i].indices[1], nextIdx),
						});
					}
				}

				return ret;
			}

			return [{ type: "string", idx: 0, str: this.str }];
		},
	},
	methods: {
		blockIdx(
			block:
				| { type: "string"; idx: number; str: string }
				| { type: "emoji"; idx: number; alt: string; url: string }
		): string {
			if (block.type == "string") {
				return block.idx + "-" + block.str;
			}
			return block.idx + "-" + block.url;
		},
	},
});
</script>

<style scoped>
.tw-emoji {
	/* This ensures the emojies are exactli the font size of the parent element */
	height: 1em;

	position: relative;
	top: -0.1rem;
	vertical-align: middle;
	margin: 0 0.1rem;
}
</style>
