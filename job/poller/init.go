package poller

import (
	"fmt"
	"logger/constants"
	"logger/handler"

	"logger/log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/FrosTiK-SD/mongik"
	"github.com/adjust/rmq/v5"
	"github.com/redis/go-redis/v9"
)

const (
	prefetchLimit = 1000
	pollDuration  = 1 * time.Second
	batchSize     = 100
	batchTimeout  = 1 * time.Second
)

func InitPublisher() *rmq.Connection {
	errChan := make(chan error, 10)
	go log.LogErrors(errChan)

	connection, err := rmq.OpenConnection("consumer", "tcp", os.Getenv("REDIS_HOST"), 2, errChan)
	if err != nil {
		panic(err)
	}

	return &connection
}

func InitConsumer() {
	// Start MongikClient
	mongikClient := mongik.NewClient(os.Getenv(constants.CONNECTION_STRING), constants.CACHING_DURATION)

	// Get the publisher to store failed attempts
	rqClient := InitPublisher()

	// Connect agin to redis for pubsub
	redisClient := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_HOST"),
	})

	// Get the queue
	logQueue, err := (*rqClient).OpenQueue(constants.REDIS_LOG_QUEUE)
	if err != nil {
		log.Debugf(err.Error())
	}

	errChan := make(chan error, 10)
	log.Debugf("Connected to REDIS")
	fmt.Println(" ---- STARTING CONSUMER ---- ")
	go log.LogErrors(errChan)

	connection, err := rmq.OpenConnection("consumer", "tcp", os.Getenv("REDIS_HOST"), 1, errChan)
	if err != nil {
		panic(err)
	}

	queue, err := connection.OpenQueue(constants.REDIS_LOG_QUEUE)
	if err != nil {
		panic(err)
	}
	if err := queue.StartConsuming(prefetchLimit, pollDuration); err != nil {
		panic(err)
	}
	if _, err := queue.AddBatchConsumer(constants.REDIS_LOG_QUEUE, batchSize, batchTimeout, handler.NewBatchLogConsumer(mongikClient, rqClient, &logQueue, redisClient)); err != nil {
		panic(err)
	}

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT)
	defer signal.Stop(signals)

	<-signals // wait for signal
	go func() {
		<-signals // hard exit on second signal (in case shutdown gets stuck)
		os.Exit(1)
	}()

	<-connection.StopAllConsuming() // wait for all Consume() calls to finish

}
