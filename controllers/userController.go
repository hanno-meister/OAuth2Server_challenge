package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/hanno-meister/OAuth2Server_challenge/initializers"
	"github.com/hanno-meister/OAuth2Server_challenge/models"
	"golang.org/x/crypto/bcrypt"
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

	// Hash password
	hashed_pw, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})

		return
	}

	user := models.User{Email: body.Email, Password: string(hashed_pw)}
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

	// Compare sent in password with hashed password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Password",
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

func ListSigningKeys(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"key": initializers.JwkKey,
	})
}

func IntrospectToken(c *gin.Context) {
	claims, _ := c.Get("claims")

	c.JSON(http.StatusOK, gin.H{
		"active":    true,
		"client_id": claims.(jwt.MapClaims)["sub"],
		"exp":       claims.(jwt.MapClaims)["exp"].(float64),
	})
}
