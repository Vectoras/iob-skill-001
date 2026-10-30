<!--
  Hello world: each private API's message, shown by a component. Delete it when the
  real work starts.
-->
<script lang="ts">
  import { resolve } from "$app/paths";
  import { HelloWorld } from "ui-components";
  import type { PageProps } from "./$types";

  let { data }: PageProps = $props();
</script>

<h1>hello-019</h1>
{#each data.helloWorlds as helloWorld, i (i)}
  <HelloWorld message={helloWorld.message} />
{/each}
<p><a href={resolve("/about")}>About</a></p>
