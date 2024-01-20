package main

import (
	"context"
	"log"
	"os"

	"logger/constants"
	"logger/controller"
	"logger/handler"
	request "logger/interfaces"
	"logger/poller"
	"logger/util"

	"github.com/FrosTiK-SD/mongik"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func main() {
	r := gin.Default()

	// Start MongikClient
	mongikClient := mongik.NewClient(os.Getenv(constants.CONNECTION_STRING), constants.CACHING_DURATION)

	// Initiate a local redis client
	redisClient := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_HOST"),
	})

	// Create Indexes before hand
	util.CreateIndex(mongikClient.MongoClient)

	// go poller.InitConsumer()
	rqClient := poller.InitPublisher()

	// Get the queue
	logQueue, err := (*rqClient).OpenQueue(constants.REDIS_LOG_QUEUE)
	if err != nil {
		log.Fatal(err)
	}

	r.Use(cors.New(util.DefaultCors()))

	connManager := make(request.ConnQueryKeeper, 0)
	handler := &handler.Handler{
		MongikClient: mongikClient,
		RQClient:     rqClient,
		LogQueue:     &logQueue,
		ConnManager:  &connManager,
	}

	// Subscribe to RedisPubSub and send messeges to connections
	go controller.LogStatusConsumer(context.Background(), mongikClient, redisClient, &connManager)

	r.POST("/", handler.HandleAddLog)
	r.GET("/", handler.HandleGetLogs)
	r.GET("/search", handler.HandleSearch)

	// Websocket:
	r.GET("/ws", handler.HandleWebsocket)

	port := "" + os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	r.Run(":" + port)

	redisClient.Close()
}
