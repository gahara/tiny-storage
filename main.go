package main

import (
	"fmt"
	"s3/src/server/handlers"
)

func main() {
	fmt.Println("Starting...")
	router := handlers.SetupRouter()

	router.Run(":8080")
}
