package router

import (
	"blogapi/api/config"
	"blogapi/api/controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func init() {
	// veri tabanı bağlantısı
	config.Conn()
}

func Router() {

	e := echo.New()
	user := e.Group("/user")
	user.POST("/register", controller.UserInsert)
	user.POST("/login", controller.UserLogin)
	user.Use(middleware.JWTWithConfig(config.JwtConfig))
	user.GET("/info", controller.UserInfo)
	post := e.Group("/post")
	post.GET("/list", controller.PostList)
	post.Use(middleware.JWTWithConfig(config.JwtConfig))
	post.POST("/insert", controller.PostInsert)
	post.POST("/del", controller.PostDel)
	category := e.Group("/category")
	category.GET("/list", controller.CategoryList)
	category.POST("/insert", controller.CategoryInsert)
	category.POST("/del", controller.CategoryDel)

	e.Start(":8080")

}
