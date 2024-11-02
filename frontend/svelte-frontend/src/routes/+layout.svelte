<script>
	import '../app.css';
</script>

<div class="app">
	<main>
		<slot />
	</main>

	<footer>
		<p>Made in Vim by Julian Orlowski</p>
	</footer>
</div>

<style>
	.app {
		min-height: 100vh;
	}

	main {
		padding: 1rem;
		width: 100%;
		margin: 0 auto;
		box-sizing: border-box;
	}

	footer {
		display: flex;
		flex-direction: column;
		justify-content: center;
		align-items: center;
		padding: 12px;
	}

	@media (min-width: 480px) {
		footer {
			padding: 12px 0;
		}
	}
</style>
