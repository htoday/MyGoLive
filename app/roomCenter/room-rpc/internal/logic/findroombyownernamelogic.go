package logic

import (
	"context"
	"fmt"

	"mygo/app/roomCenter/room-rpc/internal/svc"
	"mygo/app/roomCenter/room-rpc/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindRoomByOwnerNameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindRoomByOwnerNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindRoomByOwnerNameLogic {
	return &FindRoomByOwnerNameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindRoomByOwnerNameLogic) FindRoomByOwnerName(in *pb.FindRoomByOwnerNameReq) (*pb.FindRoomByOwnerNameResponse, error) {
	// todo: add your logic here and delete this line
	room, err := l.svcCtx.DB.QueryRoomByUsername(l.ctx, in.Username)
	fmt.Println(err)
	fmt.Println(room)
	if err != nil {
		return &pb.FindRoomByOwnerNameResponse{
			Status: 404,
			RoomId: -1,
		}, nil
	}
	return &pb.FindRoomByOwnerNameResponse{
		Status:   200,
		RoomId:   room.RoomId,
		RoomName: room.RoomName.String,
	}, nil
}
