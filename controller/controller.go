package controller

import (
	"bubble/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexHandler(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", nil)
}

func CreateATodo(ctx *gin.Context) {
	//前端页面填写一个代办事项 点击提交 发送请求
	//从请求中把数据拿出 存入数据库
	var todo model.Todo
	if err := ctx.BindJSON(&todo); err != nil {
		log.Fatal(err.Error())
	}

	//存入数据库并返回响应
	err := model.CreateATodo(&todo)
	if err != nil {
		ctx.JSON(200, gin.H{
			"error": err.Error(),
		})
	} else {
		ctx.JSON(200, todo) //返回数据 这个后期前端会有返回规定，这里没有
	}
}

func GetTodoList(ctx *gin.Context) {
	//本质就是查询todos这个表的所有数据
	todoList, err := model.GetAllList()
	if err != nil {
		ctx.JSON(200, gin.H{
			"error": err.Error(),
		})
	} else {
		ctx.JSON(200, todoList)
	}

}

func UpdateTodoList(ctx *gin.Context) {
	id, _ := ctx.Params.Get("id") //拿到id
	todo, err := model.GetATodo(id)
	if err != nil {
		ctx.JSON(200, gin.H{
			"error": err.Error(),
		})
	}
	ctx.BindJSON(&todo)
	if err := model.UpdateATodo(todo); err != nil {
		ctx.JSON(200, gin.H{
			"error": err.Error(),
		})
	} else {
		ctx.JSON(200, todo)
	}
}

func DeleteTodo(ctx *gin.Context) {
	id, ok := ctx.Params.Get("id") //拿到id
	if !ok {
		ctx.JSON(200, gin.H{
			"error": "无效的id",
		})
		return
	}
	if err := model.DeleteATodo(id); err != nil {
		ctx.JSON(200, gin.H{
			"error": err.Error(),
		})
	} else {
		ctx.JSON(200, gin.H{
			id: "deleted",
		})
	}
}
