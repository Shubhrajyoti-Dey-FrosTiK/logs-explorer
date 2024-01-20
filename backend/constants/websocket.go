package constants

import (
	request "logger/interfaces"
	"time"
)

var DEFAULT_QUERY request.SearchRequest = request.SearchRequest{
	Level:                 "",
	LevelRegex:            "",
	Message:               "",
	MessageRegex:          "",
	ResourceId:            "",
	ResourceIdRegex:       "",
	Timestamp:             time.Time{},
	TimestampRegex:        "",
	TraceId:               "",
	TraceIdRegex:          "",
	SpanId:                "",
	SpanIdRegex:           "",
	Commit:                "",
	CommitRegex:           "",
	ParentResourceId:      "",
	ParentResourceIdRegex: "",
	TimeStart:             time.Time{},
	TimeEnd:               time.Time{},
	FullTextSearch:        "",
	PageNumber:            0,
	PageSize:              60,
}
