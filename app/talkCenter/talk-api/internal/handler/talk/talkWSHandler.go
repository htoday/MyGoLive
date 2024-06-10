package talk

import (
	"log"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"mygo/app/talkCenter/talk-api/internal/svc"
)

var House = make(map[string]*Hub)

func TalkWSHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		vars := make(map[string]string)
		httpx.ParsePath(r, &vars)
		roomId := vars["room"]
		room, ok := hub.House[roomId]
		var hub *Hub
		if ok {
			hub = room
		} else {
			hub = NewHub()
			House[roomId] = hub
			go hub.run()
		}

		serveWs(hub, w, r)
	}
}
func serveWs(hub *Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := &Client{hub: hub, conn: conn, send: make(chan []byte, 256)}
	client.hub.register <- client
	helloMessage := []byte("ABC " + "进入了房间")
	client.hub.Broadcast <- helloMessage
	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.writePump()
	go client.readPump()

}
