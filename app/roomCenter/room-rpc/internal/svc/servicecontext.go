package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"mygo/app/roomCenter/room-rpc/internal/config"
)

type ServiceContext struct {
	Config      config.Config
	RoomRpcConf zrpc.RpcClientConf
	RDB         *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		RDB:         redis.MustNewRedis(c.RedisConf),
		RoomRpcConf: c.RoomRpcConf,
	}
}
