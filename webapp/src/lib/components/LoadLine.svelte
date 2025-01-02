<script>
	let { percent } = $props();
</script>

<div class="outer">
	<div class="inner" style="width: {percent}%"></div>
</div>

<style>
	.outer {
		height: 15px;
		width: 250px;
		background-color: lightgray;
	}

	.inner {
		height: 100%;
		background-color: seagreen;
		transition: width 0.5s;
	}
</style>
