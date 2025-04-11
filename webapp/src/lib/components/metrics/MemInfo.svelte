<script>
	import NumberFlow from '@number-flow/svelte';
	import { onMount, onDestroy } from 'svelte';
	import LoadLine from '$lib/components/LoadLine.svelte';

	// Mem
	let memUsed = $state(0.0);
	let memTotal = $state(0.0);
	let memPercent = $state(0.0);
	/**
	 * @type {WebSocket}
	 */
	let socketMem;

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
	});

	onDestroy(() => {
		if (socketMem) {
			socketMem.close();
		}
	});
</script>

<LoadLine percent={memPercent * 100} />
<span
	><NumberFlow value={memUsed} format={{ maximumFractionDigits: 2 }} /> / <NumberFlow
		value={memTotal}
		format={{ maximumFractionDigits: 2 }}
	/> Gb (<NumberFlow
		value={memPercent}
		format={{ style: 'percent', maximumFractionDigits: 2 }}
	/>)</span
>

<style>
	* {
		font-family: monospace;
	}
</style>
