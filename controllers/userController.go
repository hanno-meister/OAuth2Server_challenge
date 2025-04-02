package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
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

func GetToken(c *gin.Context) {

	// Get emai and PW from request body
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

	// Check if user is authorized by looking up email and password
	// SELECT * FROM users WHERE email = body.Email
	var user models.User
	initializers.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalide email",
		})

		return
	}

	if user.Password != body.Password {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid password",
		})
		return
	}

	// Generate JWT token
	token := jwt.NewWithClaims(
		jwt.SigningMethodRS256,
		jwt.MapClaims{
			"sub": user.ID,
			"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
		})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(initializers.PrivateKey)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create token",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}
