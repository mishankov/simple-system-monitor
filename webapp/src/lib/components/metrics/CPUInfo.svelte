<script>
	import NumberFlow from "@number-flow/svelte";
	import { onDestroy, onMount } from "svelte";
	import LoadLine from "$lib/components/LoadLine.svelte";

  	// CPU
	let cpus = $state([{ id: '', load: 0 }]);
	/**
	 * @type {WebSocket}
	 */
	let socketCpu;

  onMount(() => {
    socketCpu = new WebSocket(`ws://${location.host}/cpuinfo`);

		socketCpu.addEventListener('message', function (event) {
			if (event.data.length > 0) {
				cpus = JSON.parse(event.data);
			}
		});
  })

  onDestroy(() => {
    if (socketCpu) {
			socketCpu.close();
		}
  })
</script>

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

<style>
  * {
    font-family: monospace;
  }

  .cpu-list {
		display: flex;
		flex-direction: column;
		gap: 5px;
	}

	.cpu-line {
		display: grid;
		align-items: center;
		grid-template-columns: repeat(3, 1fr);
		gap: 10px;
	}

	.cpu-id {
		text-align: right;
	}
</style>