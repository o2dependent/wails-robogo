<script lang="ts">
	import logo from "./assets/images/logo-universal.png";
	import { Greet, HoldKey } from "../wailsjs/go/main/App.js";
	import { EventsEmit, EventsOn } from "../wailsjs/runtime/runtime";
	import { onMount } from "svelte";

	let isCapturing: boolean = false;
	let capturedKey: string | null = null;

	let isBouncing = false;

	const startBounce = () => {
		EventsEmit("bounce");
		isBouncing = true;
	};

	onMount(() => {
		EventsOn("bounce_end", () => {
			isBouncing = false;
		});
	});

	const captureKey = (e: KeyboardEvent) => {
		capturedKey = e.key;
	};

	const holdKey = async () => {
		if (capturedKey) {
			await HoldKey(capturedKey);
			isCapturing = false;
		}
	};

	$: {
		if (isCapturing) {
			window.addEventListener("keydown", captureKey);
		} else {
			window.removeEventListener("keydown", captureKey);
		}
	}
</script>

<main>
	<div>
		{#if isCapturing}
			<button on:click={holdKey} type="button">
				Hold <kbd>{capturedKey ?? "<NONE_ENTERED>"}</kbd>
			</button>
		{:else}
			<button on:click={() => (isCapturing = true)} type="button">
				Start key capture
			</button>
		{/if}
		<div>
			{#if isBouncing}
				<h1>Press <kbd>esc</kbd> to end bounce.</h1>
			{:else}
				<button on:click={startBounce} type="button">Start Bounce</button>
			{/if}
		</div>
	</div>
</main>
