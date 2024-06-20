<script lang="ts">
	import { goto } from '$app/navigation';
	import { pb } from '$lib/pb-client';
	import { onMount } from 'svelte';

	onMount(() => {
		if (!pb.authStore.isValid) {
			goto('/login-device');
		}
		if (pb.authStore.model?.collectionName != 'devices') {
			goto('/dashboard');
		}
	});
</script>

<slot />
