package model

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"os"
	"time"
	"tracer/utils"
)

var (
	Rdb *redis.Client
)

// 初始化链接
func initClient() (err error) {
	Rdb = redis.NewClient(&redis.Options{
		Addr: utils.RdbHost + ":" + utils.RdbPort, // redis的地址和端口
		//Password: utils.RdbPassWord,                   // 默认没有密码
		DB:       utils.RdbName,     // 默认使用的库的序号
		PoolSize: utils.RdbPoolSize, // 连接池大小
	})
	fmt.Println(
		utils.RdbHost+":"+utils.RdbPort,
		utils.RdbPassWord,
		utils.RdbName,
		utils.RdbPoolSize,
	)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err = Rdb.Ping(ctx).Result()
	return err
}

func V8Example() {
	if err := initClient(); err != nil {
		fmt.Println("连接redis数据库失败，请检查参数：", err)
		os.Exit(1)
	}
}
