<script lang="ts">
	import { getProfiles, type Profile } from '$lib';
	import { onMount } from 'svelte';
	import MaterialSymbolsSettingsOutline from '~icons/material-symbols/settings-outline';

	let profiles: Profile[] = [];
	onMount(async () => {
		profiles = await getProfiles();
	});
</script>

<div class="flex flex-col gap-2 h-screen p-2">
	<div class="flex flex-row justify-between">
		<div class="flex flex-row gap-1">
			<select class="select select-bordered select-sm" name="profile" id="profile">
				{#each profiles as profile}
					<option value={profile.id}>{profile.emoji} {profile.name}</option>
				{/each}
			</select>

			<a href="/write" class="btn btn-primary btn-sm">+ New</a>
		</div>

		<a href="/settings" class="btn btn-ghost btn-sm p-1 normal-case">
			<MaterialSymbolsSettingsOutline />
			<span class="hidden md:inline-block">Settings</span>
		</a>
	</div>

	<div class="flex-grow p-2" style="overflow-y: auto;">
		<div class="h-full flex flex-col gap-2">
			<div class="flex-grow border border-dashed">
				<h1>your notes will appear here.</h1>
			</div>
			<input
				type="text"
				name="message"
				id="message"
				class="input input-bordered"
				placeholder="Enter your question here..."
				autocomplete="off"
			/>
		</div>
	</div>
</div>
