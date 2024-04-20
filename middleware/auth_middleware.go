package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware()gin.HandlerFunc{
	return func(c *gin.Context){
		authHeader := c.Request.Header.Get("Authorization")
		log.Println("Authorization Header:", authHeader)
		c.Next()
	}
}