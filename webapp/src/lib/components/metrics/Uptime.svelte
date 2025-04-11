<script>
	import NumberFlow from '@number-flow/svelte';
	import { onDestroy, onMount } from 'svelte';

	// Uptime
	let uptimeData = $state({ days: 0, hours: 0, minutes: 0, seconds: 0 });
	/**
	 * @type {WebSocket}
	 */
	let socketUptime;

	onMount(() => {
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
		if (socketUptime) {
			socketUptime.close();
		}
	});
</script>

<div class="container">
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
</div>

<style>
	* {
		font-family: monospace;
	}

	.container {
		display: flex;
		flex-direction: column;
		align-items: center;
		gap: 5px;
	}
</style>
