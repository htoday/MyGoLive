package chat

import (
	"context"
	"mygo/app/chatCenter/chat-rpc/chatroomservice"

	"mygo/app/chatCenter/chat-api/internal/svc"
	"mygo/app/chatCenter/chat-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type StartRaffleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStartRaffleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StartRaffleLogic {
	return &StartRaffleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StartRaffleLogic) StartRaffle(req *types.StartRaffleReq) (resp *types.StartRaffleResp, err error) {
	StartRaffleResp, err := l.svcCtx.ChatRpcClient.StartRaffle(l.ctx, &chatroomservice.StartRaffleReq{
		RoomId:    req.RoomId,
		PrizeName: req.PrizeName,
		PrizeNum:  req.PrizeNum,
		Duration:  req.Duration,
	})
	if err != nil {
		logx.Error(err)
	}

	return &types.StartRaffleResp{
		Status: StartRaffleResp.Status,
	}, nil

	return
}
