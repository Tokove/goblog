module.exports = {
  devServer: {
    port: 5173, // 改成其他端口，避免和后端8080冲突
    open: true, // 启动后自动打开浏览器（可选）
    proxy: {    // 如果需要代理 API 请求，也可以配置
      '/api': {
        target: 'http://localhost:8080',
        changeOrigin: true
      }
    }
  }
}
