package handler

import (
	"context"
	"logger/constants"
	"logger/log"
	"logger/models"
	"time"

	db "github.com/FrosTiK-SD/mongik/db"
	mongik "github.com/FrosTiK-SD/mongik/models"
	"github.com/adjust/rmq/v5"
	jsoniter "github.com/json-iterator/go"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func NewBatchLogConsumer(monigk *mongik.Mongik, rqClient *rmq.Connection, queue *rmq.Queue, redisClient *redis.Client) *Handler {
	return &Handler{
		MongikClient: monigk,
		RQClient:     rqClient,
		LogQueue:     queue,
		RedisClient:  redisClient,
	}
}

func (h *Handler) Consume(batch rmq.Deliveries) {
	payloads := batch.Payloads()
	log.Debugf("-------- LOG CONSUMPTION STARTED | LEN : %d ---------", len(payloads))

	var logsList []models.Log
	for _, payload := range payloads {
		var data models.Log
		err := json.Unmarshal([]byte(payload), &data)
		if err != nil {
			// Add to queue to again try later
			(*h.LogQueue).Publish("payload")
			continue
		}
		data.Id = primitive.NewObjectID()
		data.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
		logsList = append(logsList, data)
	}

	_, err := db.InsertMany[models.Log](h.MongikClient, constants.DB, constants.COLLECTION_LOGS, logsList)
	if err != nil {
		log.Debugf(err.Error())
		batch.Reject()
		log.Debugf("rejected %v", logsList)
	} else {
		// Otherwise acknowledge
		errors := batch.Ack()
		if len(errors) == 0 {
			log.Debugf("acked %q", payloads)
		}
		for i, err := range errors {
			log.Debugf("failed to ack %q: %q", batch[i].Payload(), err)
		}

		// Now publish it to the pubsub channel
		logDataBytes, _ := json.Marshal(logsList)
		intCmd := h.RedisClient.Publish(context.Background(), constants.REDIS_LOG_STATUS, string(logDataBytes))
		log.Debugf(intCmd.String())
	}

	log.Debugf("---- FINISHED CONSUMING LOG BATCH ----")
}
