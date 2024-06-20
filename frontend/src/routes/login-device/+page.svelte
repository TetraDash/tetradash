<script lang="ts">
	import { Section, Register } from 'flowbite-svelte-blocks';
	import { Button, Checkbox, Label, Input, Alert } from 'flowbite-svelte';
	import { onMount } from 'svelte';
	import { PB_DEVICES_COLLECTION_ID, pb } from '$lib/pb-client';
	import { goto } from '$app/navigation';

	onMount(() => {
		if (pb.authStore.isValid) {
			if (pb.authStore.model?.collectionName == 'devices') {
				goto('/screens');
			} else {
				goto('/dashboard');
			}
		}
	});

	let alertMessage: string = '';
	let username: string = '';
	let password: string = '';

	function handleLogin() {
		// console.log(username, password);
		pb.collection(PB_DEVICES_COLLECTION_ID)
			.authWithPassword(username, password)
			.then((v) => {
				goto('/screens');
			})
			.catch((e) => {
				alertMessage = e.message;
			});
	}
</script>

<svelte:head>
	<title>Sign - Tetra Dash</title>
</svelte:head>

<Section name="login">
	<Register href="/">
		<svelte:fragment slot="top">
			<img class="mr-2 h-8 w-8" src="/tetradash-logo.svg" alt="logo" />
			Tetra Dash
		</svelte:fragment>
		<div class="space-y-4 p-6 sm:p-8 md:space-y-6">
			<form class="flex flex-col space-y-6" action="/" on:submit|preventDefault={handleLogin}>
				<h3 class="p-0 text-xl font-medium text-gray-900 dark:text-white">Device Sign In</h3>
				<Label class="space-y-2">
					<span>User name</span>
					<Input
						type="text"
						name="username"
						placeholder="big-screen"
						required
						bind:value={username}
					/>
				</Label>
				<Label class="space-y-2">
					<span>Password</span>
					<Input
						type="password"
						name="password"
						placeholder="•••••"
						required
						bind:value={password}
					/>
				</Label>
				<!-- <div class="flex items-start">
					<Checkbox>Remember me</Checkbox>
					<a href="/" class="ml-auto text-sm text-blue-700 hover:underline dark:text-blue-500"
						>Forgot password?</a
					>
				</div> -->

				{#if alertMessage.length > 0}
					<Alert color="red">
						{alertMessage}
						<!-- Change a few things up and try submitting again. -->
					</Alert>
				{/if}

				<Button type="submit" class="w-full1" outline>Sign in</Button>
				<p class="text-sm font-light text-gray-500 dark:text-gray-400">
					<a
						href="/login-admin"
						class="font-medium text-primary-600 hover:underline dark:text-primary-500"
						>Log in as Admin</a
					>
				</p>
			</form>
		</div>
	</Register>
</Section>
