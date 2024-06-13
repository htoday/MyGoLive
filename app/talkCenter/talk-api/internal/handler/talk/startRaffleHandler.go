package talk

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"mygo/app/talkCenter/talk-api/internal/logic/talk"
	"mygo/app/talkCenter/talk-api/internal/svc"
	"mygo/app/talkCenter/talk-api/internal/types"
)

func StartRaffleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.StartRaffleReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := talk.NewStartRaffleLogic(r.Context(), svcCtx)
		resp, err := l.StartRaffle(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
