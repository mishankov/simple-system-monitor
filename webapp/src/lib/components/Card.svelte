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
	h2 {
		color: var(--component-card-header-color);
	}

	.card {
		display: flex;
		flex-direction: column;
		align-items: center;
		gap: 5px;

		padding: 10px;

		border-radius: 10px;
		border: var(--component-card-backgroud-border);
		box-shadow:
			0 2px 4px 0 rgba(0, 0, 0, 0.2),
			0 3px 6px 0 rgba(0, 0, 0, 0.19);

		background-color: var(--component-card-backgroud-color);
	}

	.content {
		display: flex;
		flex-direction: column;
		align-items: center;
	}
</style>
