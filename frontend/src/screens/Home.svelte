<script lang="ts">
	import { goto } from "../router";
</script>

<div class="flex flex-col gap-4">
	<button class="btn" type="button" on:click={() => goto("Bounce")}>
		Bounce
	</button>
</div>
