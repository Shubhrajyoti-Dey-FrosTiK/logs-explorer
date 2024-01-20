package handler

import (
	"logger/constants"
	"logger/controller"
	request "logger/interfaces"
	"strconv"

	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func (h *Handler) HandleAddLog(ctx *gin.Context) {
	var log request.CreateLogRequest

	err := ctx.BindJSON(&log)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": constants.ERROR_INCORRENT_BODY,
			"err":     err.Error(),
		})
		return
	}

	if err := controller.CreateLogEntry(h.MongikClient, h.LogQueue, &log); err != nil {
		ctx.JSON(400, gin.H{
			"err": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Successfully added log",
	})
}

func (h *Handler) HandleGetLogs(ctx *gin.Context) {
	pageNumber, _ := strconv.Atoi(ctx.GetHeader("page-number"))
	pageSize, _ := strconv.Atoi(ctx.GetHeader("page-size"))
	noCache := false
	if ctx.GetHeader("cache-control") == constants.NO_CACHE {
		noCache = true
	}

	if pageSize == 0 {
		pageSize = constants.FIND_LIMIT
	}

	logs, err := controller.GetLogs(h.MongikClient, int64(pageNumber), int64(pageSize), noCache)

	ctx.JSON(200, gin.H{
		"logs": logs,
		"err":  err,
	})
}

func (h *Handler) HandleSearch(ctx *gin.Context) {
	// Extract the queery
	var searchQuery request.SearchRequest
	searchQuery.ExtractFromCTX(ctx)

	noCache := false
	if ctx.GetHeader("cache-control") == constants.NO_CACHE {
		noCache = true
	}

	logs, err := controller.GetSearchResult(h.MongikClient, int64(searchQuery.PageNumber), int64(searchQuery.PageSize), noCache, searchQuery)

	ctx.JSON(200, gin.H{
		"logs": logs,
		"err":  err,
		"q":    searchQuery,
	})
}
