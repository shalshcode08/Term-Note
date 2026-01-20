import { defineConfig } from 'astro/config';
import tailwind from '@astrojs/tailwind';

export default defineConfig({
  integrations: [tailwind()],
  site: 'https://termnote.dev',
  output: 'static',
  build: {
    inlineStylesheets: 'auto',
  },
});
