package controller

import (
	"logger/constants"
	request "logger/interfaces"

	mongik "github.com/FrosTiK-SD/mongik/db"
	models "github.com/FrosTiK-SD/mongik/models"
	"go.mongodb.org/mongo-driver/bson"
)

func GetQueryFromConn(currentConnections request.ConnQueryKeeper) bson.M {
	queryBuilder := make(map[string][]bson.M)

	for connUUID, conn := range currentConnections {
		skip := conn.Query.PageNumber * conn.Query.PageSize
		var querySquence []bson.M = []bson.M{
			{
				"$match": conn.Query.ExtractBSONQuery(true),
			},
			{
				"$sort": bson.M{
					"createdAt": -1,
				},
			},
			{"$limit": skip + conn.Query.PageSize},
			{"$skip": skip},
		}
		queryBuilder[connUUID] = querySquence
	}

	return bson.M{
		"$facet": queryBuilder,
	}
}

func GetBatchAggregates(mongikClient *models.Mongik, pipeline bson.A) (map[string]interface{}, error) {
	return mongik.AggregateOne[map[string]interface{}](mongikClient, constants.DB, constants.COLLECTION_LOGS, pipeline, true)
}
