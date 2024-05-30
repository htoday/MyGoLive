package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"mygo/app/roomCenter/room-api/internal/config"
	"mygo/app/roomCenter/room-rpc/roomservice"
)

type ServiceContext struct {
	Config        config.Config
	RoomRpcClient roomservice.RoomService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		RoomRpcClient: roomservice.NewRoomService(zrpc.MustNewClient(c.RoomRpcConf)),
	}
}
