package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
	"tracer/model"
)

var ctx = context.Background()

type RdbStore struct{}

// Set 插入string数据，并设置有效时常
func (r RdbStore) Set(name, phone, value string) {
	key := name + phone
	err := model.Rdb.SetEX(ctx, key, value, 60*time.Second).Err()
	if err != nil {
		panic(err)
	}
}

// Get 查询string数据
func (r RdbStore) Get(name, phone string) string {
	key := name + phone
	val, err := model.Rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		return ""
	} else if err != nil {
		panic(err)
	} else {
		return val
	}
}

// Ttl 查询剩余时间数据
func (r RdbStore) Ttl(name, phone string) float64 {
	key := name + phone
	val, err := model.Rdb.TTL(ctx, key).Result()
	if err != nil {
		panic(err)
	}
	return val.Seconds()
}
