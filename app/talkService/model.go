package main

type Message struct {
	Name    string `json:"name"`
	MsgType int    `json:"msgType"`
	Content string `json:"content"`
}
