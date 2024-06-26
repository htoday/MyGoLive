package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"mygo/app/talkCenter/talk-api/internal/config"
	"mygo/app/talkCenter/talk-rpc/talkroomservice"
	"mygo/app/userCenter/user-rpc/userservice"
)

type ServiceContext struct {
	Config        config.Config
	UserRpcClient userservice.UserService
	TalkRpcClient talkroomservice.TalkRoomService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		UserRpcClient: userservice.NewUserService(zrpc.MustNewClient(c.UserRpcConf)),
		TalkRpcClient: talkroomservice.NewTalkRoomService(zrpc.MustNewClient(c.TalkRpcConf)),
	}
}
