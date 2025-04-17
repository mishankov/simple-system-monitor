import adapter from '@sveltejs/adapter-static';

/** @type {import('@sveltejs/kit').Config} */
const config = {
	kit: {
		adapter: adapter({ pages: '../cmd/server/build' }),
		prerender: {
			handleHttpError: ({ path, message }) => {
				if (path === '/user-assets/user.css') {
					console.log('Ignoring error about /user-assets/vars.css');
					return;
				}
				throw new Error(message);
			}
		}
	}
};

export default config;
