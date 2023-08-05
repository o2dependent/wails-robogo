<script lang="ts">
	import { onMount } from "svelte";
	import {
		EventsEmit,
		EventsOff,
		EventsOn,
	} from "../../wailsjs/runtime/runtime";

	let isBouncing = false;
	onMount(() => {
		EventsOn("bounce_end", () => {
			isBouncing = false;
		});
		return () => {
			EventsOff("bounce_end");
		};
	});
	const startBounce = () => {
		EventsEmit("bounce");
		isBouncing = true;
	};
</script>

<div>
	{#if isBouncing}
		<div class="w-full h-full fixed top-0 left-0">
			<h1>Press <kbd>esc</kbd> to end bounce.</h1>
		</div>
	{:else}
		<button class="btn" on:click={startBounce} type="button"
			>Start Bounce</button
		>
	{/if}
</div>
