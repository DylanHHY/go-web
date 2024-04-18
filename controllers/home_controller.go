package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HomeIndex(c *gin.Context) {
	message := "Hello World"
	text := "Welcome to world of Golang again :)))"
	c.JSON(http.StatusOK, gin.H{"message": message, "text": text})
}
