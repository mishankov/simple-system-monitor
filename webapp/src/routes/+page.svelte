<script>
	import { browser } from "$app/environment";
	import { onMount } from "svelte";
  import NumberFlow from "@number-flow/svelte";

  let used = 0.0
  let total = 0.0
  let percent = 0.0
  let socket

  /**
	 * @param {number} num
	 */
  function round(num) {
    return Math.round(num * 100) / 100

  }

  onMount(() => {
    socket = new WebSocket(`ws://10.0.0.157:4442/meminfo`);

    socket.addEventListener("message", function (event) {
      if (event.data.length > 0) {
        let data = JSON.parse(event.data)

        used = (data.mem_total - data.mem_available) / 1024 / 1024
        total = data.mem_total / 1024 / 1024
        percent = ((data.mem_total - data.mem_available) / 1024 / 1024)/(data.mem_total / 1024 / 1024)
      }
    });
  })

  const numberFormat = {
    maximumFractionDigits: 2
  }

  // $: {
    
  // }

</script>

<div class="container">
  <div class="card">
    <h1>RAM</h1>
    <span><NumberFlow value={used} format={numberFormat}/>/<NumberFlow value={total}  format={numberFormat}/> Gb (<NumberFlow value={percent} format={{style: "percent", maximumFractionDigits: 2}}/>)</span>
  </div>
</div>

<style>
  .container {
    padding: 25px;
  }
  .card {
    display: flex;
    flex-direction: column;
    align-items: center;

    padding: 5px;
    /* background-color: burlywood; */

    border-radius: 10px;
    box-shadow: 0 4px 8px 0 rgba(0, 0, 0, 0.2), 0 6px 20px 0 rgba(0, 0, 0, 0.19);
  }
</style>
