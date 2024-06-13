package talk

import (
	"context"
	"mygo/app/talkCenter/talk-rpc/talkroomservice"

	"mygo/app/talkCenter/talk-api/internal/svc"
	"mygo/app/talkCenter/talk-api/internal/types"

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
	JoinRaffleResp, err := l.svcCtx.TalkRpcClient.JoinRaffle(l.ctx, &talkroomservice.JoinRaffleReq{
		RoomId:   req.RoomId,
		UserName: req.UserName,
	})
	if err != nil {
		return &types.JoinRaffleResp{
			Status: JoinRaffleResp.Status,
		}, err
	}
	return &types.JoinRaffleResp{
		Status: JoinRaffleResp.Status,
	}, nil

	return
}
