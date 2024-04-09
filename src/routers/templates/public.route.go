package routersTemplate

import (
	"github.com/gin-gonic/gin"
)

type S_PublicRoute struct {
	basePath string
	group    *gin.RouterGroup
}

func NewPublicRoute(basePath string, router *gin.RouterGroup) *S_PublicRoute {
	routerGroup := router.Group(basePath)
	return &S_PublicRoute{basePath: basePath, group: routerGroup}
}

func (router *S_PublicRoute) GET(relativePath string, handlers ...gin.HandlerFunc) {
	router.group.GET(relativePath, handlers...)
}

func (router *S_PublicRoute) POST(relativePath string, handlers ...gin.HandlerFunc) {
	router.group.POST(relativePath, handlers...)
}

func (router *S_PublicRoute) PUT(relativePath string, handlers ...gin.HandlerFunc) {
	router.group.PUT(relativePath, handlers...)
}

func (router *S_PublicRoute) DELETE(relativePath string, handlers ...gin.HandlerFunc) {
	router.group.DELETE(relativePath, handlers...)
}
