package routers

import (
	resumeController "bhakti-buana-api/src/controllers/resume"
	routersTemplate "bhakti-buana-api/src/routers/templates"

	"github.com/gin-gonic/gin"
)

func Resume(router *gin.RouterGroup) {
	basePath := "/resume"
	privateRoute := routersTemplate.NewPrivateRoute(basePath, router)

	privateRoute.POST("/upload", resumeController.Store)
	privateRoute.GET("/", resumeController.Index)
	privateRoute.PUT("/:id/status", resumeController.UpdateStatus)
}
