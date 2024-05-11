package main

import (
	"fmt"
	docs "s3/src/docs"
	"s3/src/internal/server/handlers"
	"s3/src/internal/server/helpers"
)

var environmentalVariables = helpers.GetEnvironmentalVariables()

func main() {
	fmt.Println("Starting...")

	router := handlers.SetupRouter(environmentalVariables)
	docs.SwaggerInfo.BasePath = "/"

	router.Run(":8080")
}
