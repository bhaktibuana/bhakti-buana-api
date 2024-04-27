package routers

import (
	aboutController "bhakti-buana-api/src/controllers/about"
	routersTemplate "bhakti-buana-api/src/routers/templates"

	"github.com/gin-gonic/gin"
)

func About(router *gin.RouterGroup) {
	basePath := "/about"
	privateRoute := routersTemplate.NewPrivateRoute(basePath, router)

	privateRoute.PUT("/:id", aboutController.Update)
	privateRoute.PUT("/:id/photo", aboutController.UpdatePhoto)
}
