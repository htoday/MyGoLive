package room

import (
	"context"
	"mygo/app/roomCenter/room-rpc/roomservice"

	"mygo/app/roomCenter/room-api/internal/svc"
	"mygo/app/roomCenter/room-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type JoinRoomLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewJoinRoomLogic(ctx context.Context, svcCtx *svc.ServiceContext) *JoinRoomLogic {
	return &JoinRoomLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *JoinRoomLogic) JoinRoom(req *types.JoinRoomReq) (resp *types.JoinRoomResp, err error) {
	// todo: add your logic here and delete this line
	JoinRoomResp, err := l.svcCtx.RoomRpcClient.JoinRoom(l.ctx, &roomservice.JoinRoomRequest{
		Username: req.Username,
		RoomId:   req.RoomId,
	})
	if err != nil {
		return nil, err
	}
	return &types.JoinRoomResp{
		Status: JoinRoomResp.Status,
	}, nil

}
