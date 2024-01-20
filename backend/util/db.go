package util

import (
	"context"
	"logger/constants"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateIndex(mongoClient *mongo.Client) {
	index := []mongo.IndexModel{
		{
			Keys: bson.D{{"level", "text"}, {"message", "text"}, {"commit", "text"}, {"metadata", "text"}, {"resourceId", "text"}, {"traceId", "text"}, {"spanId", "text"}, {"timestamp", "text"}},
		},
	}

	mongoClient.Database(constants.DB).Collection(constants.COLLECTION_LOGS).Indexes().CreateMany(context.Background(), index)
}
