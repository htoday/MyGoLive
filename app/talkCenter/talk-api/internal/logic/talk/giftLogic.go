package talk

import (
	"context"
	"mygo/app/talkCenter/talk-api/internal/svc"
	"mygo/app/talkCenter/talk-api/internal/types"
	"mygo/app/talkCenter/talk-rpc/talkroomservice"

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
	GiftResp, err := l.svcCtx.TalkRpcClient.SendGift(l.ctx, &talkroomservice.SendGiftReq{
		Name:     req.Name,
		RoomId:   req.RoomId,
		GiftType: req.GiftType,
	})
	if err != nil {
		return &types.SendGiftResp{
			Status: GiftResp.Status,
		}, err
	}
	return &types.SendGiftResp{
		Status: GiftResp.Status,
	}, nil
}
