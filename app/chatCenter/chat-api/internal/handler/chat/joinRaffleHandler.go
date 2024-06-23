package chat

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"mygo/app/chatCenter/chat-api/internal/logic/chat"
	"mygo/app/chatCenter/chat-api/internal/svc"
	"mygo/app/chatCenter/chat-api/internal/types"
)

func JoinRaffleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.JoinRaffleReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := chat.NewJoinRaffleLogic(r.Context(), svcCtx)
		resp, err := l.JoinRaffle(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
