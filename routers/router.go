package routers

import (
	"blogweb_gin/controllers"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	//静态资源
	router.Static("/static", "./static")
	router.LoadHTMLGlob("views/*")
	store := cookie.NewStore([]byte("secretKey"))
	router.Use(sessions.Sessions("mysession", store))
	// 注册：
	router.GET("/register", controllers.RegisterGet)
	router.POST("/register", controllers.RegisterPost)
	// 登陆：
	router.GET("/login", controllers.LoginGet)
	router.POST("/login", controllers.LoginPost)
	//首页
	router.GET("/", controllers.HomeGet)
	//退出
	router.GET("/exit", controllers.ExitGet)
	v1 := router.Group("/article")
	{
		v1.GET("/add", controllers.AddArticleGet)
	}

	return router
}
