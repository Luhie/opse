import {defineConfig} from 'vite'
import vue from '@vitejs/plugin-vue'
import path from 'path'
// import { fileURLToPath, URL } from 'node:url'

// https://vitejs.dev/config/
export default defineConfig({
  base: "./",
  plugins: [vue()],
  resolve: {
    alias: {
      '@': path.resolve(__dirname,'src'),
      '@wails': path.resolve(__dirname,'wailsjs'),
      // '@': fileURLToPath(new URL('./src', import.meta.url))
    }
  }

})
