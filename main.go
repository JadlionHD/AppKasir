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
	resources.GetAssets()
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	api := router.Group("/api")
	api.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	api.POST("/login", controllers.PostLoginController)
	router.GET("/validate", middleware.ValidateMiddleware, controllers.Validate)

	router.Use(static.Serve("/", resources.StaticAssets))
	router.SetHTMLTemplate(resources.Tmpl)

	router.NoRoute(func(c *gin.Context) { // fallback
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	log.Printf("Running on http://localhost:80")
	router.Run("localhost:80")

}
