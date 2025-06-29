import { defineConfig, loadEnv } from 'vite'
import react from '@vitejs/plugin-react'
import type { UserConfigExport, ConfigEnv } from 'vite'

export default ({ mode }: ConfigEnv): UserConfigExport => {
  // Load env variables based on current mode (e.g., 'development' or 'production')
  const env = loadEnv(mode, process.cwd(), '')

  return defineConfig({
    plugins: [react()],
    server: {
      proxy: {
        '/api': {
          target: env.VITE_SERVER_URL,
          changeOrigin: true,
          rewrite: (path) => path.replace(/^\/api/, ''),
        },
      },
    },
  })
}
