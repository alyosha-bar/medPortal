import { defineConfig, loadEnv } from 'vite'
import react from '@vitejs/plugin-react'
import type { UserConfigExport, ConfigEnv } from 'vite'

export default ({ mode }: ConfigEnv): UserConfigExport => {
  const env = loadEnv(mode, process.cwd(), '')

  return defineConfig({
    plugins: [react()],
    server: {
      proxy:
        mode === 'development'
          ? {
              '/api': {
                target: env.VITE_SERVER_URL,
                changeOrigin: true,
                rewrite: (path) => path.replace(/^\/api/, ''),
              },
            }
          : undefined,
    },
  })
  }
