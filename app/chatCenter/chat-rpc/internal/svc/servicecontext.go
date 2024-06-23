package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"mygo/app/chatCenter/chat-rpc/internal/config"
	"mygo/app/userCenter/user-rpc/userservice"
)

type ServiceContext struct {
	Config        config.Config
	UserRpcClient userservice.UserService
	RDB           *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		UserRpcClient: userservice.NewUserService(zrpc.MustNewClient(c.UserRpcConf)),
		RDB:           redis.MustNewRedis(c.RedisConf),
	}
}
