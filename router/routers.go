package router

import (
	"go-side-project/controllers"

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
		api.POST("/register", controllers.Register)

		// post api
		api.GET("/posts", controllers.GetAllPosts)
		api.POST("/posts", controllers.CreatePost)
		api.GET("/posts/:id", controllers.ShowPost)
		api.PUT("/posts/:id", controllers.UpdatePost)
		api.DELETE("/posts/:id", controllers.DeletePost)

	}

	return router
}
