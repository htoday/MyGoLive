package logic

import (
	"context"
	"encoding/json"
	"fmt"
	model "mygo/app/talkCenter/model"
	"strconv"
	"time"

	"mygo/app/talkCenter/talk-rpc/internal/svc"
	"mygo/app/talkCenter/talk-rpc/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type StartRaffleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewStartRaffleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StartRaffleLogic {
	return &StartRaffleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *StartRaffleLogic) StartRaffle(in *pb.StartRaffleReq) (*pb.StartRaffleResp, error) {
	key := fmt.Sprintf("raffle:%d:participants", in.RoomId)
	l.svcCtx.RDB.SetCtx(l.ctx, key, "youself")

	go func() {
		time.Sleep(time.Duration(in.Duration) * time.Second)
		winners, err := l.svcCtx.RDB.SrandmemberCtx(l.ctx, key, int(in.PrizeNum))
		if err != nil {
			l.svcCtx.RDB.Del(key)
			logx.Error(err)
			return
		}
		var message model.Message
		message.Name = "系统"
		message.MsgType = 3
		winnersMsg, err := json.Marshal(winners)
		if err != nil {
			logx.Error(err)
		}
		message.Content = string(winnersMsg)
		msgJSON, err := json.Marshal(message)
		if err != nil {
			logx.Error(err)
		}
		model.House[strconv.FormatInt(in.RoomId, 10)].Broadcast <- msgJSON
		l.svcCtx.RDB.Del(key)
	}()

	return &pb.StartRaffleResp{
		Status: 200,
	}, nil
}

type RaffleMsg struct {
	winners []string
}
