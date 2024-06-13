package room

import (
	"context"
	"mygo/app/roomCenter/room-rpc/roomservice"

	"mygo/app/roomCenter/room-api/internal/svc"
	"mygo/app/roomCenter/room-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateRoomLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateRoomLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateRoomLogic {
	return &CreateRoomLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateRoomLogic) CreateRoom(req *types.CreatRoomReq) (resp *types.CreatRoomResp, err error) {
	// todo: add your logic here and delete this line

	creatRoomResp, err := l.svcCtx.RoomRpcClient.CreateRoom(l.ctx, &roomservice.CreateRoomRequest{
		Username: req.Username,
		RoomName: req.RoomName,
	})
	if err != nil {
		return nil, err
	}
	return &types.CreatRoomResp{
		Status: creatRoomResp.Status,
		RoomId: creatRoomResp.RoomId,
	}, nil
}
