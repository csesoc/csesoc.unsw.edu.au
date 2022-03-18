module.exports = {
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
