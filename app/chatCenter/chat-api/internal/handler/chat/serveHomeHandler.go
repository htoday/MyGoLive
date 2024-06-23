package chat

import (
	"net/http"

	"mygo/app/chatCenter/chat-api/internal/svc"
)

func ServeHomeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "app/talkService/home.html")
	}
}
