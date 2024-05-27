package user

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"mygo/app/userCenter/user-api/internal/logic/user"
	"mygo/app/userCenter/user-api/internal/svc"
	"mygo/app/userCenter/user-api/internal/types"
)

func SendMessageHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SendMessageReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewSendMessageLogic(r.Context(), svcCtx)
		resp, err := l.SendMessage(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
