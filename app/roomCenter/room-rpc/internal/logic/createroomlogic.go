package logic

import (
	"context"
	"encoding/json"
	"math/rand"
	"strconv"
	"time"

	"mygo/app/roomCenter/room-rpc/internal/svc"
	"mygo/app/roomCenter/room-rpc/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateRoomLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateRoomLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateRoomLogic {
	return &CreateRoomLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateRoomLogic) CreateRoom(in *pb.CreateRoomRequest) (*pb.CreateRoomResponse, error) {
	// todo: add your logic here and delete this line
	rand.Seed(time.Now().UnixNano())
	var roomID int64
	for {
		roomID := rand.Intn(90000000) + 10000000
		exists, err := l.svcCtx.RDB.Exists("room:" + strconv.Itoa(roomID))
		if err != nil {
			return nil, err
		}
		if exists == false {
			break
		}
	}
	//生成一个八位数的房间号
	room := pb.Room{
		RoomId:    roomID,
		RoomName:  in.RoomName,
		RoomOwner: in.Username,
		ViewerNum: 0,
	}
	roomJSON, err := json.Marshal(room)
	err = l.svcCtx.RDB.Set("room:"+string(roomID), string(roomJSON))
	if err != nil {
		return &pb.CreateRoomResponse{
			RoomId: roomID,
			Status: 400,
		}, err
	}
	return &pb.CreateRoomResponse{
		RoomId: roomID,
		Status: 200,
	}, nil
}
