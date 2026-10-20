<script lang="ts">
  import favicon from "#lib/assets/favicon.svg";
  // The components' stylesheet, from the UI package.
  import "ui-components/styles.css";
  import type { LayoutProps } from "./$types";

  let { children }: LayoutProps = $props();
</script>

<svelte:head>
  <link rel="icon" href={favicon} />
</svelte:head>

{@render children()}
