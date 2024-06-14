package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"mygo/app/talkCenter/talk-rpc/internal/svc"
	"mygo/app/talkCenter/talk-rpc/pb/pb"
)

type SendGiftLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendGiftLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendGiftLogic {
	return &SendGiftLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SendGiftLogic) SendGift(in *pb.SendGiftReq) (*pb.SendGiftResp, error) {

	return &pb.SendGiftResp{}, nil
}
