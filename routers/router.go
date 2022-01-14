package routers

import (
	"blogweb_gin/controllers"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	// serving static files
	router.Static("/static", "./static")
	router.LoadHTMLGlob("views/*")
	store := cookie.NewStore([]byte("secretKey"))
	router.Use(sessions.Sessions("mysession", store))
	// registration:
	router.GET("/register", controllers.RegisterGet)
	router.POST("/register", controllers.RegisterPost)
	// login:
	router.GET("/login", controllers.LoginGet)
	router.POST("/login", controllers.LoginPost)
	// homepage:
	router.GET("/", controllers.HomeGet)
	// exit/logout:
	router.GET("/exit", controllers.ExitGet)
	v1 := router.Group("/article")
	{
		// add article:
		v1.GET("/add", controllers.AddArticleGet)
		v1.POST("/add", controllers.AddArticlePost)

		// view article:
		v1.GET("/show/:id", controllers.ShowArticleGet)

		// update article:
		v1.GET("/update/:id", controllers.UpdateArticleGet)
		v1.POST("/update/:id", controllers.UpdateArticlePost)

		// delete article:
		v1.GET("/delete/:id", controllers.DeleteArticleGet)
	}
	// showing tags info page
	router.GET("/tags", controllers.TagsGet)

	// showing album
	router.GET("/album", controllers.AlbumGet)

	// uploading picture to the album
	router.POST("/album", controllers.AlbumPost)

	return router
}
