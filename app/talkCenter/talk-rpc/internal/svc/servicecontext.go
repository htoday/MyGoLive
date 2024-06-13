package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"mygo/app/talkCenter/talk-rpc/internal/config"
	"mygo/app/userCenter/user-rpc/userservice"
)

type ServiceContext struct {
	Config        config.Config
	UserRpcClient userservice.UserService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		UserRpcClient: userservice.NewUserService(zrpc.MustNewClient(c.UserRpcConf)),
	}
}
