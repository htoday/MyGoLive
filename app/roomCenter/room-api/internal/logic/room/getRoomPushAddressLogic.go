package room

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"mygo/app/roomCenter/room-rpc/pb/pb"
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
	roomResp, err := l.svcCtx.RoomRpcClient.GetOneRoom(l.ctx, &pb.GetOneRoomRequest{
		RoomId: req.RoomId,
	})
	if err != nil {
		resp.Status = 400
		return resp, err
	}
	if req.Username != roomResp.Room.RoomOwner {
		resp.Status = 400
		return resp, nil
	}

	addr := "http://localhost:8090/control/get?room=" + strconv.FormatInt(req.RoomId, 10)
	respJSON, err := http.Get(addr)
	if err != nil {
		log.Fatal(err)
	}
	defer respJSON.Body.Close()

	body, err := ioutil.ReadAll(respJSON.Body)
	if err != nil {
		log.Fatal(err)
	}

	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		resp.Status = 400
	} else {
		resp.Status = 200
		resp.PushAdress = "rtmp://localhost:1935/live"
		resp.ChannleKey = response.Data
	}

	log.Println(response.Data)

	return resp, nil
}

type Response struct {
	Data string `json:"data"`
}
