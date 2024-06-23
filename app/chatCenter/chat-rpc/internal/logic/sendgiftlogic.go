package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"mygo/app/chatCenter/chat-api/dao"
	"mygo/app/chatCenter/chat-rpc/chatroomservice"
	"mygo/app/talkCenter/model"
	"mygo/app/userCenter/user-rpc/userservice"
	"strconv"

	"mygo/app/chatCenter/chat-rpc/internal/svc"
	"mygo/app/chatCenter/chat-rpc/pb/pb"

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

var giftList = map[string]int64{
	"flower": 1,
	"heart":  2,
	"cake":   10,
	"plane":  50,
	"rocket": 100,
}

func (l *SendGiftLogic) SendGift(in *pb.SendGiftReq) (*pb.SendGiftResp, error) {
	roomIdStr := strconv.FormatInt(in.RoomId, 10)
	fmt.Println(model.House)
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
		err = dao.NewProducer(dao.Ch, roomIdStr, string(msgString))
		if err != nil {
			fmt.Println(err)
		}
		return &chatroomservice.SendGiftResp{Status: 200}, nil
	} else {
		return &chatroomservice.SendGiftResp{Status: 400}, nil
	}

}
