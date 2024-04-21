package middleware

import (
	"log"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware()gin.HandlerFunc{
	return func(c *gin.Context){
		// Get Authorization in cookie
		tokenString, err := c.Cookie("Authorization")
		log.Println("Authorization Header:", tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		// Parse JWT
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SECRET_KEY")), nil
		})
		
		// Verify that the JWT is valid
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Error",
			})
			c.Abort()
			return
		}

		// Get claims from token
		claims, ok := token.Claims.(jwt.MapClaims)
		if ok {
			c.Set("ID", claims["ID"])
			c.Set("name", claims["name"])
			c.Set("email", claims["email"])
			c.Set("role", claims["role"])
		}

		log.Println(claims)

		c.Next()
	}
}