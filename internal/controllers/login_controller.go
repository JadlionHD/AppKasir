package controllers

import (
	"net/http"

	"github.com/JadlionHD/AppKasir/internal/models"
	"github.com/JadlionHD/AppKasir/internal/schemas"
	"github.com/JadlionHD/AppKasir/internal/utils"
	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func PostLoginController(c *gin.Context) {

	var body LoginRequest

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid username or password",
		})
		return
	}

	var user schemas.UserSchema
	models.DB.First(&user, "username = ?", body.Username)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid username or password",
		})
		return
	}

	if !utils.ComparePassword([]byte(user.Password), []byte(body.Password)) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid password",
		})
		return
	}

	token := utils.GenerateToken(user.Username, user.Permissions, user.ID)

	c.SetSameSite(http.SameSiteLaxMode)
	// 86400 seconds = 1 day
	c.SetCookie("Authorization", token, 86400, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{})
}

func Validate(c *gin.Context) {
	user, _ := c.Get("user")

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    user,
	})
}
