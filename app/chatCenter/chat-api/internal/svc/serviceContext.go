package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"mygo/app/chatCenter/chat-api/internal/config"
	"mygo/app/chatCenter/chat-rpc/chatroomservice"
)

type ServiceContext struct {
	Config        config.Config
	ChatRpcClient chatroomservice.ChatRoomService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		ChatRpcClient: chatroomservice.NewChatRoomService(zrpc.MustNewClient(c.ChatRpcConf)),
	}
}
