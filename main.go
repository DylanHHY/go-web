package main

import (
	"go-side-project/initializers"
	"go-side-project/router"
	"log"

	"github.com/joho/godotenv"
)

func init() {
	initializers.ConnectToDB()
}

func main() {
	router := router.InitRouter()

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	router.Run()
}
