package routers

import (
	authController "bhakti-buana-api/src/controllers/auth"
	routersTemplate "bhakti-buana-api/src/routers/templates"

	"github.com/gin-gonic/gin"
)

func Auth(router *gin.RouterGroup) {
	basePath := "/auth"
	publicRoute := routersTemplate.NewPublicRoute(basePath, router)

	publicRoute.POST("/login", authController.Login)
}
