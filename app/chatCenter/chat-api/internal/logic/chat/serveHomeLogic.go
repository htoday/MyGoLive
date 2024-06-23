package chat

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"mygo/app/chatCenter/chat-api/internal/svc"
)

type ServeHomeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewServeHomeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ServeHomeLogic {
	return &ServeHomeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ServeHomeLogic) ServeHome() error {
	// todo: add your logic here and delete this line

	return nil
}
