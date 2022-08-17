declare module "twemoji-parser" {
  export function parse(
    str: string
  ): Array<{ url: string; indices: Array<number>; text: string; type: string }>;
}
