package svc

import (
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"mygo/app/userCenter/user-rpc/internal/config"
	"mygo/app/userCenter/user-rpc/model"
)

type ServiceContext struct {
	Config         config.Config
	RDB            *redis.Redis
	KqPusherClient *kq.Pusher
	DB             model.ZeroUserModel
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
		DB:             model.NewZeroUserModel(sqlx.NewSqlConn("mysql", c.DB.DataSource), c.CacheConf),
		KqPusherClient: kq.NewPusher(c.KqPusherConf.Brokers, c.KqPusherConf.Topic),
	}
}
