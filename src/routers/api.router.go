package routers

import (
	"github.com/gin-gonic/gin"
)

func ApiRouter(router *gin.Engine) {
	basePath := "/api"
	apiGroup := router.Group(basePath)
	{
		Auth(apiGroup)
	}
}
