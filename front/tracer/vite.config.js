import { loadEnv, defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import path from 'path'

const CWD = process.cwd()
// env 配置文件
const BASE_ENV_CONFIG = loadEnv('', CWD)
const DEV_ENV_CONFIG = loadEnv('development', CWD)
const PROD_ENV_CONFIG = loadEnv('production', CWD)

// https://vitejs.dev/config/
export default defineConfig(({command, mode}) => {
  const TARGET_ENV_CONFIG = loadEnv(mode, CWD)
  console.info('vite config', {
    command,
    mode,
    TARGET_ENV_CONFIG
  })

  return {
    plugins: [vue()],
    root: path.resolve(__dirname),
    // 将src目录配置别名为 /@ 方便在项目中导入src目录下的文件
    resolve: {
      alias: {
        '/@': path.resolve(__dirname, './src'),
        '/#': path.resolve(__dirname, './public')
      }
    },
    // 引入第三方的配置
    optimizeDeps: {
      include: []
    },
    server: {
      host: TARGET_ENV_CONFIG.VITE_HOST,
      port: TARGET_ENV_CONFIG.VITE_POST,
      open: true,
      strictPort: false,
      https: false,
      // 反向代理
      proxy: {
        'api': {
          target: 'http://127.0.0.1:8080',
          changeOrigin: true,
          rewrite: (path) => path.replace(/^\/api/, '')
        }
      }
    },
    // 生产模式打包配置
    build: {
      outDir: 'dist',
    }
  }
})
