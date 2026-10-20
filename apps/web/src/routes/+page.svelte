<!--
  Hello world: the API's message, shown by a component. Delete it when the real work
  starts.
-->
<script lang="ts">
  import { resolve } from "$app/paths";
  import { HelloWorld } from "../lib/components";
  import type { PageProps } from "./$types";

  let { data }: PageProps = $props();
</script>

<h1>hello-008</h1>
<HelloWorld message={data.helloWorld.message} />
<p><a href={resolve("/about")}>About</a></p>
