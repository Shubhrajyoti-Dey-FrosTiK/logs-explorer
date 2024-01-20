package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Log struct {
	Id         primitive.ObjectID `json:"_id" bson:"_id"`
	Level      string             `json:"level" bson:"level"`
	Message    string             `json:"message" bson:"message"`
	ResourceId string             `json:"resourceId" bson:"resourceId"`
	Timestamp  time.Time          `json:"timestamp" bson:"timestamp"`
	TraceId    string             `json:"traceId" bson:"traceId"`
	SpanId     string             `json:"spanId" bson:"spanId"`
	Commit     string             `json:"commit" bson:"commit"`
	Metadata   map[string]string  `json:"metadata" bson:"metadata"`
	CreatedAt  primitive.DateTime `json:"createdAt" bson:"createdAt"`
}
