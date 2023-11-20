package main

import (
	"fmt"
	"os"

	"github.com/subpxl/microservice-lib/logger"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {

	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	DB = connectToDB()

	e := echo.New()
	e.POST("/register", register)

	e.POST("/login", login)

	port := os.Getenv("AUTHSERVICE_PORT")
	if port == "" {
		port = "8000"
	}

	logger.Info(fmt.Sprintf("Starting server on %s:%s ...", "127.0.0.1", port))

	serverAddress := fmt.Sprintf(":%s", port)

	e.Logger.Fatal(e.Start(serverAddress))
}
