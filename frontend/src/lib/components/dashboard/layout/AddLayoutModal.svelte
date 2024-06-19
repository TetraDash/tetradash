<script lang="ts">
	import { A, Button, Checkbox, Input, Label, Modal } from 'flowbite-svelte';

	export let show: boolean = false;
	export let onCreated: (id: string) => any;

	let width: number | null;
	let height: number | null;

	function handleCreate() {
		onCreated('');
	}

	function resetForm() {
		width = null;
		height = null;
	}

	$: if (show) {
		width = null;
		height = null;
	}
	// $: show, resetForm();
</script>

<Modal title="Add layout" bind:open={show}>
	<form on:submit|preventDefault={handleCreate}>
		<div class="mb-6">
			<Label for="width" class="mb-2">Width</Label>
			<Input type="number" id="width" placeholder="1920" required bind:value={width} />
		</div>
		<div class="mb-6">
			<Label for="height" class="mb-2">Height</Label>
			<Input type="number" id="height" placeholder="1080" required bind:value={height} />
		</div>
		<Button type="submit">Submit</Button>
	</form>
	<!-- <svelte:fragment slot="footer">
		<Button on:click={handleCreate}>I accept</Button>
		<Button color="alternative">Decline</Button>
	</svelte:fragment> -->
</Modal>
