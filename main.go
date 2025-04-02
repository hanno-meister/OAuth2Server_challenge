package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hanno-meister/OAuth2Server_challenge/controllers"
	"github.com/hanno-meister/OAuth2Server_challenge/initializers"
	"github.com/hanno-meister/OAuth2Server_challenge/middleware"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.MigrateDB()
	initializers.CreatePrivateKey()
	initializers.CreateJwkKey()
}

func main() {
	router := gin.Default()

	router.POST("/signup", controllers.Signup)
	router.POST("/token", controllers.GetToken)
	router.GET("/signingkeys", controllers.ListSigningKeys)
	router.POST("/introspection", middleware.Authentication, controllers.IntrospectToken)

	router.Run()
}
