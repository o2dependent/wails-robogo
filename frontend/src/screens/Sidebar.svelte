<script lang="ts">
	import Icon from "src/lib/Icon/Icon.svelte";

	let search = "";
	let pinnedCommands: {
		icon: string;
		name: string;
		commandId: string;
	}[] = [];
</script>

<aside class="flex flex-col gap-2.5 py-4 px-3 w-56">
	<input
		bind:value={search}
		type="Search..."
		placeholder="Type here"
		class="input"
	/>
	<div class="grid grid-cols-4">
		{#each pinnedCommands as command}
			<button type="button" class="btn tooltip" data-tip={command.name}>
				<Icon type={command.icon} />
			</button>
		{/each}
	</div>
</aside>
