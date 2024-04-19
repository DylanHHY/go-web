package router

import (
	"go-side-project/controllers"
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
	router.Use(static.Serve("/", static.LocalFile("./frontend", false)))

	api := router.Group("api/v1")
	{
		api.GET("/welcome", controllers.HomeIndex)
		api.POST("/login", controllers.Login)
		api.POST("/register", controllers.Register)
	}

	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/index.html")
	})
	router.GET("/login", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/login.html")
	})
	router.GET("/register", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/register.html")
	})

	return router
}
