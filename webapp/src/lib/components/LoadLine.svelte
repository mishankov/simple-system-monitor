<script>
	let { percent } = $props();
</script>

<div class="outer">
	<div class="inner" style="width: {percent}%"></div>
</div>

<style>
	.outer {
		height: var(--component-loadline-height);
		width: var(--component-loadline-width);
		background-color: var(--component-loadline-empty-color);
	}

	.inner {
		height: 100%;
		background-color: var(--component-loadline-filled-color);
		transition: width 0.5s;
	}
</style>
