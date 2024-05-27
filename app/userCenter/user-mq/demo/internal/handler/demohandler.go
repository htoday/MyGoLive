package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"mygo/app/userCenter/user-mq/demo/internal/logic"
	"mygo/app/userCenter/user-mq/demo/internal/svc"
	"mygo/app/userCenter/user-mq/demo/internal/types"
)

func DemoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewDemoLogic(r.Context(), svcCtx)
		resp, err := l.Demo(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
