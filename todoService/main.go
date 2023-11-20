package main

import (
	"fmt"
	"os"

	"github.com/gin-contrib/cors"

	"github.com/subpxl/microservice-lib/logger"
)

func main() {

	DB = SetupDB()
	WaitForDB()
	SetupDB()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	config.AllowHeaders = []string{"Content-Type"}

	r := SetupRouter()
	r.Use(cors.New(config))

	port := os.Getenv("TODOSERVICE_PORT")

	if port == "" {
		port = "8000"
	}

	fmt.Println("port is ", port)
	logger.Info(fmt.Sprintf("Starting server on %s:%s ...", "127.0.0.1", port))

	serverAddress := fmt.Sprintf(":%s", port)
	r.Run(serverAddress)

}
