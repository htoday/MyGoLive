package talk

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"mygo/app/talkCenter/talk-api/internal/svc"
)

type TalkWSLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTalkWSLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TalkWSLogic {
	return &TalkWSLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TalkWSLogic) TalkWS() error {
	// todo: add your logic here and delete this line

	return nil
}
