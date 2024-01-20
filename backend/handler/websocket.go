package handler

import (
	"fmt"
	"logger/constants"
	"logger/controller"
	request "logger/interfaces"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	uuid "github.com/google/uuid"
	"github.com/gorilla/websocket"
)

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

var logsUpdatedChannel = make(chan request.ConnQueryKeeper)
var mut = new(sync.Mutex)

func (h *Handler) HandleWebsocket(ctx *gin.Context) {
	writer := ctx.Writer
	req := ctx.Request
	connUUID := uuid.NewString() // UUID of the connection

	// Upgrade to a websocket connection
	conn, err := wsupgrader.Upgrade(writer, req, nil)

	fmt.Println("---- Connected Successfully ----")

	// Add to set of websocket connections
	(*h.ConnManager)[connUUID] = request.ConnQuery{
		Conn:  conn,
		Query: constants.DEFAULT_QUERY,
	}

	if err != nil {
		fmt.Printf("Failed to set websocket upgrade: %+v", err)
		return
	}

	wg := new(sync.WaitGroup)
	wg.Add(1)
	go controller.ClientInterceptor(conn, h.MongikClient, connUUID, wg, h.ConnManager)
	wg.Wait()

	// Close the connection
	conn.Close()

	// Remove the conn entry
	delete(*h.ConnManager, connUUID)

	fmt.Println("---- Disconnected Successfully ----")
}
