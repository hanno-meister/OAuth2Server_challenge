package main

import (
	"fmt"

	"github.com/hanno-meister/OAuth2Server_challenge/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {
	fmt.Println("Hello World!")
}
