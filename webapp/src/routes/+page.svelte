<script>
	import { onDestroy, onMount } from 'svelte';
	import NumberFlow from '@number-flow/svelte';
	import LoadLine from '$lib/components/LoadLine.svelte';
	import Card from '$lib/components/Card.svelte';

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

	// Uptime
	let uptimeData = $state({ days: 0, hours: 0, minutes: 0, seconds: 0 });
	/**
	 * @type {WebSocket}
	 */
	let socketUptime;

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

		socketUptime = new WebSocket(`ws://${location.host}/uptime`);

		socketUptime.addEventListener('message', function (event) {
			if (event.data.length > 0) {
				let uptimeSeconds = ~~JSON.parse(event.data).uptime;
				uptimeData.seconds = uptimeSeconds % 60;
				uptimeData.minutes = ~~(uptimeSeconds / 60) % 60;
				uptimeData.hours = ~~(uptimeSeconds / 60 / 60) % 24;
				uptimeData.days = ~~(uptimeSeconds / 60 / 60 / 24);
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

		if (socketUptime) {
			socketUptime.close();
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
	<span>
		{#if uptimeData.days != 0}
			<NumberFlow
				value={uptimeData.days}
				format={{ style: 'unit', unit: 'day', unitDisplay: 'long' }}
			/>
		{/if}

		{#if uptimeData.hours != 0}
			<NumberFlow
				value={uptimeData.hours}
				format={{ style: 'unit', unit: 'hour', unitDisplay: 'long' }}
			/>
		{/if}

		{#if uptimeData.minutes != 0}
			<NumberFlow
				value={uptimeData.minutes}
				format={{ style: 'unit', unit: 'minute', unitDisplay: 'long' }}
			/>
		{/if}

		{#if uptimeData.seconds != 0}
			<NumberFlow
				value={uptimeData.seconds}
				format={{ style: 'unit', unit: 'second', unitDisplay: 'long' }}
			/>
		{/if}
	</span>
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
