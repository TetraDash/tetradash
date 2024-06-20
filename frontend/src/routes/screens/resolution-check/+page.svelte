<script lang="ts">
	import { Heading, Button, Badge } from 'flowbite-svelte';
	import { ArrowLeftOutline, ExpandOutline, CompressOutline } from 'flowbite-svelte-icons';
	import { onMount } from 'svelte';

	var width: number = 0;
	var height: number = 0;

	let fullscreen: boolean = false;

	onMount(() => {
		width = window.innerWidth;
		height = window.innerHeight;
		fullscreen = window.innerHeight == screen.height;

		addEventListener('resize', (event) => {
			width = window.innerWidth;
			height = window.innerHeight;
			fullscreen = window.innerHeight == screen.height;
		});
	});

	function handleFullscreen() {
		document.documentElement.requestFullscreen();
	}

	function handleExitFullscreen() {
		document.exitFullscreen();
	}
</script>

<div class="flex h-screen w-screen flex-col items-center justify-center space-y-5 text-center">
	<Heading
		>{width}<Badge large>W</Badge>&nbsp;&nbsp;X&nbsp;&nbsp;&nbsp;{height}<Badge large>H</Badge
		></Heading
	>

	{#if fullscreen}
		<Button on:click={handleExitFullscreen}>
			<CompressOutline class="me-2 h-5 w-5" />
			Exit fullscreen
		</Button>
	{:else}
		<Button on:click={handleFullscreen}>
			<ExpandOutline class="me-2 h-5 w-5" />
			Fullscreen
		</Button>
	{/if}

	<Button href="/screens" outline>
		<ArrowLeftOutline class="me-2 h-5 w-5" />
		Back
	</Button>
</div>
