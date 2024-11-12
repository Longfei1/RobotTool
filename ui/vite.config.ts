import { fileURLToPath, URL } from 'node:url'

import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import { viteCommonjs } from '@originjs/vite-plugin-commonjs'
import vueSetUp from 'vite-plugin-vue-setup-extend'
//import legacy from '@vitejs/plugin-legacy'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    vue(),
    viteCommonjs(),
    vueSetUp(),
    // legacy({
    //   targets: ["ie>=11"],
    //   additionalLegacyPolyfills: ["regenerator-runtime/runtime"],
    // }),
  ],
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url))
    }
  },
  base: './',
  build: {
    
  }
})