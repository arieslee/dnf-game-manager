const path = require('path')
function resolve(dir) {
  return path.join(__dirname, dir)
}
const isDev = process.env.NODE_ENV === 'development'


module.exports = {
  // 基本路径
  publicPath: process.env.VUE_APP_PUBLIC_PATH,
  // 输出文件目录
  outputDir: 'dist',
  // assetsDir: 'static',
  lintOnSave: false,
  productionSourceMap: false,
  devServer: {
    open: true, // 自动启动浏览器
    port: 9998,
    // 浏览器弹出错误
    overlay: {
      warnings: false,
      errors: true
    }
  },
  configureWebpack: {
    // provide the app's title in webpack's name field, so that
    // it can be accessed in index.html to inject the correct title.
    // name: name,
    resolve: {
      alias: {
        '@': resolve('src')
      }
    },
    // 发布时关闭警告
    performance: {
      hints: false
    }
  },
  // webpack配置
  // see https://github.com/vuejs/vue-cli/blob/dev/docs/webpack.md
  chainWebpack(config) {
    /**
     * 删除懒加载模块的 prefetch preload，降低带宽压力
     * https://cli.vuejs.org/zh/guide/html-and-static-assets.html#prefetch
     * https://cli.vuejs.org/zh/guide/html-and-static-assets.html#preload
     * 而且预渲染时生成的 prefetch 标签是 modern 版本的，低版本浏览器是不需要的
     */
    config.plugins.delete('prefetch').delete('preload')

    config
    // https://webpack.js.org/configuration/devtool/#development
    .when(isDev, config => config.devtool('cheap-source-map'))

    config.when(!isDev, config => {
      config.optimization.splitChunks({
        chunks: 'all',
        cacheGroups: {
          libs: {
            name: 'chunk-libs',
            test: /[\\/]node_modules[\\/]/,
            priority: 10,
            chunks: 'initial' // only package third parties that are initially dependent
          },
          commons: {
            name: 'chunk-commons',
            test: resolve('src/components'), // can customize your rules
            minChunks: 3, //  minimum common number
            priority: 5,
            reuseExistingChunk: true
          }
        }
      })
      config.optimization.runtimeChunk('single')
    })
  },
  // css相关配置
  css: {
    // 是否使用css分离插件 ExtractTextPlugin
    extract: true,
    sourceMap: false,
    // css预设器配置项
    loaderOptions: {},
    // 启用 CSS modules for all css / pre-processor files.
    // modules: false,
    requireModuleExtension: true
  },
  // use thread-loader for babel & TS in production build
  // enabled by default if the machine has more than 1 cores
  parallel: require('os').cpus().length > 1,
  transpileDependencies: [
    'vuetify'
  ]
}
