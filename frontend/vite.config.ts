import { paraglideVitePlugin } from '@inlang/paraglide-js';
import tailwindcss from "@tailwindcss/vite";
import { sveltekit } from '@sveltejs/kit/vite';
import { defineConfig } from 'vite';

export default defineConfig({
	plugins: [
		tailwindcss(),
		sveltekit(),
		paraglideVitePlugin({
			project: './project.inlang',
			// Write generated paraglide artifacts to a build folder instead of src
			// This prevents the plugin from overwriting or replacing files inside `src/lib`
			// (which can cause JSONs to be converted into generated JS files in-source).
			outdir: './src/lib/paraglide'
		})
	],
	server: {
		fs: {
			allow: ['wailsjs']
		}
	}
});
