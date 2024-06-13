package logic

import (
	"context"
	"encoding/json"
	"fmt"

	"mygo/app/roomCenter/room-rpc/internal/svc"
	"mygo/app/roomCenter/room-rpc/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOneRoomViewNumLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOneRoomViewNumLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOneRoomViewNumLogic {
	return &GetOneRoomViewNumLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetOneRoomViewNumLogic) GetOneRoomViewNum(in *pb.GetOneRoomViewNumRequest) (*pb.GetOneRoomViewNumResponse, error) {
	// todo: add your logic here and delete this line
	roomKey := fmt.Sprintf("room:%d", in.RoomId)
	roomData, err := l.svcCtx.RDB.Get(roomKey)
	if err != nil {
		return nil, err
	}

	var room RDBRoom
	err = json.Unmarshal([]byte(roomData), &room)
	if err != nil {
		return nil, err
	}

	return &pb.GetOneRoomViewNumResponse{
		Status:    200,
		ViewerNum: room.viewNum,
	}, nil

}
