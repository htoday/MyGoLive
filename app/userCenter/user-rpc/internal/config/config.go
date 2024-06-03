package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	//redis.RedisConf
	KqPusherConf struct {
		Brokers []string
		Topic   string
	}
	CacheConf cache.CacheConf
	DB        struct {
		DataSource string
	}
}
