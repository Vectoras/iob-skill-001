<script module lang="ts">
  // Hello world: proves story tests run (in Chromium, with the accessibility check)
  // and that the stylesheet reaches Storybook. Delete it with HelloWorld.svelte.
  import { defineMeta } from "@storybook/addon-svelte-csf";
  import { expect } from "storybook/test";
  import { HelloWorld } from "ui-components";

  const { Story } = defineMeta({
    title: "Hello world",
    component: HelloWorld,
    args: { message: "Hello world!" },
  });
</script>

<Story
  name="Default"
  play={async ({ canvas }) => {
    const message = await canvas.findByText("Hello world!");
    // Without the stylesheet the text would be black: this colour comes only from it.
    await expect(getComputedStyle(message).color).toBe("rgb(102, 51, 153)");
  }}
/>
