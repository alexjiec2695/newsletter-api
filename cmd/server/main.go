package main

import (
	_ "git.mcontigo.com/safeplay/newsletter-api/docs"
	"git.mcontigo.com/safeplay/newsletter-api/pkg/cors"
	"git.mcontigo.com/safeplay/newsletter-api/pkg/newsletter/handler"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"os"
	"strings"
)

// @contact.name                Grupo MContigo
// @title                       Newsletter API
// @version                     1.0
// @description                 Newsletter API
func main() {
	r := gin.Default()

	r.Use(cors.AllowCORS())

	newsletterHandler := handler.Build()

	libGroup := r.Group("/subscriptions")
	libGroup.GET("", newsletterHandler.Get)
	libGroup.POST("", newsletterHandler.Post)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	port := os.Getenv("PORT")
	if strings.EqualFold(port, "") {
		port = "4500"
	}

	err := r.Run(":" + port)
	if err != nil {
		panic(err)
	}
}
