<script>
	let { header, content } = $props();
</script>

<div class="card">
	<h2>{header}</h2>
	<div class="content">
		{@render content()}
	</div>
</div>

<style>
	.card {
		display: flex;
		flex-direction: column;
		align-items: center;
		gap: 5px;

		padding: 10px;

		border-radius: 10px;
		border: 1px solid lightgray;
		box-shadow:
			0 2px 4px 0 rgba(0, 0, 0, 0.2),
			0 3px 6px 0 rgba(0, 0, 0, 0.19);
	}

	.content {
		display: flex;
		flex-direction: column;
		align-items: center;
	}
</style>
