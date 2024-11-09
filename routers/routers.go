package routers

import (
	"bubble/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	r := gin.Default()
	//静态文件哪里找
	r.Static("/static", "./static")
	//告诉gin框架哪找模板文件
	r.LoadHTMLGlob("templates/*")
	r.GET("/", controller.IndexHandler)
	//v1
	v1Group := r.Group("v1")
	{
		//待办事项
		//添加
		v1Group.POST("/todo", controller.CrateATodo)
		//查看所有的待办事项
		v1Group.GET("/todo", controller.GetTodoList)
		//查看一个待办事项
		v1Group.GET("/todo/:id", controller.GetTodo)
		//修改一个待办事项
		v1Group.PUT("/todo/:id", controller.UpdateATodo)
		//删除一个待办事项
		v1Group.DELETE("/todo/:id", controller.DeleteATodo)
	}
	return r
}
