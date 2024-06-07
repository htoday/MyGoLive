package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	"mygo/app/roomCenter/room-rpc/internal/config"
	"mygo/app/roomCenter/room-rpc/model"
)

type ServiceContext struct {
	Config      config.Config
	RoomRpcConf zrpc.RpcClientConf
	RDB         *redis.Redis
	DB          model.LiveRoomModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		RDB:         redis.MustNewRedis(c.RedisConf),
		RoomRpcConf: c.RoomRpcConf,
		DB:          model.NewLiveRoomModel(sqlx.NewSqlConn("mysql", c.DB.DataSource), c.CacheConf),
	}
}
