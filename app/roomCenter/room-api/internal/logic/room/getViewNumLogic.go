package room

import (
	"context"
	"mygo/app/roomCenter/room-rpc/roomservice"

	"mygo/app/roomCenter/room-api/internal/svc"
	"mygo/app/roomCenter/room-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetViewNumLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetViewNumLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetViewNumLogic {
	return &GetViewNumLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetViewNumLogic) GetViewNum(req *types.GetViewNumReq) (resp *types.GetViewNumResp, err error) {
	// todo: add your logic here and delete this line
	GetViewNumResp, err := l.svcCtx.RoomRpcClient.GetOneRoomViewNum(l.ctx, &roomservice.GetOneRoomViewNumRequest{
		RoomId: req.RoomId,
	})
	if err != nil {
		return &types.GetViewNumResp{
			Status:    400,
			ViewerNum: GetViewNumResp.ViewerNum,
		}, err
	}
	return &types.GetViewNumResp{
		Status:    GetViewNumResp.Status,
		ViewerNum: GetViewNumResp.ViewerNum,
	}, nil
}
