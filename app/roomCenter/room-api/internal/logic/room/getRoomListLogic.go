package room

import (
	"context"
	"mygo/app/roomCenter/room-rpc/roomservice"

	"mygo/app/roomCenter/room-api/internal/svc"
	"mygo/app/roomCenter/room-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRoomListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetRoomListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRoomListLogic {
	return &GetRoomListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRoomListLogic) GetRoomList(req *types.GetRoomListReq) (resp *types.GetRoomListResp, err error) {
	// todo: add your logic here and delete this line
	getRoomResp, err := l.svcCtx.RoomRpcClient.GetRoom(l.ctx, &roomservice.GetRoomRequest{
		Page: req.Page,
	})
	if err != nil {
		return nil, err
	}
	var rooms []types.Room
	for _, room := range getRoomResp.Rooms {
		rooms = append(rooms, types.Room{
			RoomId:    room.RoomId,
			RoomName:  room.RoomName,
			RoomOwner: room.RoomOwner,
			ViewerNum: room.ViewerNum,
		})
	}
	return &types.GetRoomListResp{
		Status:   getRoomResp.Status,
		RoomList: rooms,
	}, nil

}
