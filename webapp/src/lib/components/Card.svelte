<script>
	let { header, content } = $props();
</script>

<div class="card">
	<h1>{header}</h1>
	<div class="content">
		{@render content()}
	</div>
</div>

<style>
	.card {
		display: flex;
		flex-direction: column;
		align-items: center;

		padding: 10px;

		border-radius: 10px;
		box-shadow:
			0 4px 8px 0 rgba(0, 0, 0, 0.2),
			0 6px 20px 0 rgba(0, 0, 0, 0.19);
	}

	.content {
		display: flex;
		flex-direction: column;
		align-items: center;
	}
</style>
