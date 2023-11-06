<script lang="ts">
	import { appConf, getProfiles, type Profile } from '$lib';
	import { onMount } from 'svelte';

	import MaterialSymbolsLightModeOutline from '~icons/material-symbols/light-mode-outline';
	import MaterialSymbolsDarkModeRounded from '~icons/material-symbols/dark-mode-rounded';
	import { themes } from '$lib/constants';

	let profiles: Profile[] = [];
	onMount(async () => {
		profiles = await getProfiles();
	});

	const handleThemeChange = (e: Event) => {
		const theme = (e.target as HTMLSelectElement).value;
		appConf.update((conf) => ({ ...conf, theme }));
	};
</script>

<!-- <div class="flex flex-col h-[100dvh]">
	<div class="flex flex-row justify-between">
		<select class="select select-bordered select-sm" name="profile" id="profile">
			{#each profiles as profile}
				<option value={profile.id}>{profile.emoji} {profile.name}</option>
			{/each}
		</select>

		<select
			class="select select-bordered select-sm"
			name="theme"
			id="theme"
			on:change={handleThemeChange}
		>
			{#each themes as th}
				<option selected={th === $appConf.theme} value={th}>{th}</option>
			{/each}
		</select>
	</div>

	<div class="h-full border-red-200"></div>
</div> -->

<div class="flex flex-col gap-2 h-screen p-2">
	<div class="flex flex-row justify-between">
		<select class="select select-bordered select-sm" name="profile" id="profile">
			{#each profiles as profile}
				<option value={profile.id}>{profile.emoji} {profile.name}</option>
			{/each}
		</select>

		<select
			class="select select-bordered select-sm"
			name="theme"
			id="theme"
			on:change={handleThemeChange}
		>
			{#each themes as th}
				<option selected={th === $appConf.theme} value={th}>{th}</option>
			{/each}
		</select>
	</div>

	<div class="flex-grow p-2" style="overflow-y: auto;">
		<div class="h-full flex flex-col gap-2">
			<div class="flex-grow border border-dashed">
				<h1>your notes will appear here.</h1>
			</div>
			<textarea class="w-full h-auto resize-none textarea textarea-bordered" rows="2" />
		</div>
	</div>
</div>
