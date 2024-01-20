package request

import "github.com/gorilla/websocket"

type ConnQuery struct {
	Conn  *websocket.Conn
	Query SearchRequest
}

type ConnQueryKeeper = map[string]ConnQuery
