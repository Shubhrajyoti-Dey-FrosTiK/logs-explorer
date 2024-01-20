package poller

import (
	"logger/log"
	"os"

	"github.com/adjust/rmq/v5"
)

func InitPublisher() *rmq.Connection {
	errChan := make(chan error, 10)
	go log.LogErrors(errChan)

	connection, err := rmq.OpenConnection("consumer", "tcp", os.Getenv("REDIS_HOST"), 1, errChan)
	if err != nil {
		panic(err)
	}

	return &connection
}
