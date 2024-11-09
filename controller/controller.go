package controller

import (
	"bubble/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

//url   -->controller -->logic  -->model
//请求来啦 -->控制器    -->业务逻辑 -->模拟层的增删改查

func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}
func CrateATodo(c *gin.Context) {
	//前端页面填写嗲办事项 点击提交 会发请求到这里
	//1.从请求中把数据拿出来
	var todo models.Todo
	c.BindJSON(&todo)
	//2.存入数据库
	//err = DB.Create(&todo).Error
	//if err!=nil{}
	//3.返回响应
	err := models.CreateATodo(&todo)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, todo)
		//c.JSON(http.StatusOK,gin.H{
		//	"code":2000,
		//	"msg":"success",
		//	"data":todo,
		//})
	}
}
func GetTodoList(c *gin.Context) {
	//查询todo这个表里的所有数据
	todoList, err := models.GetAllTodo()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, todoList)
	}
}
func GetTodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"error": "无效的id",
		})
		return
	}
	todo, err := models.GetATodo(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}
func UpdateATodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"error": "无效的id",
		})
		return
	}
	todo, err := models.GetATodo(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	c.BindJSON(&todo)
	err = models.UpdateATodo(todo)
	if err != nil { //主键存在就更新，不存在就插入
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}
func DeleteATodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
		return
	}
	err := models.DeleteATodo(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{id: "deleted"})
	}
}
