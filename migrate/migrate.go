package main

import (
	"go-side-project/initializers"
	model "go-side-project/models"
)

func init() {
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&model.User{})
	initializers.DB.AutoMigrate(&model.Post{})
}
