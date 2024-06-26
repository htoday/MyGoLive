// Code generated by goctl. DO NOT EDIT.
// Source: chat.proto

package chatroomservice

import (
	"context"

	"mygo/app/chatCenter/chat-rpc/pb/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CreatRoomReq    = pb.CreatRoomReq
	CreatRoomResp   = pb.CreatRoomResp
	JoinRaffleReq   = pb.JoinRaffleReq
	JoinRaffleResp  = pb.JoinRaffleResp
	SendGiftReq     = pb.SendGiftReq
	SendGiftResp    = pb.SendGiftResp
	StartRaffleReq  = pb.StartRaffleReq
	StartRaffleResp = pb.StartRaffleResp

	ChatRoomService interface {
		CreatRoom(ctx context.Context, in *CreatRoomReq, opts ...grpc.CallOption) (*CreatRoomResp, error)
		SendGift(ctx context.Context, in *SendGiftReq, opts ...grpc.CallOption) (*SendGiftResp, error)
		StartRaffle(ctx context.Context, in *StartRaffleReq, opts ...grpc.CallOption) (*StartRaffleResp, error)
		JoinRaffle(ctx context.Context, in *JoinRaffleReq, opts ...grpc.CallOption) (*JoinRaffleResp, error)
	}

	defaultChatRoomService struct {
		cli zrpc.Client
	}
)

func NewChatRoomService(cli zrpc.Client) ChatRoomService {
	return &defaultChatRoomService{
		cli: cli,
	}
}

func (m *defaultChatRoomService) CreatRoom(ctx context.Context, in *CreatRoomReq, opts ...grpc.CallOption) (*CreatRoomResp, error) {
	client := pb.NewChatRoomServiceClient(m.cli.Conn())
	return client.CreatRoom(ctx, in, opts...)
}

func (m *defaultChatRoomService) SendGift(ctx context.Context, in *SendGiftReq, opts ...grpc.CallOption) (*SendGiftResp, error) {
	client := pb.NewChatRoomServiceClient(m.cli.Conn())
	return client.SendGift(ctx, in, opts...)
}

func (m *defaultChatRoomService) StartRaffle(ctx context.Context, in *StartRaffleReq, opts ...grpc.CallOption) (*StartRaffleResp, error) {
	client := pb.NewChatRoomServiceClient(m.cli.Conn())
	return client.StartRaffle(ctx, in, opts...)
}

func (m *defaultChatRoomService) JoinRaffle(ctx context.Context, in *JoinRaffleReq, opts ...grpc.CallOption) (*JoinRaffleResp, error) {
	client := pb.NewChatRoomServiceClient(m.cli.Conn())
	return client.JoinRaffle(ctx, in, opts...)
}
