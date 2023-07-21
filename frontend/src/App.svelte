<script lang="ts">
	import logo from "./assets/images/logo-universal.png";
	import { Greet, HoldKey } from "../wailsjs/go/main/App.js";
	import { EventsEmit, EventsOn } from "../wailsjs/runtime/runtime";
	import { onMount } from "svelte";

	let resultText: string = "Please enter your name below 👇";
	let name: string;

	let isCapturing: boolean = false;
	let capturedKey: string | null = null;

	let isBouncing = false;

	function greet(): void {
		Greet(name).then((result) => (resultText = result));
	}

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
	<img alt="Wails logo" id="logo" src={logo} />
	<div class="result" id="result">{resultText}</div>
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

<style>
	#logo {
		display: block;
		width: 50%;
		height: 50%;
		margin: auto;
		padding: 10% 0 0;
		background-position: center;
		background-repeat: no-repeat;
		background-size: 100% 100%;
		background-origin: content-box;
	}

	.result {
		height: 20px;
		line-height: 20px;
		margin: 1.5rem auto;
	}

	.input-box .btn {
		width: 60px;
		height: 30px;
		line-height: 30px;
		border-radius: 3px;
		border: none;
		margin: 0 0 0 20px;
		padding: 0 8px;
		cursor: pointer;
	}

	.input-box .btn:hover {
		background-image: linear-gradient(to top, #cfd9df 0%, #e2ebf0 100%);
		color: #333333;
	}

	.input-box .input {
		border: none;
		border-radius: 3px;
		outline: none;
		height: 30px;
		line-height: 30px;
		padding: 0 10px;
		background-color: rgba(240, 240, 240, 1);
		-webkit-font-smoothing: antialiased;
	}

	.input-box .input:hover {
		border: none;
		background-color: rgba(255, 255, 255, 1);
	}

	.input-box .input:focus {
		border: none;
		background-color: rgba(255, 255, 255, 1);
	}
</style>
