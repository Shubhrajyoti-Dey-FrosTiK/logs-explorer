package controller

import (
	"context"
	"fmt"
	"logger/constants"
	request "logger/interfaces"
	"sync"

	models "github.com/FrosTiK-SD/mongik/models"
	"github.com/gorilla/websocket"
	jsoniter "github.com/json-iterator/go"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/bson"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

// Whenever there is an update send the updated response to the clients
func LogStatusConsumer(ctx context.Context, mongikClient *models.Mongik, redisClient *redis.Client, connManager *request.ConnQueryKeeper) {
	subscriber := redisClient.Subscribe(ctx, constants.REDIS_LOG_STATUS)

	for {
		fmt.Println("Waiting for message")
		_, err := subscriber.ReceiveMessage(ctx)
		if err != nil {
			// Used for breaking out of the loop
			break
		}

		query := GetQueryFromConn(*connManager)
		logsMap, _ := GetBatchAggregates(mongikClient, bson.A{query})

		fmt.Println("CONNECTION_COUNT : ", len(*connManager))
		for connUUID, conn := range *connManager {
			logsBytes, _ := json.Marshal(logsMap[connUUID])
			conn.Conn.WriteMessage(1, logsBytes)
		}
	}
}

func ClientInterceptor(conn *websocket.Conn, mongikClient *models.Mongik, connUUID string, wg *sync.WaitGroup, connManager *request.ConnQueryKeeper) {
	for {
		// Read message from browser
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}

		var search request.SearchRequest
		err = json.Unmarshal(msg, &search)
		if err != nil {
			fmt.Println("Error unmarshalling JSON: ", err)
			continue
		}

		if search.PageSize == 0 {
			search.PageSize = 60
		}

		// Set the updated search query in the channel
		(*connManager)[connUUID] = request.ConnQuery{
			Conn:  conn,
			Query: search,
		}

		// Query the DB
		logs, _ := GetSearchResult(mongikClient, int64(search.PageNumber), int64(search.PageSize), true, search)
		logsBytes, _ := json.Marshal(logs)

		// Send back to the client
		if err = conn.WriteMessage(msgType, logsBytes); err != nil {
			break
		}
	}

	defer wg.Done()
}
