package logic

import (
	"context"
	"fmt"

	"mygo/app/talkCenter/talk-rpc/internal/svc"
	"mygo/app/talkCenter/talk-rpc/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type JoinRaffleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewJoinRaffleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *JoinRaffleLogic {
	return &JoinRaffleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *JoinRaffleLogic) JoinRaffle(in *pb.JoinRaffleReq) (*pb.JoinRaffleResp, error) {
	key := fmt.Sprintf("raffle:%d:participants", in.RoomId)
	exists, err := l.svcCtx.RDB.Exists(key)
	if err != nil {
		return &pb.JoinRaffleResp{Status: 400}, err
	}
	if exists == false {
		return &pb.JoinRaffleResp{Status: 400}, nil
	}
	l.svcCtx.RDB.Set(key, in.UserName)
	return &pb.JoinRaffleResp{Status: 200}, nil
}
