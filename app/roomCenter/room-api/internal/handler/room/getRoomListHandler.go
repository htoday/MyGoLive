package room

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"mygo/app/roomCenter/room-api/internal/logic/room"
	"mygo/app/roomCenter/room-api/internal/svc"
	"mygo/app/roomCenter/room-api/internal/types"
)

func GetRoomListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetRoomListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := room.NewGetRoomListLogic(r.Context(), svcCtx)
		resp, err := l.GetRoomList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
