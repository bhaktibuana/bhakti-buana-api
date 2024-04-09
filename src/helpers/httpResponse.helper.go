package helpers

import (
	"bhakti-buana-api/src/constants"

	"github.com/gin-gonic/gin"
)

type S_HttpResponse struct {
	Message string      `json:"message"`
	Status  bool        `json:"status"`
	Data    interface{} `json:"data"`
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

	response := S_HttpResponse{
		Message: message,
		Status:  httpStatus >= 200 && httpStatus < 300,
		Data:    data,
	}

	if response.Status {
		context.JSON(httpStatus, response)
	} else {
		context.AbortWithStatusJSON(httpStatus, response)
	}
}
