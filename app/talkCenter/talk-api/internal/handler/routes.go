// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	talk "mygo/app/talkCenter/talk-api/internal/handler/talk"
	"mygo/app/talkCenter/talk-api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/:room",
				Handler: talk.ServeHomeHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/ws/:room",
				Handler: talk.WsHandler(serverCtx),
			},
		},
	)
}
