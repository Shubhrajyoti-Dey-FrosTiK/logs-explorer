package handler

import (
	models "github.com/FrosTiK-SD/mongik/models"
	"github.com/adjust/rmq/v5"
	"github.com/redis/go-redis/v9"
)

type Handler struct {
	MongikClient *models.Mongik
	RQClient     *rmq.Connection
	LogQueue     *rmq.Queue
	RedisClient  *redis.Client
}
