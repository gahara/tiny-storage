package main

import (
	"fmt"
	docs "s3/src/server/docs"
	"s3/src/server/handlers"
	"s3/src/server/helpers"
)

var EnvironmentarVariables = helpers.GetEnvironmentalVariables()

func main() {
	fmt.Println("Starting...")

	router := handlers.SetupRouter()
	docs.SwaggerInfo.BasePath = "/"

	router.Run(":8080")
}
