package talk

import (
	"context"

	"mygo/app/talkCenter/talk-api/internal/svc"
	"mygo/app/talkCenter/talk-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreatRoomLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreatRoomLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreatRoomLogic {
	return &CreatRoomLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreatRoomLogic) CreatRoom(req *types.CreatRoomReq) (resp *types.CreatRoomResp, err error) {
	// todo: add your logic here and delete this line

	return
}
