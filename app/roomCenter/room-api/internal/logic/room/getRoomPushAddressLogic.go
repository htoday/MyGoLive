package room

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"mygo/app/roomCenter/room-rpc/roomservice"
	"net/http"
	"strconv"

	"mygo/app/roomCenter/room-api/internal/svc"
	"mygo/app/roomCenter/room-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRoomPushAddressLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetRoomPushAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRoomPushAddressLogic {
	return &GetRoomPushAddressLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRoomPushAddressLogic) GetRoomPushAddress(req *types.GetRoomPushAddressReq) (resp *types.GetRoomPushAddressResp, err error) {
	CheckRoomResp, err := l.svcCtx.RoomRpcClient.FindRoomByOwnerName(l.ctx, &roomservice.FindRoomByOwnerNameReq{
		Username: req.Username,
	})
	if CheckRoomResp.Status != 200 {
		return &types.GetRoomPushAddressResp{
			Status: 200,
		}, nil
	}
	addr := "http://localhost:8090/control/get?room=" + strconv.FormatInt(CheckRoomResp.RoomId, 10)
	respJSON, err := http.Get(addr)
	if err != nil {
		log.Println(err)
	}
	defer respJSON.Body.Close()

	body, err := ioutil.ReadAll(respJSON.Body)
	if err != nil {
		log.Println(err)
	}

	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		resp.Status = 400
	} else {
		resp.PushAddress = "rtmp://localhost:1935/live/"
		resp.ChannelKey = response.Data
		resp.Status = 200
		resp.RoomId = CheckRoomResp.RoomId
		resp.RoomName = CheckRoomResp.RoomName
	}

	log.Println(response.Data)

	return resp, nil
}

type Response struct {
	Data string `json:"data"`
}
