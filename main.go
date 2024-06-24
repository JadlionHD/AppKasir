package main

import (
	"log"
	"net/http"

	"github.com/JadlionHD/AppKasir/internal/controllers"
	"github.com/JadlionHD/AppKasir/internal/middleware"
	"github.com/JadlionHD/AppKasir/internal/models"
	"github.com/JadlionHD/AppKasir/internal/resources"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func init() {
	models.Database()

}

func main() {
	assets := resources.Assets()
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.GET("/api/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.POST("/api/login", controllers.PostLoginController)
	router.GET("/test-validate", middleware.ValidateMiddleware, controllers.Validate)

	router.Use(static.Serve("/", assets))
	router.NoRoute(func(c *gin.Context) {
		c.File("./internal/resources/website/build/index.html")
	})

	log.Printf("Running on http://localhost:80")
	router.Run("localhost:80")

}
