package chat

import (
	"context"
	"mygo/app/chatCenter/chat-rpc/chatroomservice"

	"mygo/app/chatCenter/chat-api/internal/svc"
	"mygo/app/chatCenter/chat-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type JoinRaffleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewJoinRaffleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *JoinRaffleLogic {
	return &JoinRaffleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *JoinRaffleLogic) JoinRaffle(req *types.JoinRaffleReq) (resp *types.JoinRaffleResp, err error) {
	JoinRaffleResp, err := l.svcCtx.ChatRpcClient.JoinRaffle(l.ctx, &chatroomservice.JoinRaffleReq{
		RoomId:   req.RoomId,
		UserName: req.UserName,
	})
	if err != nil {
		logx.Error(err)
	}

	return &types.JoinRaffleResp{
		Status: JoinRaffleResp.Status,
	}, nil

}
