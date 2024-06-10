package talk

import (
	"net/http"

	"mygo/app/talkCenter/talk-api/internal/svc"
)

func ServeHomeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "app/talkService/home.html")
	}
}
