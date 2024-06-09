// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	room "mygo/app/roomCenter/room-api/internal/handler/room"
	"mygo/app/roomCenter/room-api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/closeRoom",
				Handler: room.CloseRoomHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/creatRoom",
				Handler: room.CreatRoomHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/getRoomList",
				Handler: room.GetRoomListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/getRoomPushAdress",
				Handler: room.GetRoomPushAddressHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/getViewNum",
				Handler: room.GetViewNumHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/joinRoom",
				Handler: room.JoinRoomHandler(serverCtx),
			},
		},
		rest.WithPrefix("/v1"),
	)
}
