<!--
  Hello world: proves a component with plain CSS, styled from the stylesheet its
  consumers import (a custom property), used across packages. Delete it, with its
  story, when the real work starts.
-->
<script lang="ts">
  let { message }: { message: string } = $props();
</script>

<p class="hello-world">{message}</p>

<style>
  .hello-world {
    color: var(--hello-world-color);
    font-weight: bold;
  }
</style>
