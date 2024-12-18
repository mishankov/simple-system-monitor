<script>
	import { browser } from "$app/environment";
	import { onMount } from "svelte";

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

        used = round((data.mem_total - data.mem_available) / 1024 / 1024)
        total = round(data.mem_total / 1024 / 1024)
        percent = round(((data.mem_total - data.mem_available) / 1024 / 1024)/(data.mem_total / 1024 / 1024)*100)
      }
    });
  })

  if (browser) {
      
    }

  // $: {
    
  // }

</script>

<span>{used}/{total} Gb ({percent}%)</span>