package talk

import (
	"context"
	"encoding/json"
	"fmt"
	"mygo/app/talkCenter/model"
	"mygo/app/talkCenter/talk-api/internal/svc"
	"mygo/app/talkCenter/talk-api/internal/types"
	"mygo/app/userCenter/user-rpc/userservice"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
)

type GiftLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGiftLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GiftLogic {
	return &GiftLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

var giftList = map[string]int64{
	"flower": 1,
	"heart":  2,
	"cake":   10,
	"plane":  50,
	"rocket": 100,
}

func (l *GiftLogic) Gift(req *types.SendGiftReq) (resp *types.SendGiftResp, err error) {

	roomIdStr := strconv.FormatInt(req.RoomId, 10)
	thisHub := model.House[roomIdStr]
	fmt.Println(model.House)
	if thisHub == nil {
		fmt.Println("房间不存在")

		return &types.SendGiftResp{Status: 400}, nil
	}
	spendResp, err := l.svcCtx.UserRpcClient.Spend(l.ctx, &userservice.SpendReq{
		Username: req.Name,
		Num:      giftList[req.GiftType],
	})
	if err != nil {
		fmt.Println(err)
	}
	if spendResp.Status == 200 {
		msg := model.Message{
			Name:    req.Name,
			MsgType: 1,
			Content: req.GiftType,
		}
		msgString, err := json.Marshal(msg)
		if err != nil {
			fmt.Println(err)
		}
		thisHub.Broadcast <- msgString
		return &types.SendGiftResp{Status: 200}, nil
	} else {
		return &types.SendGiftResp{Status: 400}, nil
	}
	//fmt.Println(model.House)
	//GiftResp, err := l.svcCtx.TalkRpcClient.SendGift(l.ctx, &talkroomservice.SendGiftReq{
	//	Name:     req.Name,
	//	RoomId:   req.RoomId,
	//	GiftType: req.GiftType,
	//})
	//if err != nil {
	//	return &types.SendGiftResp{
	//		Status: GiftResp.Status,
	//	}, err
	//}
	//return &types.SendGiftResp{
	//	Status: GiftResp.Status,
	//}, nil

}
