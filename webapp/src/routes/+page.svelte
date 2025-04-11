<script>
	import Card from '$lib/components/Card.svelte';
	import Uptime from '$lib/components/metrics/Uptime.svelte';
	import CpuInfo from '$lib/components/metrics/CPUInfo.svelte';
	import MemInfo from '$lib/components/metrics/MemInfo.svelte';
</script>

<svelte:head>
	<title>Simple system monitor</title>
</svelte:head>

<div class="container">
	<div class="cards">
		<Card header="RAM" content={ram} />
		<Card header="CPU" content={cpu} />
		<Card header="Uptime" content={uptime} />
	</div>
</div>

{#snippet ram()}
	<MemInfo />
{/snippet}

{#snippet cpu()}
	<CpuInfo />
{/snippet}

{#snippet uptime()}
	<Uptime />
{/snippet}

<style>
	:global(*) {
		font-family: monospace;
	}

	.container {
		height: 100vh;
		width: 100vw;
		display: flex;
		justify-content: center;
	}

	.cards {
		padding: 25px;
		width: 450px;

		display: flex;
		flex-direction: column;
		gap: 10px;
	}
</style>
