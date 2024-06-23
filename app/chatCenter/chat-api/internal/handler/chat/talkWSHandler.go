package chat

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/streadway/amqp"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"mygo/app/chatCenter/chat-api/dao"
	"net/http"
	"time"

	"mygo/app/chatCenter/chat-api/internal/svc"
)

var Upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

type Request struct {
	Room string `path:"room"`
}

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

func TalkWSHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req Request
		httpx.Parse(r, &req)
		roomId := req.Room
		serveWS(w, r, roomId)
	}
}
func serveWS(w http.ResponseWriter, r *http.Request, roomId string) {
	conn, err := Upgrader.Upgrade(w, r, nil)
	if err != nil {
		logx.Error(err)
		return
	}
	go ReadPump(conn, roomId)
	go WritePump(conn, roomId)
}

func ReadPump(Conn *websocket.Conn, roomId string) {
	defer func() {

	}()
	Conn.SetReadLimit(maxMessageSize)
	Conn.SetReadDeadline(time.Now().Add(pongWait))
	Conn.SetPongHandler(func(string) error { Conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, message, err := Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				logx.Error(err)
			}
			break
		}
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		var msgJSON Message
		err = json.Unmarshal(message, &msgJSON)
		if err != nil {
			fmt.Println("error:", err)
		}
		fmt.Println(msgJSON)
		fmt.Println(msgJSON.Content)
		if msgJSON.MsgType == 0 {
			err = dao.NewProducer(dao.Ch, roomId, string(message))
			if err != nil {
				fmt.Println("ReadPumpError:", err)
			}
		}

	}
}

func WritePump(Conn *websocket.Conn, roomId string) {
	ticker := time.NewTicker(pingPeriod)
	defer ticker.Stop()
	msgs, err := dao.NewConsumer(dao.Ch, roomId)
	if err != nil {
		fmt.Println("WritePumpError:", err)
	}
	buffer := &bytes.Buffer{}
	for {
		select {
		case d, ok := <-msgs:
			Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			buffer.Write(d.Body)
			buffer.Write(newline)
			collectMessages(buffer, msgs)

			if buffer.Len() > 0 {
				w, err := Conn.NextWriter(websocket.TextMessage)
				if err != nil {
					return
				}
				if _, err := w.Write(buffer.Bytes()); err != nil {
					return
				}
				if err := w.Close(); err != nil {
					return
				}
				buffer.Reset()
			}
			//w, err := Conn.NextWriter(websocket.TextMessage)
			//if err != nil {
			//	return
			//}
			//w.Write(d.Body)
			//
			//// Add queued chat messages to the current websocket message.
			//n := len(msgs)
			//for i := 0; i < n; i++ {
			//	w.Write(newline)
			//	dd := <-msgs
			//	w.Write(dd.Body)
			//}
			//
			//if err := w.Close(); err != nil {
			//	return
			//}

		case <-ticker.C:
			Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

type Message struct {
	Name    string `json:"name"`
	MsgType int    `json:"msgType"`
	Content string `json:"content"`
}

func collectMessages(buffer *bytes.Buffer, msgs <-chan amqp.Delivery) {
	timeout := time.After(100 * time.Millisecond)
	for {
		select {
		case d, ok := <-msgs:
			if !ok {
				return
			}
			buffer.Write(d.Body)
			buffer.Write(newline)
		case <-timeout:
			return
		}
	}
}
