<!-- Hello world: prerendered, no API call. Delete it when the real work starts. -->
<script lang="ts">
  // Internal links go through resolve(): it adds the base path and checks the route
  // exists (the lint rule svelte/no-navigation-without-resolve requires it).
  import { resolve } from "$app/paths";
</script>

<h1>About</h1>
<p><a href={resolve("/")}>Home</a></p>
