package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/hanno-meister/OAuth2Server_challenge/initializers"
)

func Authentication(c *gin.Context) {
	var body struct {
		Token string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	// Decode and Validate
	token, err := jwt.Parse(body.Token, func(token *jwt.Token) (interface{}, error) {
		return initializers.PublicKey, nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodRS256.Alg()}))

	if err != nil {
		c.AbortWithStatusJSON(401, gin.H{
			"error": err.Error()})
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		//Check the exp date
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		// Attach claims
		c.Set("claims", claims)

		//Continue
		c.Next()

	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	c.Next()
}
