package logic

import (
	"context"
	"fmt"

	"mygo/app/chatCenter/chat-rpc/internal/svc"
	"mygo/app/chatCenter/chat-rpc/pb/pb"

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
	// todo: add your logic here and delete this line
	key := fmt.Sprintf("raffle:%d:participants", in.RoomId)
	exists, err := l.svcCtx.RDB.Exists(key)
	if err != nil {
		return &pb.JoinRaffleResp{Status: 400}, err
	}
	if exists == false {
		return &pb.JoinRaffleResp{Status: 400}, nil
	}
	l.svcCtx.RDB.Sadd(key, in.UserName)
	return &pb.JoinRaffleResp{Status: 200}, nil

}
