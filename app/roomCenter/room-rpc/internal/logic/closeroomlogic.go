package logic

import (
	"context"
	"strconv"

	"mygo/app/roomCenter/room-rpc/internal/svc"
	"mygo/app/roomCenter/room-rpc/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CloseRoomLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCloseRoomLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CloseRoomLogic {
	return &CloseRoomLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CloseRoomLogic) CloseRoom(in *pb.CloseRoomRequest) (*pb.CloseRoomResponse, error) {

	err := l.svcCtx.DB.Delete(l.ctx, in.RoomId)
	_, err = l.svcCtx.RDB.Del("room:" + strconv.FormatInt(in.RoomId, 10))
	if err != nil {
		return &pb.CloseRoomResponse{
			Status: 400,
		}, nil
	}
	return &pb.CloseRoomResponse{
		Status: 200,
	}, nil
}
