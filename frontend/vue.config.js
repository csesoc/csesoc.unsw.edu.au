module.exports = {
    pluginOptions: {
        "style-resources-loader": {
            preProcessor: "scss",
            patterns: ['@/styles/global.scss']
        }
    },
    devServer: {
        proxy: {
            '/api': {
                target: "http://backend:1323",
                ws: true,
                changeOrigin: true,
                headers: {
                    Connection: 'keep-alive'
                }
            }
        },
        historyApiFallback: false
    },
} 