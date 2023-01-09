package main

import (
	"fmt"
	"tracer/app/blog"
	"tracer/app/user"
	"tracer/model"
	"tracer/utils"
)

func main() {
	// 数据库配置
	model.InitDb()

	// redis配置
	model.V8Example()

	// 加载多个APP的路由配置, 此处目前只加载了一个
	include(blog.Routers, user.Routers)

	// 初始化路由
	r := Init()
	if err := r.Run(utils.HttpPort); err != nil {
		fmt.Printf("Startup service failed, err:%v", err)
	}
}