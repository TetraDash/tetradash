<script lang="ts">
	import { ComputerSpeakerSolid, UserSolid } from 'flowbite-svelte-icons';
	import { Button, Heading, Card, Hr } from 'flowbite-svelte';
	import { onMount } from 'svelte';
	import { PB_DEVICES_COLLECTION_ID, PB_USER_COLLECTION_ID, pb } from '$lib/pb-client';
	import { goto } from '$app/navigation';
	import { dev } from '$app/environment';

	onMount(() => {
		if (pb.authStore.isValid) {
			if (pb.authStore.model!.collection == PB_USER_COLLECTION_ID) {
				goto('/dashboard');
			} else if (pb.authStore.model!.collection == PB_DEVICES_COLLECTION_ID) {
				goto('/screens');
			}
		}
	});

	function handleSignOut() {
		pb.authStore.clear();
	}
</script>

<svelte:head>
	<title>Sigin - Tetra Dash</title>
</svelte:head>

<div class="flex h-screen items-center justify-center">
	<Card>
		<div class="text-center">
			<Heading>Sigin As</Heading>
		</div>
		<Hr />
		<div class="mx-auto flex w-fit flex-col space-y-5">
			<Button size="xl" href="/login-device">
				<ComputerSpeakerSolid class="me-2 h-6 w-6" />
				Device
			</Button>
			<Button size="xl" href="/login-admin">
				<UserSolid class="me-2 h-6 w-6" />
				Admin
			</Button>

			{#if dev && pb.authStore.isValid}
				<Button outline on:click={handleSignOut}>Sign out</Button>
			{/if}
		</div>
	</Card>
</div>
