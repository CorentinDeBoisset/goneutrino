<template>
  <span>
    <span v-for="block in parsedBlocks" :key="blockIdx(block)">
      <span v-if="block.type == 'string'">{{ block.str }}</span>
      <img
        v-else-if="block.type == 'emoji'"
        class="tw-emoji"
        :src="block.url"
        :alt="block.alt"
        draggable="false"
      />
    </span>
  </span>
</template>

<script setup lang="ts">
import { parse } from "twemoji-parser";
import { computed } from "vue";

export interface Props {
  str: string;
}

interface StringBlock {
  type: "string";
  idx: number;
  str: string;
}
interface EmojiBlock {
  type: "emoji";
  idx: number;
  alt: string;
  url: string;
}

const props = withDefaults(defineProps<Props>(), {
  str: () => "",
});

const parsedBlocks = computed((): Array<StringBlock | EmojiBlock> => {
  const parsedEmojis = parse(props.str);
  if (parsedEmojis.length > 0) {
    const ret: Array<StringBlock | EmojiBlock> = [];

    if (parsedEmojis[0].indices[0] > 0) {
      ret.push({
        type: "string",
        idx: 0,
        str: props.str.substring(0, parsedEmojis[0].indices[0]),
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
        nextIdx = props.str.length;
      } else {
        nextIdx = parsedEmojis[i + 1].indices[0];
      }
      if (parsedEmojis[i].indices[1] < nextIdx) {
        ret.push({
          type: "string",
          idx: parsedEmojis[i].indices[1],
          str: props.str.substring(parsedEmojis[i].indices[1], nextIdx),
        });
      }
    }

    return ret;
  }

  return [{ type: "string", idx: 0, str: props.str }];
});

function blockIdx(block: StringBlock | EmojiBlock) {
  if (block.type == "string") {
    return block.idx + "-" + block.str;
  }
  return block.idx + "-" + block.url;
}
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
