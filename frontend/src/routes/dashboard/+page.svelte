<script lang="ts">
	import DashboardSidebar from '$lib/components/dashboard/DashboardSidebar.svelte';
	import { Heading, Hr } from 'flowbite-svelte';
	import { Section } from 'flowbite-svelte-blocks';
</script>

<svelte:head>
	<title>Dashboard - Tetra Dash</title>
</svelte:head>

<DashboardSidebar>
	<Section>
		<Heading>Dashboard</Heading>
		<Hr />
	</Section>
</DashboardSidebar>
