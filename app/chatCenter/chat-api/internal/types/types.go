// Code generated by goctl. DO NOT EDIT.
package types

type CreatRoomReq struct {
	RoomId int64 `json:"roomId"`
}

type CreatRoomResp struct {
	Status int64 `json:"status"`
}

type JoinRaffleReq struct {
	UserName string `json:"userName"`
	RoomId   int64  `json:"roomId"`
}

type JoinRaffleResp struct {
	Status int64 `json:"status"`
}

type SendGiftReq struct {
	RoomId   int64  `json:"roomId"`
	Name     string `json:"name"`
	GiftType string `json:"giftType"`
}

type SendGiftResp struct {
	Status int64 `json:"status"`
}

type SendRaffleReq struct {
	RoomId int64  `json:"roomId"`
	Name   string `json:"name"` //谁发起的
}

type SendRaffleResp struct {
	Status int64 `json:"status"`
}

type StartRaffleReq struct {
	RoomId    int64  `json:"roomId"`
	PrizeName string `json:"prizeName"`
	PrizeNum  int64  `json:"prizeNum"`
	Duration  int64  `json:"duration"`
}

type StartRaffleResp struct {
	Status int64 `json:"status"`
}