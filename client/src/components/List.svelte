<script>
  import Item from "./Item.svelte";
  import { onMount } from "svelte";
  let golies = [];
  onMount(async () => {
    const res = await fetch("http://localhost:8000/goly");
    golies = await res.json();
  });
</script>

{#each golies as goly}
  <Item {goly} />
{/each}
