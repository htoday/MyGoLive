package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"mygo/app/talkCenter/talk-rpc/internal/svc"
	"mygo/app/talkCenter/talk-rpc/pb/pb"
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

	return &pb.CreatRoomResp{
		Status: 200,
	}, nil
}
