package router

import (
	controllers "go-side-project/controllers"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	// CORS
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = append(config.AllowHeaders, "Authorization")

	router.Use(cors.New(config))
	router.Use(static.Serve("/", static.LocalFile("./fronted", false)))

	router.GET("/api/v1/welcome", func(c *gin.Context) {
		message := "Hello World"
		text := "Welcome to world of Golang again :)))"
		c.JSON(http.StatusOK, gin.H{"message": message, "text": text})
	})

	router.GET("/", controllers.HomeIndex)
	// router.POST("/api/register", controllers.Register)
	// router.POST("api/login", controllers.Login)

	return router
}
