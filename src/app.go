package app

import (
	"bhakti-buana-api/src/database"
	"bhakti-buana-api/src/routers"
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func DBConnection(uri, dbName string) {
	err := database.Connect(uri, dbName)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Database connected successfully ✨")
}

func Middlewares(app *gin.Engine) {
	// Middleware to set the trusted headers (trust proxy)
	app.Use(func(context *gin.Context) {
		context.Request.Header.Set("X-Real-IP", context.GetHeader("X-Real-IP"))
		context.Request.Header.Set("X-Forwarded-For", context.GetHeader("X-Forwarded-For"))
		context.Request.Header.Set("X-Forwarded-Proto", context.GetHeader("X-Forwarded-Proto"))
		context.Next()
	})

	// Middleware to disable Cross-Origin Embedder Policy
	app.Use(func(context *gin.Context) {
		context.Writer.Header().Set("Cross-Origin-Embedder-Policy", "unsafe-none")
		context.Next()
	})

	app.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowCredentials: true,
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	}))
}

func Routes(app *gin.Engine) {
	routers.Index(app)
}

func Serve(app *gin.Engine, port string) {
	app.Run(port)
}
