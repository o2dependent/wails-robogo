<script lang="ts">
	import logo from "./assets/images/logo-universal.png";
	import { Greet, HoldKey } from "../wailsjs/go/main/App.js";
	import { EventsEmit, EventsOn } from "../wailsjs/runtime/runtime";
	import { onMount } from "svelte";
	import Sidebar from "./lib/Sidebar.svelte";
	import Router from "./Router.svelte";

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

<div class="flex h-full">
	<Sidebar />
	<Router />
</div>
