package logic

import (
	"context"
	"database/sql"
	"encoding/json"
	"math/rand"
	"mygo/app/roomCenter/room-rpc/internal/svc"
	"mygo/app/roomCenter/room-rpc/model"
	"mygo/app/roomCenter/room-rpc/pb/pb"
	"strconv"
	"time"

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
	rand.Seed(time.Now().UnixNano())
	var roomID int64
	for {
		roomID = int64(rand.Intn(90000000) + 10000000)
		exists, err := l.svcCtx.RDB.Exists("room:" + strconv.FormatInt(roomID, 10))
		if err != nil {
			return nil, err
		}
		if exists == false {
			break
		}
	}
	CacheRoom := RDBRoom{
		roomId:  roomID,
		viewNum: 0,
	}
	roomJSON, err := json.Marshal(CacheRoom)
	if err != nil {
		return nil, err
	}
	err = l.svcCtx.RDB.Set("room:"+strconv.FormatInt(roomID, 10), string(roomJSON))
	if err != nil {
		return nil, err
	}
	room := model.LiveRoom{
		RoomId:    roomID,
		RoomName:  sql.NullString{String: in.RoomName, Valid: true},
		RoomOwner: sql.NullString{String: in.Username, Valid: true},
		ImgPath:   sql.NullString{String: "app/roomCenter/roomImage", Valid: true},
		ImgName:   sql.NullString{String: "default.jpg", Valid: true},
	}
	l.svcCtx.DB.Insert(l.ctx, &room)
	return &pb.CreateRoomResponse{
		RoomId: roomID,
		Status: 200,
	}, nil
}

type RDBRoom struct {
	roomId  int64 `json:"roomId"`
	viewNum int64 `json:"viewNum"`
}

//func (l *CreateRoomLogic) CreateRoom1(in *pb.CreateRoomRequest) (*pb.CreateRoomResponse, error) {
//	// todo: add your logic here and delete this line
//	rand.Seed(time.Now().UnixNano())
//	var roomID int64
//	for {
//		roomID = int64(rand.Intn(90000000) + 10000000)
//		exists, err := l.svcCtx.RDB.Exists("room:" + strconv.FormatInt(roomID, 10))
//		if err != nil {
//			return nil, err
//		}
//		if exists == false {
//			break
//		}
//	}
//	//生成一个八位数的房间号
//	room := pb.Room{
//		RoomId:    roomID,
//		RoomName:  in.RoomName,
//		RoomOwner: in.Username,
//		ViewerNum: 0, //json编码可能会忽略0值，需要在结构体定义时删除omitempty
//	}
//	roomJSON, err := json.MarshalIndent(room, "", " ")
//
//	//_, err = l.svcCtx.RDB.Lpush("rooms", string(roomJSON))
//	err = l.svcCtx.RDB.Set("room:"+strconv.FormatInt(roomID, 10), string(roomJSON))
//	if err != nil {
//		return nil, err
//	}
//	// 将房间的ID和创建时间添加到有序集合中
//	_, err = l.svcCtx.RDB.Zadd("rooms", int64(time.Now().Unix()), strconv.FormatInt(roomID, 10))
//	if err != nil {
//		return &pb.CreateRoomResponse{
//			RoomId: roomID,
//			Status: 400,
//		}, err
//	}
//	return &pb.CreateRoomResponse{
//		RoomId: roomID,
//		Status: 200,
//	}, nil
//}
