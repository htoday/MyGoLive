package svc

import (
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"mygo/app/userCenter/user-rpc/internal/config"
)

type ServiceContext struct {
	Config         config.Config
	RDB            *redis.Redis
	KqPusherClient *kq.Pusher
}

func NewServiceContext(c config.Config) *ServiceContext {
	RDBConf := redis.RedisConf{
		Host: "localhost:6379",
		Type: "node",
		Pass: "",
	}
	r, err := redis.NewRedis(RDBConf)
	if err != nil {
		panic(err)
	}
	return &ServiceContext{
		Config:         c,
		RDB:            r,
		KqPusherClient: kq.NewPusher(c.KqPusherConf.Brokers, c.KqPusherConf.Topic),
	}
}
