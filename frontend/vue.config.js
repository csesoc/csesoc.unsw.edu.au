const ImageminWebpWebpackPlugin = require('imagemin-webp-webpack-plugin');

module.exports = {
  configureWebpack: {
    plugins: [
      // Make sure that the plugin is after any plugins that add images
      new ImageminWebpWebpackPlugin(['src/assets/*'], {
        overrideExtension: true,
        config: [
          {
            test: /\.(jpe?g|png|gif)/,
            options: {
              quality: 75
            }
          }
        ]
      })
    ]
  },
  css: {
    loaderOptions: {
      scss: {
        additionalData: '@import "@/styles/global.scss";'
      }
    }
  },
  devServer: {
    proxy: {
      '/api': {
        target: 'http://backend:1323',
        ws: true,
        changeOrigin: true,
        headers: {
          Connection: 'keep-alive'
        }
      }
    },
    historyApiFallback: false
  }
};
