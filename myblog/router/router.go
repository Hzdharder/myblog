package router

import (
	"github.com/gin-gonic/gin"
	"myblog/controller"
)

func Start() {
	e := gin.Default()
	e.LoadHTMLGlob("templates/*")
	e.Static("/assets", "./assets")

	//e.GET("/index", controller.ListUser)
	e.GET("/index", controller.Index)

	e.GET("/register", controller.GoRegister)
	e.POST("/register", controller.Register)

	e.GET("/login", controller.GoLogin)
	e.POST("/login", controller.Login)

	e.GET("/post", controller.GoAddPost)
	e.POST("/post", controller.AddPost)
	e.GET("/post_index", controller.GetPostIndex)
	e.GET("/delete_post", controller.DeletePost)

	e.GET("/post_detail", controller.PostDetail)

	e.GET("/face", controller.ShowFace)

	e.GET("/", controller.ShowStar)

	e.GET("/modify_post", controller.GoModifyPost)
	e.POST("/modify_post", controller.ModifyPost)
	e.Run()
}
