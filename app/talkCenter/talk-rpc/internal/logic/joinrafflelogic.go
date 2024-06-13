package logic

import (
	"context"

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
	_, ok := Raffles[in.RoomId]
	if !ok {
		return &pb.JoinRaffleResp{
			Status: 400,
		}, nil
	}
	Raffles[in.RoomId].mu.Lock()
	Raffles[in.RoomId].Participants = append(Raffles[in.RoomId].Participants, in.UserName)
	Raffles[in.RoomId].mu.Unlock()
	return &pb.JoinRaffleResp{}, nil
}
