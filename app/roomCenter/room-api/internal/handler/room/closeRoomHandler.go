package room

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"mygo/app/roomCenter/room-api/internal/logic/room"
	"mygo/app/roomCenter/room-api/internal/svc"
	"mygo/app/roomCenter/room-api/internal/types"
)

func CloseRoomHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CloseRoomReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := room.NewCloseRoomLogic(r.Context(), svcCtx)
		resp, err := l.CloseRoom(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
