<script>
	import { onDestroy, onMount } from 'svelte';
	import NumberFlow from '@number-flow/svelte';
	import LoadLine from '$lib/components/LoadLine.svelte';
	import Card from '$lib/components/Card.svelte';
	import Uptime from '$lib/components/metrics/Uptime.svelte';

	// Mem
	let memUsed = $state(0.0);
	let memTotal = $state(0.0);
	let memPercent = $state(0.0);
	/**
	 * @type {WebSocket}
	 */
	let socketMem;

	// CPU
	let cpus = $state([{ id: '', load: 0 }]);
	/**
	 * @type {WebSocket}
	 */
	let socketCpu;



	onMount(() => {
		socketMem = new WebSocket(`ws://${location.host}/meminfo`);

		socketMem.addEventListener('message', function (event) {
			if (event.data.length > 0) {
				let data = JSON.parse(event.data);

				memUsed = (data.mem_total - data.mem_available) / 1024 / 1024;
				memTotal = data.mem_total / 1024 / 1024;
				memPercent =
					(data.mem_total - data.mem_available) / 1024 / 1024 / (data.mem_total / 1024 / 1024);
			}
		});

		socketCpu = new WebSocket(`ws://${location.host}/cpuinfo`);

		socketCpu.addEventListener('message', function (event) {
			if (event.data.length > 0) {
				cpus = JSON.parse(event.data);
			}
		});

		
	});

	onDestroy(() => {
		if (socketMem) {
			socketMem.close();
		}

		if (socketCpu) {
			socketCpu.close();
		}
	});

	const numberFormat = {
		maximumFractionDigits: 2
	};
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
	<LoadLine percent={memPercent * 100} />
	<span
		><NumberFlow value={memUsed} format={numberFormat} /> / <NumberFlow
			value={memTotal}
			format={numberFormat}
		/> Gb (<NumberFlow
			value={memPercent}
			format={{ style: 'percent', maximumFractionDigits: 2 }}
		/>)</span
	>
{/snippet}

{#snippet cpu()}
	<div class="cpu-list">
		{#each cpus as cpu (cpu.id)}
			{#if cpu.id != ''}
				<div class="cpu-line">
					<span class="cpu-id">{cpu.id}</span>
					<LoadLine percent={cpu.load * 100} />
					<NumberFlow value={cpu.load} format={{ style: 'percent', maximumFractionDigits: 2 }} />
				</div>
			{/if}
		{/each}
	</div>
{/snippet}

{#snippet uptime()}
	<Uptime/>
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

	.cpu-list {
		display: flex;
		flex-direction: column;
		gap: 5px;
	}

	.cpu-line {
		/* display: flex;
		flex-direction: row;
		gap: 10px;
		align-items: center; */

		display: grid;
		align-items: center;
		grid-template-columns: repeat(3, 1fr);
		gap: 10px;
	}

	.cpu-id {
		text-align: right;
	}
</style>
