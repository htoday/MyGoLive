package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"mygo/app/chatCenter/chat-api/dao"
	"mygo/app/talkCenter/model"
	"strconv"
	"time"

	"mygo/app/chatCenter/chat-rpc/internal/svc"
	"mygo/app/chatCenter/chat-rpc/pb/pb"

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
	// todo: add your logic here and delete this line
	key := fmt.Sprintf("raffle:%d:participants", in.RoomId)
	l.svcCtx.RDB.Sadd(key, "yourself")
	roomIdStr := strconv.FormatInt(in.RoomId, 10)
	var message model.Message
	message.Name = "系统"
	message.MsgType = 2
	message.Content = "抽奖开始"
	msgJSON, err := json.Marshal(message)
	if err != nil {
		logx.Error(err)
	}
	err = dao.NewProducer(dao.Ch, roomIdStr, string(msgJSON))
	if err != nil {
		logx.Error(err)
	}
	go func() {
		time.Sleep(time.Duration(in.Duration) * time.Second)
		winners, err := l.svcCtx.RDB.Srandmember(key, int(in.PrizeNum))
		if err != nil {
			l.svcCtx.RDB.Del(key)
			logx.Error(err)
			fmt.Println(err)
			return
		}
		if winners == nil {
			var message model.Message
			message.Name = "系统"
			message.MsgType = 3
			winnersMsg, err := json.Marshal("没有人中将哦")
			if err != nil {
				logx.Error(err)
			}
			message.Content = string(winnersMsg)
			msgJSON, err := json.Marshal(message)
			if err != nil {
				logx.Error(err)
			}
			err = dao.NewProducer(dao.Ch, roomIdStr, string(msgJSON))
			if err != nil {
				logx.Error(err)
			}
			l.svcCtx.RDB.Del(key)
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
		err = dao.NewProducer(dao.Ch, roomIdStr, string(msgJSON))
		if err != nil {
			logx.Error(err)
		}
		l.svcCtx.RDB.Del(key)
	}()

	return &pb.StartRaffleResp{
		Status: 200,
	}, nil
}
