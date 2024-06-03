package logic

import (
	"context"
	"encoding/json"
	"fmt"

	"mygo/app/roomCenter/room-rpc/internal/svc"
	"mygo/app/roomCenter/room-rpc/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOneRoomLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOneRoomLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOneRoomLogic {
	return &GetOneRoomLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetOneRoomLogic) GetOneRoom(in *pb.GetOneRoomRequest) (*pb.GetOneRoomResponse, error) {
	// todo: add your logic here and delete this line
	roomKey := fmt.Sprintf("room:%d", in.RoomId)
	roomData, err := l.svcCtx.RDB.Get(roomKey)
	if err != nil {
		return nil, err
	}

	var room pb.Room
	err = json.Unmarshal([]byte(roomData), &room)
	if err != nil {
		return nil, err
	}

	return &pb.GetOneRoomResponse{
		Status: 200,
		Room:   &room,
	}, nil

}
