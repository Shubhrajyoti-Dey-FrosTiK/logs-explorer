package controller

import (
	"logger/constants"
	request "logger/interfaces"
	"logger/models"

	db "github.com/FrosTiK-SD/mongik/db"
	mongik "github.com/FrosTiK-SD/mongik/models"
	"github.com/adjust/rmq/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateLogEntry(mongikClient *mongik.Mongik, queue *rmq.Queue, log *request.CreateLogRequest) error {
	// Convert to bytes
	jsonBytes, err := json.Marshal(log)
	if err != nil {
		return err
	}

	// Publish to redis queue
	err = (*queue).Publish(string(jsonBytes))
	if err != nil {
		return err
	}

	// Reset the cache
	db.DBCacheReset(mongikClient.CacheClient, constants.COLLECTION_LOGS)

	return nil
}

func GetLogs(mongikClient *mongik.Mongik, pageNumber int64, pageSize int64, noCache bool) (*[]models.Log, error) {
	var logs []models.Log
	skip := pageNumber * pageSize
	logs, err := db.Find[models.Log](mongikClient, constants.DB, constants.COLLECTION_LOGS, bson.M{}, noCache, &options.FindOptions{
		Sort: bson.M{"timestamp": -1},
	}, &options.FindOptions{
		Skip:  &skip,
		Limit: &pageSize,
	})
	return &logs, err
}

func GetSearchResult(mongikClient *mongik.Mongik, pageNumber int64, pageSize int64, noCache bool, query request.SearchRequest) (*[]models.Log, error) {
	var logs []models.Log
	var queryBSON bson.M = query.ExtractBSONQuery(false)
	skip := pageNumber * pageSize
	limit := (pageNumber * pageSize) + pageSize
	logs, err := db.Find[models.Log](mongikClient, constants.DB, constants.COLLECTION_LOGS, queryBSON, noCache, &options.FindOptions{
		Sort: bson.M{"timestamp": -1},
	}, &options.FindOptions{
		Skip:  &skip,
		Limit: &limit,
	})
	return &logs, err
}
