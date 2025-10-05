import { sveltekit } from '@sveltejs/kit/vite';
import { defineConfig } from 'vite';

export default defineConfig({
	plugins: [sveltekit()],
  server: {
    proxy: {
      '/providers': 'http://localhost:8080',
      '/formats': 'http://localhost:8080'
    }
  }
});
