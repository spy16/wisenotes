<script>
	import { appConf } from '$lib';
	import '../app.postcss';
</script>

<svelte:head>
	<title>WiseNotes</title>
</svelte:head>

<div data-theme={$appConf.theme} class="w-full h-full overflow-hidden">
	<slot />
</div>
