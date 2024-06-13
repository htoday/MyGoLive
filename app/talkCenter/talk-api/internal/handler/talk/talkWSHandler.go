package talk

import (
	"encoding/json"
	"fmt"
	"log"
	model "mygo/app/talkCenter/talk-api/talk"
	"net/http"
	"sync"

	"github.com/zeromicro/go-zero/rest/httpx"
	"mygo/app/talkCenter/talk-api/internal/svc"
)

type Request struct {
	Room string `path:"room"`
}

func TalkWSHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var req Request
		httpx.Parse(r, &req)
		roomId := req.Room
		fmt.Println(roomId)
		model.MutexForRoomMutexes.Lock()
		roomMutex, ok := model.RoomMutexes[roomId]
		if ok {
			roomMutex.Lock()
		} else {
			model.RoomMutexes[roomId] = new(sync.Mutex)
			model.RoomMutexes[roomId].Lock()
		}
		model.MutexForRoomMutexes.Unlock()
		room, ok := model.House[roomId]
		var hub *model.Hub
		if ok {
			hub = room
		} else {
			hub = model.NewHub(roomId)
			model.House[roomId] = hub
			go hub.Run()
		}

		serveWs(hub, w, r)
	}
}
func serveWs(hub *model.Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := model.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := &model.Client{Hub: hub, Conn: conn, Send: make(chan []byte, 256)}
	client.Hub.Clients[client] = true
	//client.Hub.Register <- client
	helloMessage := model.Message{
		Name:    "ABC",
		MsgType: 0,
		Content: "进入了房间",
	}
	helloMessageString, err := json.Marshal(helloMessage)
	client.Hub.Broadcast <- helloMessageString

	model.RoomMutexes[hub.RoomId].Unlock()
	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.WritePump()
	go client.ReadPump()

}
