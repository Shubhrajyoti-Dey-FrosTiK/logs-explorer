package handler

import (
	request "logger/interfaces"

	models "github.com/FrosTiK-SD/mongik/models"
	"github.com/adjust/rmq/v5"
	"github.com/gorilla/websocket"
)

type Handler struct {
	MongikClient     *models.Mongik
	RQClient         *rmq.Connection
	LogQueue         *rmq.Queue
	WebsocketClients []*websocket.Conn
	ConnManager      *request.ConnQueryKeeper
}
