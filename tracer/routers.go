package main

import (
	"tracer/middleware"
	"tracer/utils"

	"github.com/gin-gonic/gin"
)

type Option func(*gin.Engine)

// var options = []Option{}
var options = make([]Option, 0)

// Include 注册app的路由配置
func include(opts ...Option) {
	options = append(options, opts...)
}

// Init 初始化
func Init() *gin.Engine {
	gin.SetMode(utils.AppMode) // 有三种模式，默认debug：debug, release, test
	r := gin.New()             // 创建无中间件的路由
	// r := gin.Default()      // 默认使用了Logger和Recovery中间件
	r.Use(middleware.Cors(), middleware.Logger())
	for _, opt := range options {
		opt(r)
	}
	return r
}