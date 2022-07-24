package routers

import (
	"bubble/controller"

	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	r := gin.Default()

	//一定记得设置html目录 不然无法找到
	r.LoadHTMLGlob("templates/*")

	//想要访问工程下的静态文件（图片）前必须进行设置
	r.Static("/static", "static")

	r.GET("/", controller.IndexHandler)

	//v1(api)
	v1Group := r.Group("/v1")
	{
		//代办事项
		//添加
		v1Group.POST("/todo", controller.CreateATodo)

		//查看所有代办事项
		v1Group.GET("/todo", controller.GetTodoList)

		//查看某一个代办事项
		v1Group.GET("/todo/:id", func(ctx *gin.Context) {
			//这个项目里没有
		})

		//修改
		v1Group.PUT("/todo/:id", controller.UpdateTodoList)

		//删除
		v1Group.DELETE("/todo/:id", controller.DeleteTodo)

	}
	return r
}
