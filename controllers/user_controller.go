package controllers

import (
	DB "go-side-project/initializers"
	model "go-side-project/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context){
	message := "Hello World"
	text := "Welcome to world of Golang again :)))"
	c.JSON(http.StatusOK, gin.H{"message": message, "text": text})
}

func Register(c *gin.Context){
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request data"})
		return
	}
	user.Password = user.Password + "safety"
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to encrypt password"})
        return
    }

	user.Password = string(hashedPassword)
	user.Role = "USER"

	if err := DB.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})


}