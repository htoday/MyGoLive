package logic

import (
	"context"
	"encoding/json"
	"math/rand"
	model "mygo/app/talkCenter/talk-api/talk"
	"strconv"
	"sync"
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
	_, ok := Raffles[in.RoomId]
	if ok {
		return &pb.StartRaffleResp{
			Status: 400,
		}, nil
	}
	go func() {

		var r Raffle
		r = Raffle{
			Participants: []string{},
			prizeNum:     in.PrizeNum,
			duration:     in.Duration,
			PrizeName:    in.PrizeName,
			mu:           sync.Mutex{},
		}
		Raffles[in.RoomId] = &r
		time.Sleep(time.Duration(in.Duration) * time.Second)
		if int64(len(r.Participants)) < r.prizeNum {
			delete(Raffles, in.RoomId)
			var message model.Message
			message.MsgType = 3
			message.Content = "Raffle failed, not enough participants"
			msgString, err := json.Marshal(message)
			if err != nil {
				logx.Error(err, "json marshal failed")
			}
			model.House[strconv.FormatInt(in.RoomId, 10)].Broadcast <- msgString
			return
		}
		winners := r.DrawWinners(int(r.prizeNum))
		var message model.RaffleResult
		message.MsgType = 3
		message.Winners = winners
		message.PrizeName = r.PrizeName
		msgString, err := json.Marshal(message)
		if err != nil {
			logx.Error(err, "json marshal failed")
		}
		model.House[strconv.FormatInt(in.RoomId, 10)].Broadcast <- msgString
		delete(Raffles, in.RoomId)
	}()

	return &pb.StartRaffleResp{
		Status: 200,
	}, nil
}

var Raffles = make(map[int64]*Raffle)

type Raffle struct {
	Participants []string
	prizeNum     int64
	duration     int64
	PrizeName    string
	mu           sync.Mutex
}

func (r *Raffle) DrawWinners(n int) []string {
	r.mu.Lock()
	defer r.mu.Unlock()

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(r.Participants), func(i, j int) {
		r.Participants[i], r.Participants[j] = r.Participants[j], r.Participants[i]
	})

	winners := r.Participants[:n]
	r.Participants = r.Participants[n:]

	return winners
}
