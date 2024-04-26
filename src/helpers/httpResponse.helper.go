package helpers

import (
	"bhakti-buana-api/src/constants"

	"github.com/gin-gonic/gin"
)

type S_BaseResponse struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
}

type S_HttpResponse struct {
	S_BaseResponse
	Data interface{} `json:"data"`
}

type S_HttpResponsePaginate struct {
	S_HttpResponse
	Pagination S_Pagination `json:"pagination"`
}

func baseResponse(message string, httpStatus int, context *gin.Context, useLanguage bool) S_BaseResponse {
	if useLanguage {
		selectedLanguage, _ := context.Get("language")
		var messageString string
		if val, ok := selectedLanguage.(string); ok {
			messageString = constants.Languages[val][message]
		}
		if messageString != "" {
			message = messageString
		}
	}

	response := S_BaseResponse{
		Message: message,
		Status:  httpStatus >= 200 && httpStatus < 300,
	}

	return response
}

// HttpResponse Helper
/*
 * @param message string
 * @param httpStatus int
 * @param context *gin.Context
 * @param data interface{}
 * @param options ...interface{} (useLanguage)
 * @returns
 */
func HttpResponse(message string, httpStatus int, context *gin.Context, data interface{}, options ...interface{}) {
	useLanguage := true
	if len(options) > 0 {
		if val, ok := options[0].(bool); ok {
			useLanguage = val
		}
	}

	response := S_HttpResponse{
		S_BaseResponse: baseResponse(message, httpStatus, context, useLanguage),
		Data:           data,
	}

	if response.Status {
		context.JSON(httpStatus, response)
	} else {
		context.AbortWithStatusJSON(httpStatus, response)
	}
}

// HttpResponsePaginate Helper
/*
 * @param message string
 * @param httpStatus int
 * @param context *gin.Context
 * @param data interface{}
 * @param pagination S_Pagination
 * @param options ...interface{} (useLanguage)
 * @returns
 */
func HttpResponsePaginate(message string, httpStatus int, context *gin.Context, data interface{}, pagination S_Pagination, options ...interface{}) {
	useLanguage := true
	if len(options) > 0 {
		if val, ok := options[0].(bool); ok {
			useLanguage = val
		}
	}

	response := S_HttpResponsePaginate{
		S_HttpResponse: S_HttpResponse{
			S_BaseResponse: baseResponse(message, httpStatus, context, useLanguage),
			Data:           data,
		},
		Pagination: pagination,
	}

	if response.Status {
		context.JSON(httpStatus, response)
	} else {
		context.AbortWithStatusJSON(httpStatus, response)
	}
}
