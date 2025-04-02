package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hanno-meister/OAuth2Server_challenge/initializers"
	"github.com/hanno-meister/OAuth2Server_challenge/models"
)

func Signup(c *gin.Context) {

	// Get email/pw from req body
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	// Create the user and save in DB
	user := models.User{Email: body.Email, Password: body.Password}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user",
		})

		return
	}

	// Respond
	c.JSON(http.StatusOK, gin.H{
		"message": "Succefully created user",
	})
}
