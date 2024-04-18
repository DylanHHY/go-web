package main

import (
	"go-side-project/initializers"
	"go-side-project/router"
)

func init() {
	initializers.ConnectToDB()
}

func main() {
	router := router.InitRouter()
	router.Run()
}
