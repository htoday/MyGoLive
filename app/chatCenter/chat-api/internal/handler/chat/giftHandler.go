package chat

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"mygo/app/chatCenter/chat-api/internal/logic/chat"
	"mygo/app/chatCenter/chat-api/internal/svc"
	"mygo/app/chatCenter/chat-api/internal/types"
)

func GiftHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SendGiftReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := chat.NewGiftLogic(r.Context(), svcCtx)
		resp, err := l.Gift(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
