import { sveltekit } from '@sveltejs/kit/vite';
import { defineConfig } from 'vite';

export default defineConfig({
	plugins: [sveltekit()],
	server: {
		proxy: {
			'/meminfo': {
				target: 'http://localhost:4442',
				ws: true
			},
			'/cpuinfo': {
				target: 'http://localhost:4442',
				ws: true
			},
			'/uptime': {
				target: 'http://localhost:4442',
				ws: true
			}
		}
	}
});
