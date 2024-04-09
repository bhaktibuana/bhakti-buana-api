package routers

import (
	"bhakti-buana-api/src/configs"
	"bhakti-buana-api/src/helpers"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(router *gin.Engine) {
	router.Use(func(context *gin.Context) {
		scheme := context.Request.Header.Get("X-Forwarded-Proto")
		language := context.Request.Header.Get("Language")

		context.Set("language", language)

		if scheme == "" {
			scheme = "http"
		}

		if configs.AppConfig().GIN_MODE == "release" {
			baseUrl := configs.AppConfig().BASE_URL
			context.Set("baseUrl", baseUrl)
		} else {
			baseUrl := fmt.Sprintf("%s://%s", scheme, context.Request.Host)
			context.Set("baseUrl", baseUrl)
		}

		context.Next()
	})

	ApiRouter(router)

	router.NoRoute(func(context *gin.Context) {
		baseUrl, _ := context.Get("baseUrl")
		url := fmt.Sprintf("%s%s", baseUrl, context.Request.URL.Path)
		helpers.HttpResponse("URL not found.", http.StatusNotFound, context, map[string]interface{}{"url": url})
	})

	router.GET("/", func(context *gin.Context) {
		baseUrl, _ := context.Get("baseUrl")
		url := fmt.Sprintf("%s", baseUrl)
		helpers.HttpResponse("Bhakti Buana API", http.StatusOK, context, map[string]interface{}{"url": url})
	})
}
