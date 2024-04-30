package main

import (
	"fmt"
	docs "s3/src/internal/server/docs"
	"s3/src/internal/server/handlers"
	"s3/src/internal/server/helpers"
)

var EnvironmentarVariables = helpers.GetEnvironmentalVariables()

func main() {
	fmt.Println("Starting...")

	router := handlers.SetupRouter()
	docs.SwaggerInfo.BasePath = "/"

	router.Run(":8080")
}
