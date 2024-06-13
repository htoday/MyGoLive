package logic

import (
	"context"
	"encoding/json"
	"fmt"
	model "mygo/app/talkCenter/talk-api/talk"
	"mygo/app/talkCenter/talk-rpc/talkroomservice"
	"mygo/app/userCenter/user-rpc/userservice"
	"strconv"

	"mygo/app/talkCenter/talk-rpc/internal/svc"
	"mygo/app/talkCenter/talk-rpc/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendGiftLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendGiftLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendGiftLogic {
	return &SendGiftLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SendGiftLogic) SendGift(in *pb.SendGiftReq) (*pb.SendGiftResp, error) {
	// todo: add your logic here and delete this line
	thisHub := model.House[strconv.FormatInt(in.RoomId, 10)]
	if thisHub == nil {
		fmt.Println("房间不存在")
		return &talkroomservice.SendGiftResp{Status: 400}, nil
	}
	spendResp, err := l.svcCtx.UserRpcClient.Spend(l.ctx, &userservice.SpendReq{
		Username: in.Name,
		Num:      giftList[in.GiftType],
	})
	if err != nil {
		fmt.Println(err)
	}
	if spendResp.Status == 200 {
		msg := model.Message{
			Name:    in.Name,
			MsgType: 1,
			Content: in.GiftType,
		}
		msgString, err := json.Marshal(msg)
		if err != nil {
			fmt.Println(err)
		}
		thisHub.Broadcast <- msgString
		return &talkroomservice.SendGiftResp{Status: 200}, nil
	} else {
		return &talkroomservice.SendGiftResp{Status: 400}, nil
	}

}

var giftList = map[string]int64{
	"flower": 1,
	"heart":  2,
	"cake":   10,
	"plane":  50,
	"rocket": 100,
}
