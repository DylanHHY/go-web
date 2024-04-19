package router

import (
	"go-side-project/controllers"
	"net/http"
	"os"

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

	router.GET("/", renderHTMLPage)
	router.GET("/login", renderHTMLPage)
	router.GET("/register", renderHTMLPage)
	// router.POST("/api/register", controllers.Register)
	// router.POST("api/login", controllers.Login)

	return router
}

func renderHTMLPage(c *gin.Context){
    var filePath string
    switch c.Request.URL.Path {
    case "/login":
        filePath = "login.html"
    case "/register":
        filePath = "register.html"
    default:
        filePath = "index.html"
    }

    // 檢查文件是否存在
    if _, err := os.Stat("./fronted/" + filePath); os.IsNotExist(err) {
        c.String(http.StatusNotFound, "File not found")
        return
    }

    // 返回 HTML 文件
    c.File("./fronted/" + filePath)
}