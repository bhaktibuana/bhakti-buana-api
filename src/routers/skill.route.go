package routers

import (
	skillController "bhakti-buana-api/src/controllers/skill"
	routersTemplate "bhakti-buana-api/src/routers/templates"

	"github.com/gin-gonic/gin"
)

func Skill(router *gin.RouterGroup) {
	basePath := "/skill"
	privateRoute := routersTemplate.NewPrivateRoute(basePath, router)
	publicRoute := routersTemplate.NewPublicRoute(basePath, router)

	privateRoute.POST("/create", skillController.Store)
	privateRoute.GET("/", skillController.Index)
	privateRoute.PUT("/:id/update", skillController.Update)
	publicRoute.GET("/show-public", skillController.ShowPublic)
}
