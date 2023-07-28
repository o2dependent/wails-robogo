<script lang="ts">
	import { route } from "./router";
	import Bounce from "./screens/Bounce.svelte";
	import Home from "./screens/Home.svelte";
	import NotFound from "./screens/NotFound.svelte";
</script>

{#if $route.name === "Home"}
	<Home />
{:else if $route.name === "Bounce"}
	<Bounce />
{:else}
	<NotFound />
{/if}
