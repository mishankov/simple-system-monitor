<script>
	import { browser } from "$app/environment";
	import { onDestroy, onMount } from "svelte";
  import NumberFlow from "@number-flow/svelte";
	import LoadLine from "$lib/components/LoadLine.svelte";
	import Card from "$lib/components/Card.svelte";

  let memUsed = $state(0.0)
  let memTotal = $state(0.0)
  let memPercent = $state(0.0)
  /**
	 * @type {WebSocket}
	 */
  let socketMem

  let cpus = $state([{id: "", load: 0}])
  /**
	 * @type {WebSocket}
	 */
  let socketCpu

  onMount(() => {
    socketMem = new WebSocket(`ws://10.0.0.157:4442/meminfo`);

    socketMem.addEventListener("message", function (event) {
      if (event.data.length > 0) {
        let data = JSON.parse(event.data)

        memUsed = (data.mem_total - data.mem_available) / 1024 / 1024
        memTotal = data.mem_total / 1024 / 1024
        memPercent = ((data.mem_total - data.mem_available) / 1024 / 1024)/(data.mem_total / 1024 / 1024)
      }
    });

    socketCpu = new WebSocket(`ws://10.0.0.157:4442/cpuinfo`);

    socketCpu.addEventListener("message", function (event) {
      if (event.data.length > 0) {
        cpus = JSON.parse(event.data)
      }
    });
  })

  onDestroy(() => {
    if (socketMem) {
      socketMem.close()
    }

    if (socketCpu) {
      socketCpu.close()
    }
  })

  const numberFormat = {
    maximumFractionDigits: 2
  }
</script>

<div class="container">
  <Card header="RAM" content={ram}/>
  <Card header="CPU" content={cpu}/>

  {#snippet ram()}
    <LoadLine percent={memPercent*100}/>
    <span><NumberFlow value={memUsed} format={numberFormat}/> / <NumberFlow value={memTotal}  format={numberFormat}/> Gb (<NumberFlow value={memPercent} format={{style: "percent", maximumFractionDigits: 2}}/>)</span>
  {/snippet}

  {#snippet cpu()}
  <div class="list">
    {#each cpus as cpu (cpu.id)}
    {#if cpu.id != ""}
      <div class="line">
        <span>{cpu.id}</span>
        <LoadLine percent={cpu.load*100}/>
        <NumberFlow value={cpu.load} format={{style: "percent", maximumFractionDigits: 2}}/>
      </div>
    {/if}
    {/each}
  </div>
  {/snippet}
</div>

<style>
  .container {
    padding: 25px;

    display: flex;
    flex-direction: column;
    gap: 10px;
  }

  .list {
    display: flex;
    flex-direction: column;
    gap: 5px;
  }

  .line {
    display: flex;
    flex-direction: row;
    gap: 10px;
    align-items: center;
  }
</style>
