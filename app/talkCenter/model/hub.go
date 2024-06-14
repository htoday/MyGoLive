package model

import (
	"fmt"
	"github.com/gorilla/websocket"
	"sync"
	"time"
)

var House = make(map[string]*Hub)

var RoomMutexes = make(map[string]*sync.Mutex)
var MutexForRoomMutexes = new(sync.Mutex)

type Client struct {
	Hub *Hub

	// The websocket connection.
	Conn *websocket.Conn

	// Buffered channel of outbound messages.
	Send chan []byte
}

type Hub struct {
	RoomId string
	// Registered Clients.
	Clients map[*Client]bool

	// Inbound messages from the clients.
	Broadcast chan []byte

	// Register requests from the clients.
	Register chan *Client

	// Unregister requests from Clients.
	unregister chan *Client

	MessageQueue [][]byte
}

func NewHub(roomId string) *Hub {
	hub := &Hub{
		RoomId:       roomId,
		Broadcast:    make(chan []byte),
		Register:     make(chan *Client),
		unregister:   make(chan *Client),
		Clients:      make(map[*Client]bool),
		MessageQueue: make([][]byte, 0),
	}
	go func() {
		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()
		for range ticker.C {
			hub.BroadcastQueuedMessages()
		}
	}()
	return hub
}

func (h *Hub) Run() {
	defer func() {
		close(h.Register)
		close(h.unregister)
		close(h.Broadcast)
	}()
	for {
		select {
		//case client := <-h.Register:
		//	h.clients[client] = true
		case client := <-h.unregister:
			RoomMutexes[h.RoomId].Lock()
			if _, ok := h.Clients[client]; ok {
				delete(h.Clients, client)
				close(client.Send)
				if len(h.Clients) == 0 {
					delete(House, h.RoomId)
					fmt.Println(h.RoomId, "房间已经关闭")
					RoomMutexes[h.RoomId].Unlock()
					return
				}
			}
			RoomMutexes[h.RoomId].Unlock()
		case message := <-h.Broadcast:
			h.MessageQueue = append(h.MessageQueue, message)
			//for client := range h.Clients {
			//	select {
			//	case client.Send <- message:
			//	default:
			//		close(client.Send)
			//		delete(h.Clients, client)
			//	}
			//}
		}
	}
}

func (h *Hub) BroadcastQueuedMessages() {
	for _, message := range h.MessageQueue {
		for client := range h.Clients {
			select {
			case client.Send <- message:
			default:
				close(client.Send)
				delete(h.Clients, client)
			}
		}
	}
	// 清空消息队列
	h.MessageQueue = nil
}
func GetHouse() map[string]*Hub {
	return House
}
