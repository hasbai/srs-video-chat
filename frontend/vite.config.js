import { defineConfig } from 'vite'
import Vue from '@vitejs/plugin-vue'
import path from 'path'

export default defineConfig({
  resolve: {
    alias: {
      '@': path.resolve(__dirname, 'src'),
      '@c': path.resolve(__dirname, 'src/components'),
    },
  },
  base: '/',
  plugins: [Vue()],
  define: {
    'process.env': {},
  },
  server: {
    proxy: {
      '/ws': {
        target: 'http://localhost:8000/',
        changeOrigin: true,
        ws: true,
        secure: false,
      },
    },
  },
})
