package main

import (
	"fmt"
	"s3/src/server/handlers"
	"s3/src/server/helpers"
)

var EnvironmentarVariables = helpers.GetEnvironmentalVariables()

func main() {
	fmt.Println("Starting...")

	router := handlers.SetupRouter()

	router.Run(":8080")
}
