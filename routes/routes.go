package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rizperdana/prototype-golang-rest/controllers"
)

func Routes(router *gin.Engine) {
	router.GET("/", welcome)

	router.GET("/todo", controllers.GetListTodos)
	router.POST("/todo", controllers.CreateTodo)
	router.POST("/todo/:todoId", controllers.GetDetailTodo)
}

func welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Welcome to  API",
	})
}
