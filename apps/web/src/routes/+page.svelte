<!--
  Hello world: a message shown by a component (this app calls no private API).
  Delete it when the real work starts.
-->
<script lang="ts">
  import { resolve } from "$app/paths";
  import { HelloWorld } from "../lib/components";
</script>

<h1>hello-016</h1>
<HelloWorld message="Hello world!" />
<p><a href={resolve("/about")}>About</a></p>
