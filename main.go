package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/serkanerip/go-pg/controllers"
)

var server controllers.Server

func main() {
	var err error
	isDev := os.Getenv("GO_DEVELOPMENT")
	if isDev != "" {
		err = godotenv.Load()
		if err != nil {
			fmt.Printf("cannot load env file %v", err)
		}
	}

	server.Initialize()
	server.Run(":8080")
}
