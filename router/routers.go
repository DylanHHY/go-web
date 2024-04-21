package router

import (
	"go-side-project/controllers"
	midd "go-side-project/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	// CORS
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = append(config.AllowHeaders, "Authorization")

	router.Use(cors.New(config))

	api := router.Group("api/v1")
	{

		api.GET("/welcome", controllers.HomeIndex)
		
		// user api
		api.POST("/login", controllers.Login)
		api.DELETE("/logout", controllers.Logout)
		api.POST("/register", controllers.Register)

		// post api
		api.GET("/posts",controllers.GetAllPosts)
		api.POST("/posts", midd.AuthMiddleware(), controllers.CreatePost)
		api.GET("/posts/:id", midd.AuthMiddleware(), controllers.ShowPost)
		api.PUT("/posts/:id", midd.AuthMiddleware(), controllers.UpdatePost)
		api.DELETE("/posts/:id",  midd.AuthMiddleware(), controllers.DeletePost)

	}

	return router
}
