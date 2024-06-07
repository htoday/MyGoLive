package logic

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"mygo/app/roomCenter/room-rpc/internal/svc"
	"mygo/app/roomCenter/room-rpc/pb/pb"
)

type GetRoomLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetRoomLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRoomLogic {
	return &GetRoomLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetRoomLogic) GetRoom(in *pb.GetRoomRequest) (*pb.GetRoomResponse, error) {
	pageSize := int64(5)
	resps, err := l.svcCtx.DB.Query(l.ctx, in.Page, pageSize)
	if err != nil && !errors.Is(err, sqlc.ErrNotFound) {
		return nil, err
	}

	var rooms []*pb.Room
	for _, resp := range *resps {
		rooms = append(rooms, &pb.Room{
			RoomId:    resp.RoomId,
			RoomName:  resp.RoomName.String,
			RoomOwner: resp.RoomOwner.String,
		})

	}
	return &pb.GetRoomResponse{
		Status: 200,
		Rooms:  rooms,
	}, nil
}

//func (l *GetRoomLogic) GetRoom(in *pb.GetRoomRequest) (*pb.GetRoomResponse, error) {
//	// todo: add your logic here and delete this line
//
//	var rooms []*roomservice.Room
//	page := in.Page
//	pageSize := int64(5)
//	start := (page - 1) * pageSize
//	end := start + pageSize - 1
//
//	// 将房间的ID和创建时间添加到有序集合中
//	roomIDs, err := l.svcCtx.RDB.Zrevrange("rooms", start, end)
//	if err != nil {
//		return nil, err
//	}
//
//	// Get the room data for each room ID
//	rooms = make([]*pb.Room, len(roomIDs))
//	for i, roomID := range roomIDs {
//		roomData, err := l.svcCtx.RDB.Get("room:" + roomID)
//		if err != nil {
//			return nil, err
//		}
//
//		var room pb.Room
//		err = json.Unmarshal([]byte(roomData), &room)
//		if err != nil {
//			return nil, err
//		}
//		rooms[i] = &room
//	}
//	return &pb.GetRoomResponse{
//		Status: 200,
//		Rooms:  rooms,
//	}, nil
//}
