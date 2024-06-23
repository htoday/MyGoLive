package chat

import (
	"context"
	"mygo/app/chatCenter/chat-rpc/chatroomservice"

	"mygo/app/chatCenter/chat-api/internal/svc"
	"mygo/app/chatCenter/chat-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GiftLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGiftLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GiftLogic {
	return &GiftLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GiftLogic) Gift(req *types.SendGiftReq) (resp *types.SendGiftResp, err error) {
	GitfResp, err := l.svcCtx.ChatRpcClient.SendGift(l.ctx, &chatroomservice.SendGiftReq{
		RoomId:   req.RoomId,
		Name:     req.Name,
		GiftType: req.GiftType,
	})
	if err != nil {
		logx.Error(err)
	}

	return &types.SendGiftResp{
		Status: GitfResp.Status,
	}, nil
}
