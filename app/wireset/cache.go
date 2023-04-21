package wireset

import (
	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
	"log"
	"github.com/one-meta/meta/app/entity/config"
)

// CacheProvider 其它需要注入的也可以写在后面
var CacheProvider = wire.NewSet(NewRedis)

// NewRedis 创建redis客户端
func NewRedis() (*redis.Client, func(), error) {
	dsn := config.CFG.Cache.Redis.DSN()
	// 有两种创建client的方式
	// See: https://redis.uptrace.dev/guide/go-redis.html#connecting-to-redis-server
	opt, err := redis.ParseURL(dsn)
	if err != nil {
		log.Fatal(err)
	}
	rdb := redis.NewClient(opt)
	return rdb, func() {
		err := rdb.Close()
		if err != nil {
			return
		}
	}, err
}
