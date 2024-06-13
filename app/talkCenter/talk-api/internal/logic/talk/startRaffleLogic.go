package talk

import (
	"context"
	"mygo/app/talkCenter/talk-rpc/talkroomservice"

	"mygo/app/talkCenter/talk-api/internal/svc"
	"mygo/app/talkCenter/talk-api/internal/types"

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
	StartRaffleResp, err := l.svcCtx.TalkRpcClient.StartRaffle(l.ctx, &talkroomservice.StartRaffleReq{
		RoomId: req.RoomId,
	})
	if err != nil {
		return &types.StartRaffleResp{
			Status: StartRaffleResp.Status,
		}, err
	}
	return &types.StartRaffleResp{
		Status: StartRaffleResp.Status,
	}, nil

}
