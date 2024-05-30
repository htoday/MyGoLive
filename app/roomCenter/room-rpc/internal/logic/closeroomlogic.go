package logic

import (
	"context"

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
	// todo: add your logic here and delete this line

	return &pb.CloseRoomResponse{}, nil
}
