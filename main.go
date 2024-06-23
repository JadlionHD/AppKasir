package main

import (
	"log"
	"net/http"

	"github.com/JadlionHD/AppKasir/internal/controllers"
	"github.com/JadlionHD/AppKasir/internal/middleware"
	"github.com/JadlionHD/AppKasir/internal/models"
	"github.com/JadlionHD/AppKasir/internal/resources"
	"github.com/gin-gonic/gin"
)

func init() {
	models.Database()

}

func main() {
	assets, _ := resources.Assets()
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.NoRoute(gin.WrapH(http.FileServer(http.FS(assets))))

	router.GET("/api/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.POST("/api/login", controllers.PostLoginController)
	router.GET("/test-validate", middleware.ValidateMiddleware, controllers.Validate)
	log.Printf("Running on http://localhost:80")
	router.Run("localhost:80")

}
