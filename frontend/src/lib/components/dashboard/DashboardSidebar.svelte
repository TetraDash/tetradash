<script>
	import { page } from '$app/stores';
	import {
		Sidebar,
		SidebarBrand,
		SidebarGroup,
		SidebarItem,
		SidebarWrapper
	} from 'flowbite-svelte';
	import {
		ChartPieSolid,
		GridSolid,
		MobilePhoneSolid,
		ArrowLeftToBracketOutline
	} from 'flowbite-svelte-icons';
	let spanClass = 'flex-1 ms-3 whitespace-nowrap';
	$: activeUrl = $page.url.pathname;

	let site = {
		name: 'Tetra Dash',
		href: '/',
		img: '/tetradash-logo.svg'
	};

	function handleSignOut() {}
</script>

<div class="flex h-screen w-screen">
	<Sidebar {activeUrl} class="h-full bg-gray-50">
		<SidebarWrapper>
			<SidebarGroup>
				<SidebarBrand {site} />
				<SidebarItem label="Dashboard" href="/dashboard">
					<svelte:fragment slot="icon">
						<ChartPieSolid
							class="h-6 w-6 text-gray-500 transition duration-75 group-hover:text-gray-900 dark:text-gray-400 dark:group-hover:text-white"
						/>
					</svelte:fragment>
				</SidebarItem>
				<SidebarItem label="Devices" {spanClass} href="/dashboard/devices">
					<svelte:fragment slot="icon">
						<MobilePhoneSolid
							class="h-6 w-6 text-gray-500 transition duration-75 group-hover:text-gray-900 dark:text-gray-400 dark:group-hover:text-white"
						/>
					</svelte:fragment>
					<!-- <svelte:fragment slot="subtext">
						<span
							class="ms-3 inline-flex items-center justify-center rounded-full bg-gray-200 px-2 text-sm font-medium text-gray-800 dark:bg-gray-700 dark:text-gray-300"
						>
							Pro
						</span>
					</svelte:fragment> -->
				</SidebarItem>
				<SidebarItem label="Layouts" {spanClass} href="/dashboard/layouts">
					<svelte:fragment slot="icon">
						<GridSolid
							class="h-6 w-6 text-gray-500 transition duration-75 group-hover:text-gray-900 dark:text-gray-400 dark:group-hover:text-white"
						/>
						<!-- <MailBoxSolid
							class="h-6 w-6 text-gray-500 transition duration-75 group-hover:text-gray-900 dark:text-gray-400 dark:group-hover:text-white"
						/> -->
					</svelte:fragment>
					<!-- <svelte:fragment slot="subtext">
						<span
							class="ms-3 inline-flex h-3 w-3 items-center justify-center rounded-full bg-primary-200 p-3 text-sm font-medium text-primary-600 dark:bg-primary-900 dark:text-primary-200"
						>
							3
						</span>
					</svelte:fragment> -->
				</SidebarItem>
				<!-- <SidebarItem
					label="Sidebar"
					href="/docs/components/sidebar"
					active={activeUrl === '/docs/components/sidebar'}
				>
					<svelte:fragment slot="icon">
						<UserSolid
							class="h-6 w-6 text-gray-500 transition duration-75 group-hover:text-gray-900 dark:text-gray-400 dark:group-hover:text-white"
						/>
					</svelte:fragment>
				</SidebarItem> -->
				<!-- <SidebarItem label="Sign In">
					<svelte:fragment slot="icon">
						<ArrowRightToBracketOutline
							class="h-6 w-6 text-gray-500 transition duration-75 group-hover:text-gray-900 dark:text-gray-400 dark:group-hover:text-white"
						/>
					</svelte:fragment>
				</SidebarItem> -->
				<SidebarItem
					label="Sign Out"
					on:click={() => {
						handleSignOut();
					}}
				>
					<svelte:fragment slot="icon">
						<ArrowLeftToBracketOutline
							class="h-6 w-6 text-gray-500 transition duration-75 group-hover:text-gray-900 dark:text-gray-400 dark:group-hover:text-white"
						/>
					</svelte:fragment>
				</SidebarItem>
			</SidebarGroup>
		</SidebarWrapper>
	</Sidebar>
	<div class="w-full">
		<slot />
	</div>
</div>
