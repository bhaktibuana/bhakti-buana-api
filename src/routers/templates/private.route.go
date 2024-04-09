package routersTemplate

import (
	"bhakti-buana-api/src/middlewares"

	"github.com/gin-gonic/gin"
)

type S_PrivateRoute struct {
	basePath string
	group    *gin.RouterGroup
}

func NewPrivateRoute(basePath string, router *gin.RouterGroup) *S_PrivateRoute {
	routerGroup := router.Group(basePath)
	return &S_PrivateRoute{basePath: basePath, group: routerGroup}
}

func (router *S_PrivateRoute) GET(relativePath string, handlers ...gin.HandlerFunc) {
	handlers = append([]gin.HandlerFunc{middlewares.IsAuthenticate}, handlers...)
	router.group.GET(relativePath, handlers...)
}

func (router *S_PrivateRoute) POST(relativePath string, handlers ...gin.HandlerFunc) {
	handlers = append([]gin.HandlerFunc{middlewares.IsAuthenticate}, handlers...)
	router.group.POST(relativePath, handlers...)
}

func (router *S_PrivateRoute) PUT(relativePath string, handlers ...gin.HandlerFunc) {
	handlers = append([]gin.HandlerFunc{middlewares.IsAuthenticate}, handlers...)
	router.group.PUT(relativePath, handlers...)
}

func (router *S_PrivateRoute) DELETE(relativePath string, handlers ...gin.HandlerFunc) {
	handlers = append([]gin.HandlerFunc{middlewares.IsAuthenticate}, handlers...)
	router.group.DELETE(relativePath, handlers...)
}
