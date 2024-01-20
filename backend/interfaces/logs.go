package request

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateLogRequest struct {
	Id         primitive.ObjectID `json:"_id" bson:"_id"`
	Level      string             `json:"level" bson:"level"`
	Message    string             `json:"message" bson:"message"`
	ResourceId string             `json:"resourceId" bson:"resourceId"`
	Timestamp  time.Time          `json:"timestamp" bson:"timestamp"`
	TraceId    string             `json:"traceId" bson:"traceId"`
	SpanId     string             `json:"spanId" bson:"spanId"`
	Commit     string             `json:"commit" bson:"commit"`
	Metadata   map[string]string  `json:"metadata" bson:"metadata"`
}

type SearchRequest struct {
	Level                 string    `json:"level" bson:"level"`
	LevelRegex            string    `json:"levelRegex" bson:"levelRegex"`
	Message               string    `json:"message" bson:"message"`
	MessageRegex          string    `json:"messageRegex" bson:"messageRegex"`
	ResourceId            string    `json:"resourceId" bson:"resourceId"`
	ResourceIdRegex       string    `json:"resourceIdRegex" bson:"resourceIdRegex"`
	Timestamp             time.Time `json:"timestamp" bson:"timestamp"`
	TimestampRegex        string    `json:"timestampRegex" bson:"timestampRegex"`
	TraceId               string    `json:"traceId" bson:"traceId"`
	TraceIdRegex          string    `json:"traceIdRegex" bson:"traceIdRegex"`
	SpanId                string    `json:"spanId" bson:"spanId"`
	SpanIdRegex           string    `json:"spanIdRegex" bson:"spanIdRegex"`
	Commit                string    `json:"commit" bson:"commit"`
	CommitRegex           string    `json:"commitRegex" bson:"commitRegex"`
	ParentResourceId      string    `json:"parentResourceId" bson:"parentResourceId"`
	ParentResourceIdRegex string    `json:"parentResourceIdRegex" bson:"parentResourceIdRegex"`
	TimeStart             time.Time `json:"timeStart" bson:"timeStart"`
	TimeEnd               time.Time `json:"timeEnd" bson:"timeEnd"`
	FullTextSearch        string    `json:"fullTextSearch" bson:"fullTextSearch"`
	PageSize              int       `json:"pageSize" bson:"pageSize"`
	PageNumber            int       `json:"pageNumber" bson:"pageNumber"`
}

func (c *SearchRequest) ExtractFromCTX(ctx *gin.Context) {
	c.Level = ctx.Query("level")
	c.LevelRegex = ctx.Query("levelRegex")
	c.Message = ctx.Query("message")
	c.MessageRegex = ctx.Query("messageRegex")
	c.ResourceId = ctx.Query("resourceId")
	c.ResourceIdRegex = ctx.Query("resourceIdRegex")
	c.Timestamp, _ = time.Parse(time.RFC3339, ctx.Query("timestamp"))
	c.TimestampRegex = ctx.Query("timestampRegex")
	c.TraceId = ctx.Query("traceId")
	c.TraceIdRegex = ctx.Query("traceIdRegex")
	c.SpanId = ctx.Query("spanId")
	c.SpanIdRegex = ctx.Query("spanIdRegex")
	c.Commit = ctx.Query("commit")
	c.CommitRegex = ctx.Query("commitRegex")
	c.ParentResourceId = ctx.Query("parentResourceId")
	c.ParentResourceIdRegex = ctx.Query("parentResourceIdRegex")
	c.TimeStart, _ = time.Parse(time.RFC3339, ctx.Query("timeStart"))
	c.TimeEnd, _ = time.Parse(time.RFC3339, ctx.Query("timeEnd"))
	c.FullTextSearch = ctx.Query("fullTextSearch")
	c.PageSize, _ = strconv.Atoi(ctx.Query("pageSize"))
	c.PageNumber, _ = strconv.Atoi(ctx.Query("pageNumber"))

	if c.PageSize == 0 {
		c.PageSize = 60
	}
}

func (c *SearchRequest) ExtractBSONQuery(omitFullTextSearch bool) bson.M {
	query := bson.M{}

	if c.Level != "" {
		query["level"] = c.Level
	}
	if c.LevelRegex != "" {
		query["level"] = bson.M{
			"$regex": c.LevelRegex,
		}
	}
	if c.Message != "" {
		query["message"] = c.Message
	}
	if c.MessageRegex != "" {
		query["message"] = bson.M{
			"$regex": c.MessageRegex,
		}
	}
	if c.ResourceId != "" {
		query["resourceId"] = c.ResourceId
	}
	if c.ResourceIdRegex != "" {
		query["resourceId"] = bson.M{
			"$regex": c.ResourceIdRegex,
		}
	}
	if !c.Timestamp.IsZero() {
		query["timestamp"] = c.Timestamp
	}
	if c.TimestampRegex != "" {
		query["timestamp"] = bson.M{
			"$regex": c.TimestampRegex,
		}
	}
	if c.TraceId != "" {
		query["traceId"] = c.TraceId
	}
	if c.TraceIdRegex != "" {
		query["traceId"] = bson.M{
			"$regex": c.TraceIdRegex,
		}
	}
	if c.SpanId != "" {
		query["spanId"] = c.SpanId
	}
	if c.SpanIdRegex != "" {
		query["spanId"] = bson.M{
			"$regex": c.SpanIdRegex,
		}
	}
	if c.Commit != "" {
		query["commit"] = c.Commit
	}
	if c.CommitRegex != "" {
		query["commit"] = bson.M{
			"$regex": c.CommitRegex,
		}
	}
	if c.ParentResourceId != "" {
		query["parentResourceId"] = c.ParentResourceId
	}
	if c.ParentResourceIdRegex != "" {
		query["parentResourceId"] = bson.M{
			"$regex": c.ParentResourceIdRegex,
		}
	}
	if !c.Timestamp.IsZero() {
		query["timestamp"] = bson.M{
			"$gte": c.Timestamp,
		}
	}
	if c.TimestampRegex != "" {
		query["timestamp"] = bson.M{
			"$regex": c.TimestampRegex,
		}
	}

	if !c.TimeStart.IsZero() && !c.TimeEnd.IsZero() {
		query["timestamp"] = bson.M{
			"$gte": c.TimeStart,
			"$lte": c.TimeEnd,
		}
	} else if !c.TimeStart.IsZero() {
		query["timestamp"] = bson.M{
			"$lte": c.TimeEnd,
		}
	} else if !c.TimeEnd.IsZero() {
		query["timestamp"] = bson.M{
			"$gte": c.TimeStart,
		}
	}
	if !omitFullTextSearch && c.FullTextSearch != "" {
		query["$text"] = bson.M{
			"$search": c.FullTextSearch,
		}
	}

	return query
}
