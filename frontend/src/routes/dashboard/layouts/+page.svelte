<script lang="ts">
	import DashboardSidebar from '$lib/components/dashboard/DashboardSidebar.svelte';
	import { PB_LAYOUTS_COLLECTION_ID, pb } from '$lib/pb-client';
	import { Button, Heading, Hr, Input, Label, Modal } from 'flowbite-svelte';
	import { Section } from 'flowbite-svelte-blocks';
	import { PlusOutline } from 'flowbite-svelte-icons';

	let showAddLayoutModal: boolean = false;

	let name: string = '';
	let width: number | null;
	let height: number | null;

	function handleCreate() {
		const data = {
			width,
			height
		};

		pb.collection(PB_LAYOUTS_COLLECTION_ID)
			.create(data)
			.then((v) => {})
			.catch((e) => {});
	}
</script>

<svelte:head>
	<title>Layouts - Tetra Dash</title>
</svelte:head>

<DashboardSidebar>
	<Section>
		<div class="flex">
			<Heading>Layouts</Heading>
			<Button
				on:click={() => {
					showAddLayoutModal = true;
				}}
			>
				<PlusOutline /> Add
			</Button>
		</div>
		<Hr />
	</Section>
</DashboardSidebar>

<Modal title="Add layout" bind:open={showAddLayoutModal}>
	<form on:submit|preventDefault={handleCreate}>
		<div class="mb-6">
			<Label for="name" class="mb-2">Name</Label>
			<Input type="text" id="name" placeholder="Youtube analytics" required bind:value={name} />
		</div>
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
</Modal>
