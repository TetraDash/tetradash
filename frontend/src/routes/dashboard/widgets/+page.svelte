<script lang="ts">
	import DashboardSidebar from '$lib/components/dashboard/DashboardSidebar.svelte';
	import { Button, Heading, Hr } from 'flowbite-svelte';
	import { Section } from 'flowbite-svelte-blocks';
	import { PlusOutline } from 'flowbite-svelte-icons';
</script>

<DashboardSidebar>
	<Section>
		<div class="flex">
			<Heading>Widgets</Heading>
			<!-- <Button
				on:click={() => {
					showAddLayoutModal = true;
				}}
			>
				<PlusOutline /> Add
			</Button> -->
		</div>
		<Hr />
	</Section>
</DashboardSidebar>
