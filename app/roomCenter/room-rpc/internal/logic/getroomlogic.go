package logic

import (
	"context"
	"encoding/json"
	"mygo/app/roomCenter/room-rpc/roomservice"

	"mygo/app/roomCenter/room-rpc/internal/svc"
	"mygo/app/roomCenter/room-rpc/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
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
	// todo: add your logic here and delete this line
	var cursor uint64
	var rooms []*roomservice.Room
	var roomsJSON []string
	page := in.Page
	pageSize := int64(10)
	for i := int64(0); i < page; i++ {
		var keys []string
		var err error

		keys, cursor, err = l.svcCtx.RDB.Scan(cursor, "room:*", pageSize)
		if err != nil {
			return nil, err
		}

		if i == page-1 {
			roomsJSON = keys
		}
	}
	for _, roomJSON := range roomsJSON {
		var room roomservice.Room
		err := json.Unmarshal([]byte(roomJSON), &room)
		if err != nil {
			return nil, err
		}
		rooms = append(rooms, &room)
	}
	return &pb.GetRoomResponse{
		Rooms: rooms,
	}, nil
}
