package model

type Message struct {
	Name    string `json:"name"`
	MsgType int    `json:"msgType"`
	Content string `json:"content"`
}
type RaffleResult struct {
	Winners   []string `json:"winners"`
	MsgType   int      `json:"msgType"`
	PrizeName string   `json:"prizeName"`
}
