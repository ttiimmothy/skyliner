import { defineConfig } from 'vitest/config';

export default defineConfig({
	test: {
		include: ['src/**/*.{test,spec}.{js,ts}'],
		environment: 'jsdom',
		setupFiles: ['src/test/setup.ts'],
		globals: true,
		coverage: {
			provider: 'v8',
			reporter: ['text', 'json', 'html'],
			exclude: [
				'node_modules/',
				'src/test/',
				'**/*.d.ts',
				'**/*.config.{js,ts}',
				'**/app.html'
			]
		}
	},
	resolve: {
		alias: {
			'$app': new URL('./src/app', import.meta.url).pathname,
			'$lib': new URL('./src/lib', import.meta.url).pathname
		}
	}
});
