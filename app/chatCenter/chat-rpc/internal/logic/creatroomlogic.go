package logic

import (
	"context"

	"mygo/app/chatCenter/chat-rpc/internal/svc"
	"mygo/app/chatCenter/chat-rpc/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreatRoomLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreatRoomLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreatRoomLogic {
	return &CreatRoomLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreatRoomLogic) CreatRoom(in *pb.CreatRoomReq) (*pb.CreatRoomResp, error) {
	// todo: add your logic here and delete this line

	return &pb.CreatRoomResp{}, nil
}
