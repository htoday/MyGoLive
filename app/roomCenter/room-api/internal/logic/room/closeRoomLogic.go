package room

import (
	"context"
	"mygo/app/roomCenter/room-rpc/roomservice"

	"mygo/app/roomCenter/room-api/internal/svc"
	"mygo/app/roomCenter/room-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CloseRoomLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCloseRoomLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CloseRoomLogic {
	return &CloseRoomLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CloseRoomLogic) CloseRoom(req *types.CloseRoomReq) (resp *types.CloseRoomResp, err error) {
	// todo: add your logic here and delete this line
	CloseRoomResp, err := l.svcCtx.RoomRpcClient.CloseRoom(l.ctx, &roomservice.CloseRoomRequest{
		RoomId:   req.RoomId,
		Username: req.Username,
	})
	if err != nil {
		return nil, err
	}
	return &types.CloseRoomResp{
		Status: CloseRoomResp.Status,
	}, nil

}
