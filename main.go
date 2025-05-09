package main

import (
	"fmt"
	"gouserservice/config"
	"gouserservice/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("My Golang UserService project")

	// load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error loading .env file %v", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Port not set")
	}

	config.ConnectDB()

	router := gin.Default()

	routes.UserRoute(router)

	router.Run(":" + port)
}
