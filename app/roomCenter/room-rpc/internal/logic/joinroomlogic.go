package logic

import (
	"context"
	"strconv"

	"mygo/app/roomCenter/room-rpc/internal/svc"
	"mygo/app/roomCenter/room-rpc/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type JoinRoomLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewJoinRoomLogic(ctx context.Context, svcCtx *svc.ServiceContext) *JoinRoomLogic {
	return &JoinRoomLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *JoinRoomLogic) JoinRoom(in *pb.JoinRoomRequest) (*pb.JoinRoomResponse, error) {
	// todo: add your logic here and delete this line
	exists, err := l.svcCtx.RDB.Exists("room:" + strconv.FormatInt(in.RoomId, 10))
	if err != nil {
		return nil, err
	}
	if exists == false {
		return &pb.JoinRoomResponse{Status: 400}, nil
	}
	return &pb.JoinRoomResponse{Status: 200}, nil
}
